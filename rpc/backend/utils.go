package backend

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"sort"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	ethtypes "github.com/ethereum/go-ethereum/core/types"

	abci "github.com/cometbft/cometbft/abci/types"
	tmrpctypes "github.com/cometbft/cometbft/rpc/core/types"

	"github.com/stratosnet/stratos-chain/rpc/types"
	evmtypes "github.com/stratosnet/stratos-chain/x/evm/types"
)

type txGasAndReward struct {
	gasUsed uint64
	reward  *big.Int
}

type sortGasAndReward []txGasAndReward

func (s sortGasAndReward) Len() int { return len(s) }
func (s sortGasAndReward) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortGasAndReward) Less(i, j int) bool {
	return s[i].reward.Cmp(s[j].reward) < 0
}

// SetTxDefaults populates tx message with default values in case they are not
// provided on the args
func (b *Backend) SetTxDefaults(args evmtypes.TransactionArgs) (evmtypes.TransactionArgs, error) {
	if args.GasPrice != nil && (args.MaxFeePerGas != nil || args.MaxPriorityFeePerGas != nil) {
		return args, errors.New("both gasPrice and (maxFeePerGas or maxPriorityFeePerGas) specified")
	}

	head := b.CurrentHeader()
	if head == nil {
		return args, errors.New("latest header is nil")
	}

	baseFee, err := b.BaseFee()
	if err != nil {
		return args, err
	}

	// If user specifies both maxPriorityfee and maxFee, then we do not
	// need to consult the chain for defaults. It's definitely a London tx.
	if args.MaxPriorityFeePerGas == nil || args.MaxFeePerGas == nil {
		// In this clause, user left some fields unspecified.
		if baseFee != nil && args.GasPrice == nil {
			if args.MaxPriorityFeePerGas == nil {
				tip, err := b.SuggestGasTipCap()
				if err != nil {
					return args, err
				}
				args.MaxPriorityFeePerGas = (*hexutil.Big)(tip)
			}

			if args.MaxFeePerGas == nil {
				gasFeeCap := new(big.Int).Add(
					(*big.Int)(args.MaxPriorityFeePerGas),
					new(big.Int).Mul(baseFee, big.NewInt(2)),
				)
				args.MaxFeePerGas = (*hexutil.Big)(gasFeeCap)
			}

			if args.MaxFeePerGas.ToInt().Cmp(args.MaxPriorityFeePerGas.ToInt()) < 0 {
				return args, fmt.Errorf("maxFeePerGas (%v) < maxPriorityFeePerGas (%v)", args.MaxFeePerGas, args.MaxPriorityFeePerGas)
			}

		} else {
			if args.MaxFeePerGas != nil || args.MaxPriorityFeePerGas != nil {
				return args, errors.New("maxFeePerGas or maxPriorityFeePerGas specified but london is not active yet")
			}

			if args.GasPrice == nil {
				price, err := b.SuggestGasTipCap()
				if err != nil {
					return args, err
				}
				if baseFee != nil {
					// The legacy tx gas price suggestion should not add 2x base fee
					// because all fees are consumed, so it would result in a spiral
					// upwards.
					price.Add(price, baseFee)
				}
				args.GasPrice = (*hexutil.Big)(price)
			}
		}
	} else {
		// Both maxPriorityfee and maxFee set by caller. Sanity-check their internal relation
		if args.MaxFeePerGas.ToInt().Cmp(args.MaxPriorityFeePerGas.ToInt()) < 0 {
			return args, fmt.Errorf("maxFeePerGas (%v) < maxPriorityFeePerGas (%v)", args.MaxFeePerGas, args.MaxPriorityFeePerGas)
		}
	}

	if args.Value == nil {
		args.Value = new(hexutil.Big)
	}
	if args.Nonce == nil {
		// get the nonce from the account retriever
		// ignore error in case tge account doesn't exist yet
		nonce, err := b.getAccountNonce(*args.From, types.EthPendingBlockNumber)
		if err != nil {
			return args, err
		}
		args.Nonce = (*hexutil.Uint64)(&nonce)
	}

	if args.Data != nil && args.Input != nil && !bytes.Equal(*args.Data, *args.Input) {
		return args, errors.New("both 'data' and 'input' are set and not equal. Please use 'input' to pass transaction call data")
	}

	if args.To == nil {
		// Contract creation
		var input []byte
		if args.Data != nil {
			input = *args.Data
		} else if args.Input != nil {
			input = *args.Input
		}

		if len(input) == 0 {
			return args, errors.New("contract creation without any data provided")
		}
	}

	if args.Gas == nil {
		// For backwards-compatibility reason, we try both input and data
		// but input is preferred.
		input := args.Input
		if input == nil {
			input = args.Data
		}

		callArgs := evmtypes.TransactionArgs{
			From:                 args.From,
			To:                   args.To,
			Gas:                  args.Gas,
			GasPrice:             args.GasPrice,
			MaxFeePerGas:         args.MaxFeePerGas,
			MaxPriorityFeePerGas: args.MaxPriorityFeePerGas,
			Value:                args.Value,
			Data:                 input,
			AccessList:           args.AccessList,
		}

		blockNr := types.NewBlockNumber(big.NewInt(0))
		estimated, err := b.EstimateGas(callArgs, &blockNr)
		if err != nil {
			return args, err
		}
		args.Gas = &estimated
		b.logger.Debug("estimate gas usage automatically", "gas", args.Gas)
	}

	if args.ChainID == nil {
		ethChainId := b.ChainConfig().ChainID
		args.ChainID = (*hexutil.Big)(ethChainId)
	}

	return args, nil
}

// getAccountNonce returns the account nonce for the given account address.
// If the pending value is true, it will iterate over the mempool (pending)
// txs in order to compute and return the pending tx sequence.
// Todo: include the ability to specify a blockNumber
func (b *Backend) getAccountNonce(address common.Address, height types.BlockNumber) (uint64, error) {
	if height == types.EthPendingBlockNumber {
		return b.GetTxPool().Nonce(address), nil
	}
	block, err := b.GetTendermintBlockByNumber(height)
	if err != nil {
		return 0, err
	}

	if block.Block == nil {
		return 0, fmt.Errorf("failed to get block for %d height", height)
	}
	sdkCtx, _, err := b.GetEVMContext().GetSdkContextWithHeader(&block.Block.Header)
	if err != nil {
		return 0, err
	}
	return b.GetEVMKeeper().GetNonce(sdkCtx, address), nil
}

// output: targetOneFeeHistory
func (b *Backend) processBlock(
	tendermintBlock *tmrpctypes.ResultBlock,
	ethBlock *types.Block,
	rewardPercentiles []float64,
	tendermintBlockResult *tmrpctypes.ResultBlockResults,
	targetOneFeeHistory *types.OneFeeHistory,
) error {
	blockHeight := tendermintBlock.Block.Height
	blockBaseFee, err := b.BaseFee()
	if err != nil {
		return err
	}

	// set basefee
	targetOneFeeHistory.BaseFee = blockBaseFee

	// set gas used ratio
	gasLimitUint64 := hexutil.Uint64(ethBlock.GasLimit.ToInt().Uint64())

	gasUsedBig := ethBlock.GasUsed

	gasusedfloat, _ := new(big.Float).SetInt(gasUsedBig.ToInt()).Float64()

	if gasLimitUint64 <= 0 {
		return fmt.Errorf("gasLimit of block height %d should be bigger than 0 , current gaslimit %d", blockHeight, gasLimitUint64)
	}

	gasUsedRatio := gasusedfloat / float64(gasLimitUint64)
	blockGasUsed := gasusedfloat
	targetOneFeeHistory.GasUsedRatio = gasUsedRatio

	rewardCount := len(rewardPercentiles)
	targetOneFeeHistory.Reward = make([]*big.Int, rewardCount)
	for i := 0; i < rewardCount; i++ {
		targetOneFeeHistory.Reward[i] = big.NewInt(0)
	}

	// check tendermintTxs
	tendermintTxs := tendermintBlock.Block.Txs
	tendermintTxResults := tendermintBlockResult.TxsResults
	tendermintTxCount := len(tendermintTxs)

	var sorter sortGasAndReward

	for i := 0; i < tendermintTxCount; i++ {
		eachTendermintTx := tendermintTxs[i]
		eachTendermintTxResult := tendermintTxResults[i]

		tx, err := b.clientCtx.TxConfig.TxDecoder()(eachTendermintTx)
		if err != nil {
			b.logger.Debug("failed to decode transaction in block", "height", blockHeight, "error", err.Error())
			continue
		}
		txGasUsed := uint64(eachTendermintTxResult.GasUsed)
		for _, msg := range tx.GetMsgs() {
			ethMsg, ok := msg.(*evmtypes.MsgEthereumTx)
			if !ok {
				continue
			}
			tx := ethMsg.AsTransaction()
			reward := tx.EffectiveGasTipValue(blockBaseFee)
			if reward == nil {
				reward = big.NewInt(0)
			}
			sorter = append(sorter, txGasAndReward{gasUsed: txGasUsed, reward: reward})
		}
	}

	// return an all zero row if there are no transactions to gather data from
	ethTxCount := len(sorter)
	if ethTxCount == 0 {
		return nil
	}

	sort.Sort(sorter)

	var txIndex int
	sumGasUsed := sorter[0].gasUsed

	for i, p := range rewardPercentiles {
		thresholdGasUsed := uint64(blockGasUsed * p / 100)
		for sumGasUsed < thresholdGasUsed && txIndex < ethTxCount-1 {
			txIndex++
			sumGasUsed += sorter[txIndex].gasUsed
		}
		targetOneFeeHistory.Reward[i] = sorter[txIndex].reward
	}

	return nil
}

// AllTxLogsFromEvents parses all ethereum logs from cosmos events
func AllTxLogsFromEvents(events []abci.Event) ([][]*ethtypes.Log, error) {
	allLogs := make([][]*ethtypes.Log, 0, 4)
	for _, event := range events {
		var logs []*ethtypes.Log
		// <v012 support
		if event.Type == evmtypes.EventTypeTxLog {
			var err error
			logs, err = ParseTxLogsFromEventV011(event)
			if err != nil {
				return nil, err
			}
		} else {
			msg, err := sdk.ParseTypedEvent(event)
			if err != nil {
				// ignore in case of not typed event (in case of not migrated code)
				continue
			}

			evtTxLog, ok := msg.(*evmtypes.EventTxLog)
			if !ok {
				continue
			}

			logs, err = ParseTxLogsFromEvent(evtTxLog)
			if err != nil {
				return nil, err
			}
		}

		allLogs = append(allLogs, logs)
	}
	return allLogs, nil
}

// TxLogsFromEvents parses ethereum logs from cosmos events for specific msg index
func TxLogsFromEvents(events []abci.Event, msgIndex int) ([]*ethtypes.Log, error) {
	for _, event := range events {
		// <v012 support
		if event.Type == evmtypes.EventTypeTxLog {
			if msgIndex > 0 {
				// not the eth tx we want
				msgIndex--
				continue
			}

			return ParseTxLogsFromEventV011(event)
		}

		msg, err := sdk.ParseTypedEvent(event)
		if err != nil {
			// ignore in case of not typed event (in case of not migrated code)
			continue
		}

		evtTxLog, ok := msg.(*evmtypes.EventTxLog)
		if !ok {
			continue
		}

		if msgIndex > 0 {
			// not the eth tx we want
			msgIndex--
			continue
		}

		return ParseTxLogsFromEvent(evtTxLog)
	}
	return nil, fmt.Errorf("eth tx logs not found for message index %d", msgIndex)
}

// ParseTxLogsFromEventV011 parse tx logs from one event
func ParseTxLogsFromEventV011(event abci.Event) ([]*ethtypes.Log, error) {
	logs := make([]*evmtypes.Log, 0, len(event.Attributes))

	for _, attr := range event.Attributes {
		bAttrValue := []byte(attr.Value)

		var log evmtypes.Log

		if err := json.Unmarshal(bAttrValue, &log); err != nil {
			return nil, err
		}

		logs = append(logs, &log)
	}

	return evmtypes.LogsToEthereum(logs), nil
}

// ParseTxLogsFromEvent parse tx logs from one event
func ParseTxLogsFromEvent(evtTxLog *evmtypes.EventTxLog) ([]*ethtypes.Log, error) {
	logBytes := evtTxLog.GetTxLogs()

	var logs []*evmtypes.Log

	if err := json.Unmarshal(logBytes, &logs); err != nil {
		return nil, err
	}
	return evmtypes.LogsToEthereum(logs), nil
}
