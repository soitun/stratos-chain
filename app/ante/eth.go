package ante

import (
	"math/big"

	"cosmossdk.io/errors"
	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	authante "github.com/cosmos/cosmos-sdk/x/auth/ante"

	"github.com/ethereum/go-ethereum/common"
	ethcore "github.com/ethereum/go-ethereum/core"
	ethtypes "github.com/ethereum/go-ethereum/core/types"

	stratos "github.com/stratosnet/stratos-chain/types"
	evmkeeper "github.com/stratosnet/stratos-chain/x/evm/keeper"
	"github.com/stratosnet/stratos-chain/x/evm/statedb"
	"github.com/stratosnet/stratos-chain/x/evm/tracers"
	evmtypes "github.com/stratosnet/stratos-chain/x/evm/types"
)

// EthTxPayloadVerificationDecorator validates base tx payload and check some limitations
type EthTxPayloadVerificationDecorator struct {
}

func NewEthTxPayloadVerificationDecorator() EthTxPayloadVerificationDecorator {
	return EthTxPayloadVerificationDecorator{}
}

// AnteHandle validates msgs count, some signature protection and applied limitations
func (tpvd EthTxPayloadVerificationDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (newCtx sdk.Context, err error) {
	if len(tx.GetMsgs()) > 1 {
		return ctx, errors.Wrapf(sdkerrors.ErrInvalidType, "evm transactions only operates with 1 msg")
	}

	for _, msg := range tx.GetMsgs() {
		msgEthTx, ok := msg.(*evmtypes.MsgEthereumTx)
		if !ok {
			return ctx, errors.Wrapf(sdkerrors.ErrUnknownRequest, "invalid message type %T, expected %T", msg, (*evmtypes.MsgEthereumTx)(nil))
		}

		ethTx := msgEthTx.AsTransaction()
		// EIP-155 only allowed
		if !ethTx.Protected() {
			return ctx, errors.Wrapf(sdkerrors.ErrNotSupported, "legacy pre-eip-155 transactions not supported")
		}

		// forbid EIP-2930 update with access list
		if len(ethTx.AccessList()) > 0 {
			return ctx, errors.Wrapf(sdkerrors.ErrNotSupported, "eip-2930 transactions not supported")
		}
	}

	return next(ctx, tx, simulate)
}

// EthTxCosmosMsgVerifierDecorator validates base tx payload if cosmos msg occured and protect from looping
type EthTxCosmosMsgVerifierDecorator struct {
	evmKeeper EVMKeeper
}

func NewEthTxCosmosMsgVerifierDecorator(ek EVMKeeper) EthTxCosmosMsgVerifierDecorator {
	return EthTxCosmosMsgVerifierDecorator{
		evmKeeper: ek,
	}
}

func (ecmvd EthTxCosmosMsgVerifierDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (newCtx sdk.Context, err error) {
	msg := tx.GetMsgs()[0]
	msgEthTx, ok := msg.(*evmtypes.MsgEthereumTx)
	if !ok {
		return ctx, errors.Wrapf(sdkerrors.ErrUnknownRequest, "invalid message type %T, expected %T", msg, (*evmtypes.MsgEthereumTx)(nil))
	}

	// NOTE: We could not trust From param from msg itself, it should be taken from data and signatures
	params := ecmvd.evmKeeper.GetParams(ctx)
	ethCfg := params.ChainConfig.EthereumConfig()
	signer := ethtypes.MakeSigner(ethCfg, big.NewInt(ctx.BlockHeight()))

	baseFee := ecmvd.evmKeeper.GetBaseFee(ctx, ethCfg)

	coreMsg, err := msgEthTx.AsMessage(signer, baseFee)
	if err != nil {
		return ctx, errors.Wrapf(
			err,
			"failed to create an ethereum core.Message from signer %T", signer,
		)
	}

	// skip check if not cosmos msg
	if !evmtypes.IsCosmosHandler(coreMsg.To()) {
		return next(ctx, tx, simulate)
	}

	cMsg, err := ecmvd.evmKeeper.GetSdkMsg(coreMsg.From().Bytes(), coreMsg.Data())
	if err != nil {
		return ctx, err
	}

	if err := cMsg.ValidateBasic(); err != nil {
		return ctx, err
	}

	return next(ctx, tx, simulate)
}

// EthSigVerificationDecorator validates an ethereum signatures
type EthSigVerificationDecorator struct {
	evmKeeper EVMKeeper
}

// NewEthSigVerificationDecorator creates a new EthSigVerificationDecorator
func NewEthSigVerificationDecorator(ek EVMKeeper) EthSigVerificationDecorator {
	return EthSigVerificationDecorator{
		evmKeeper: ek,
	}
}

// AnteHandle validates checks that the registered chain id is the same as the one on the message, and
// that the signer address matches the one defined on the message.
// It's not skipped for RecheckTx, because it set `From` address which is critical from other ante handler to work.
// Failure in RecheckTx will prevent tx to be included into block, especially when CheckTx succeed, in which case user
// won't see the error message.
func (esvd EthSigVerificationDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (newCtx sdk.Context, err error) {
	params := esvd.evmKeeper.GetParams(ctx)

	ethCfg := params.ChainConfig.EthereumConfig()
	blockNum := big.NewInt(ctx.BlockHeight())
	signer := ethtypes.MakeSigner(ethCfg, blockNum)

	for _, msg := range tx.GetMsgs() {
		msgEthTx, ok := msg.(*evmtypes.MsgEthereumTx)
		if !ok {
			return ctx, errors.Wrapf(sdkerrors.ErrUnknownRequest, "invalid message type %T, expected %T", msg, (*evmtypes.MsgEthereumTx)(nil))
		}

		sender, err := signer.Sender(msgEthTx.AsTransaction())
		if err != nil {
			return ctx, errors.Wrapf(
				sdkerrors.ErrorInvalidSigner,
				"couldn't retrieve sender address ('%s') from the ethereum transaction: %s",
				msgEthTx.From,
				err.Error(),
			)
		}

		// set up the sender to the transaction field if not already
		msgEthTx.From = sender.Hex()
	}

	return next(ctx, tx, simulate)
}

// EthAccountVerificationDecorator validates an account balance checks
type EthAccountVerificationDecorator struct {
	ak        evmtypes.AccountKeeper
	evmKeeper EVMKeeper
}

// NewEthAccountVerificationDecorator creates a new EthAccountVerificationDecorator
func NewEthAccountVerificationDecorator(ak evmtypes.AccountKeeper, ek EVMKeeper) EthAccountVerificationDecorator {
	return EthAccountVerificationDecorator{
		ak:        ak,
		evmKeeper: ek,
	}
}

// AnteHandle validates checks that the sender balance is greater than the total transaction cost.
// The account will be set to store if it doesn't exis, i.e cannot be found on store.
// This AnteHandler decorator will fail if:
// - any of the msgs is not a MsgEthereumTx
// - from address is empty
// - account balance is lower than the transaction cost
func (avd EthAccountVerificationDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (newCtx sdk.Context, err error) {
	if !ctx.IsCheckTx() {
		return next(ctx, tx, simulate)
	}

	for i, msg := range tx.GetMsgs() {
		msgEthTx, ok := msg.(*evmtypes.MsgEthereumTx)
		if !ok {
			return ctx, errors.Wrapf(sdkerrors.ErrUnknownRequest, "invalid message type %T, expected %T", msg, (*evmtypes.MsgEthereumTx)(nil))
		}

		txData, err := evmtypes.UnpackTxData(msgEthTx.Data)
		if err != nil {
			return ctx, errors.Wrapf(err, "failed to unpack tx data any for tx %d", i)
		}

		// sender address should be in the tx cache from the previous AnteHandle call
		from := msgEthTx.GetFrom()
		if from.Empty() {
			return ctx, errors.Wrap(sdkerrors.ErrInvalidAddress, "from address cannot be empty")
		}

		// check whether the sender address is EOA
		fromAddr := common.BytesToAddress(from)
		acct := avd.evmKeeper.GetAccount(ctx, fromAddr)

		if acct == nil {
			acc := avd.ak.NewAccountWithAddress(ctx, from)
			avd.ak.SetAccount(ctx, acc)
			acct = statedb.NewEmptyAccount()
		} else if acct.IsContract() {
			return ctx, errors.Wrapf(sdkerrors.ErrInvalidType,
				"the sender is not EOA: address %s, codeHash <%s>", fromAddr, acct.CodeHash)
		}

		if err := evmkeeper.CheckSenderBalance(sdkmath.NewIntFromBigInt(acct.Balance), txData); err != nil {
			return ctx, errors.Wrap(err, "failed to check sender balance")
		}

	}
	return next(ctx, tx, simulate)
}

// EthGasConsumeDecorator validates enough intrinsic gas for the transaction and
// gas consumption.
type EthGasConsumeDecorator struct {
	evmKeeper    EVMKeeper
	maxGasWanted uint64
}

// NewEthGasConsumeDecorator creates a new EthGasConsumeDecorator
func NewEthGasConsumeDecorator(
	evmKeeper EVMKeeper,
	maxGasWanted uint64,
) EthGasConsumeDecorator {
	return EthGasConsumeDecorator{
		evmKeeper,
		maxGasWanted,
	}
}

// AnteHandle validates that the Ethereum tx message has enough to cover intrinsic gas
// (during CheckTx only) and that the sender has enough balance to pay for the gas cost.
//
// Intrinsic gas for a transaction is the amount of gas that the transaction uses before the
// transaction is executed. The gas is a constant value plus any cost inccured by additional bytes
// of data supplied with the transaction.
//
// This AnteHandler decorator will fail if:
// - the message is not a MsgEthereumTx
// - sender account cannot be found
// - transaction's gas limit is lower than the intrinsic gas
// - user doesn't have enough balance to deduct the transaction fees (gas_limit * gas_price)
// - transaction or block gas meter runs out of gas
// - sets the gas meter limit
func (egcd EthGasConsumeDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (newCtx sdk.Context, err error) {
	params := egcd.evmKeeper.GetParams(ctx)

	ethCfg := params.ChainConfig.EthereumConfig()

	blockHeight := big.NewInt(ctx.BlockHeight())
	homestead := ethCfg.IsHomestead(blockHeight)
	istanbul := ethCfg.IsIstanbul(blockHeight)
	london := ethCfg.IsLondon(blockHeight)
	evmDenom := params.EvmDenom
	gasWanted := uint64(0)
	var events []*evmtypes.EventTx

	for _, msg := range tx.GetMsgs() {
		msgEthTx, ok := msg.(*evmtypes.MsgEthereumTx)
		if !ok {
			return ctx, errors.Wrapf(sdkerrors.ErrUnknownRequest, "invalid message type %T, expected %T", msg, (*evmtypes.MsgEthereumTx)(nil))
		}

		txData, err := evmtypes.UnpackTxData(msgEthTx.Data)
		if err != nil {
			return ctx, errors.Wrap(err, "failed to unpack tx data")
		}

		if ctx.IsCheckTx() {
			// We can't trust the tx gas limit, because we'll refund the unused gas.
			if txData.GetGas() > egcd.maxGasWanted {
				gasWanted += egcd.maxGasWanted
			} else {
				gasWanted += txData.GetGas()
			}
		} else {
			gasWanted += txData.GetGas()
		}

		fees, err := egcd.evmKeeper.DeductTxCostsFromUserBalance(
			ctx,
			*msgEthTx,
			txData,
			evmDenom,
			homestead,
			istanbul,
			london,
		)
		if err != nil {
			return ctx, errors.Wrapf(err, "failed to deduct transaction costs from user balance")
		}

		events = append(events, &evmtypes.EventTx{Fee: fees.String()})
	}

	for _, event := range events {
		_ = ctx.EventManager().EmitTypedEvent(event)
	}

	// TODO: deprecate after https://github.com/cosmos/cosmos-sdk/issues/9514  is fixed on SDK
	blockGasLimit := stratos.BlockGasLimit(ctx)

	// NOTE: safety check
	if blockGasLimit > 0 {
		// generate a copy of the gas pool (i.e block gas meter) to see if we've run out of gas for this block
		// if current gas consumed is greater than the limit, this funcion panics and the error is recovered on the Baseapp
		gasPool := sdk.NewGasMeter(blockGasLimit)
		gasPool.ConsumeGas(ctx.GasMeter().GasConsumedToLimit(), "gas pool check")
	}

	// Set ctx.GasMeter with a limit of GasWanted (gasLimit)
	gasConsumed := ctx.GasMeter().GasConsumed()
	ctx = ctx.WithGasMeter(stratos.NewInfiniteGasMeterWithLimit(gasWanted))
	ctx.GasMeter().ConsumeGas(gasConsumed, "copy gas consumed")

	// we know that we have enough gas on the pool to cover the intrinsic gas
	return next(ctx, tx, simulate)
}

// CanTransferDecorator checks if the sender is allowed to transfer funds according to the EVM block
// context rules.
type CanTransferDecorator struct {
	evmKeeper EVMKeeper
}

// NewCanTransferDecorator creates a new CanTransferDecorator instance.
func NewCanTransferDecorator(evmKeeper EVMKeeper) CanTransferDecorator {
	return CanTransferDecorator{
		evmKeeper: evmKeeper,
	}
}

// AnteHandle creates an EVM from the message and calls the BlockContext CanTransfer function to
// see if the address can execute the transaction.
func (ctd CanTransferDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (sdk.Context, error) {
	params := ctd.evmKeeper.GetParams(ctx)
	ethCfg := params.ChainConfig.EthereumConfig()
	signer := ethtypes.MakeSigner(ethCfg, big.NewInt(ctx.BlockHeight()))

	for _, msg := range tx.GetMsgs() {
		msgEthTx, ok := msg.(*evmtypes.MsgEthereumTx)
		if !ok {
			return ctx, errors.Wrapf(sdkerrors.ErrUnknownRequest, "invalid message type %T, expected %T", msg, (*evmtypes.MsgEthereumTx)(nil))
		}

		baseFee := ctd.evmKeeper.GetBaseFee(ctx, ethCfg)

		coreMsg, err := msgEthTx.AsMessage(signer, baseFee)
		if err != nil {
			return ctx, errors.Wrapf(
				err,
				"failed to create an ethereum core.Message from signer %T", signer,
			)
		}

		// NOTE: pass in an empty coinbase address and nil tracer as we don't need them for the check below
		cfg := &evmtypes.EVMConfig{
			ChainConfig: ethCfg,
			Params:      params,
			CoinBase:    common.Address{},
			BaseFee:     baseFee,
		}
		stateDB := statedb.New(ctx, ctd.evmKeeper, statedb.NewEmptyTxConfig(common.BytesToHash(ctx.HeaderHash().Bytes())))
		evm := ctd.evmKeeper.NewEVM(ctx, coreMsg, cfg, tracers.NewNoOpTracer(), stateDB)

		// check that caller has enough balance to cover asset transfer for **topmost** call
		// NOTE: here the gas consumed is from the context with the infinite gas meter
		if coreMsg.Value().Sign() > 0 && !evm.Context.CanTransfer(stateDB, coreMsg.From(), coreMsg.Value()) {
			return ctx, errors.Wrapf(
				sdkerrors.ErrInsufficientFunds,
				"failed to transfer %s from address %s using the EVM block context transfer function",
				coreMsg.Value(),
				coreMsg.From(),
			)
		}

		if evmtypes.IsLondon(ethCfg, ctx.BlockHeight()) {
			if baseFee == nil {
				return ctx, errors.Wrap(
					evmtypes.ErrInvalidBaseFee,
					"base fee is supported but evm block context value is nil",
				)
			}
			if coreMsg.GasFeeCap().Cmp(baseFee) < 0 {
				return ctx, errors.Wrapf(
					sdkerrors.ErrInsufficientFee,
					"max fee per gas less than block base fee (%s < %s)",
					coreMsg.GasFeeCap(), baseFee,
				)
			}
		}
	}

	return next(ctx, tx, simulate)
}

// EthIncrementSenderSequenceDecorator increments the sequence of the signers.
type EthIncrementSenderSequenceDecorator struct {
	ak evmtypes.AccountKeeper
}

// NewEthIncrementSenderSequenceDecorator creates a new EthIncrementSenderSequenceDecorator.
func NewEthIncrementSenderSequenceDecorator(ak evmtypes.AccountKeeper) EthIncrementSenderSequenceDecorator {
	return EthIncrementSenderSequenceDecorator{
		ak: ak,
	}
}

// AnteHandle handles incrementing the sequence of the signer (i.e sender). If the transaction is a
// contract creation, the nonce will be incremented during the transaction execution and not within
// this AnteHandler decorator.
func (issd EthIncrementSenderSequenceDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (sdk.Context, error) {
	for _, msg := range tx.GetMsgs() {
		msgEthTx, ok := msg.(*evmtypes.MsgEthereumTx)
		if !ok {
			return ctx, errors.Wrapf(sdkerrors.ErrUnknownRequest, "invalid message type %T, expected %T", msg, (*evmtypes.MsgEthereumTx)(nil))
		}

		txData, err := evmtypes.UnpackTxData(msgEthTx.Data)
		if err != nil {
			return ctx, errors.Wrap(err, "failed to unpack tx data")
		}

		// increase sequence of sender
		acc := issd.ak.GetAccount(ctx, msgEthTx.GetFrom())
		if acc == nil {
			return ctx, errors.Wrapf(
				sdkerrors.ErrUnknownAddress,
				"account %s is nil", common.BytesToAddress(msgEthTx.GetFrom().Bytes()),
			)
		}
		nonce := acc.GetSequence()

		// we merged the nonce verification to nonce increment, so when tx includes multiple messages
		// with same sender, they'll be accepted.
		if txData.GetNonce() != nonce {
			return ctx, errors.Wrapf(
				sdkerrors.ErrInvalidSequence,
				"invalid nonce; got %d, expected %d", txData.GetNonce(), nonce,
			)
		}

		if err := acc.SetSequence(nonce + 1); err != nil {
			return ctx, errors.Wrapf(err, "failed to set sequence to %d", acc.GetSequence()+1)
		}

		issd.ak.SetAccount(ctx, acc)
	}

	return next(ctx, tx, simulate)
}

// EthValidateBasicDecorator is adapted from ValidateBasicDecorator from cosmos-sdk, it ignores ErrNoSignatures
type EthValidateBasicDecorator struct {
	evmKeeper EVMKeeper
}

// NewEthValidateBasicDecorator creates a new EthValidateBasicDecorator
func NewEthValidateBasicDecorator(ek EVMKeeper) EthValidateBasicDecorator {
	return EthValidateBasicDecorator{
		evmKeeper: ek,
	}
}

// AnteHandle handles basic validation of tx
func (vbd EthValidateBasicDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (sdk.Context, error) {
	// no need to validate basic on recheck tx, call next antehandler
	if ctx.IsReCheckTx() {
		return next(ctx, tx, simulate)
	}

	err := tx.ValidateBasic()
	// ErrNoSignatures is fine with eth tx
	if err != nil && !errors.IsOf(err, sdkerrors.ErrNoSignatures) {
		return ctx, errors.Wrap(err, "tx basic validation failed")
	}

	// For eth type cosmos tx, some fields should be veified as zero values,
	// since we will only verify the signature against the hash of the MsgEthereumTx.Data
	if wrapperTx, ok := tx.(protoTxProvider); ok {
		protoTx := wrapperTx.GetProtoTx()
		body := protoTx.Body
		if body.Memo != "" || body.TimeoutHeight != uint64(0) || len(body.NonCriticalExtensionOptions) > 0 {
			return ctx, errors.Wrap(sdkerrors.ErrInvalidRequest,
				"for eth tx body Memo TimeoutHeight NonCriticalExtensionOptions should be empty")
		}

		if len(body.ExtensionOptions) != 1 {
			return ctx, errors.Wrap(sdkerrors.ErrInvalidRequest, "for eth tx length of ExtensionOptions should be 1")
		}

		txFee := sdk.Coins{}
		txGasLimit := uint64(0)

		params := vbd.evmKeeper.GetParams(ctx)
		ethCfg := params.ChainConfig.EthereumConfig()
		baseFee := vbd.evmKeeper.GetBaseFee(ctx, ethCfg)

		for _, msg := range protoTx.GetMsgs() {
			msgEthTx, ok := msg.(*evmtypes.MsgEthereumTx)
			if !ok {
				return ctx, errors.Wrapf(sdkerrors.ErrUnknownRequest, "invalid message type %T, expected %T", msg, (*evmtypes.MsgEthereumTx)(nil))
			}

			txGasLimit += msgEthTx.GetGas()

			txData, err := evmtypes.UnpackTxData(msgEthTx.Data)
			if err != nil {
				return ctx, errors.Wrap(err, "failed to unpack MsgEthereumTx Data")
			}

			ethTx := ethtypes.NewTx(txData.AsEthereumData())
			// Reject transactions over defined size to prevent DOS attacks
			if uint64(ethTx.Size()) > evmtypes.TxMaxSize {
				return ctx, errors.Wrapf(sdkerrors.ErrInvalidRequest, ethcore.ErrOversizedData.Error())
			}

			// return error if contract creation or call are disabled through governance
			if !params.EnableCreate && txData.GetTo() == nil {
				return ctx, errors.Wrap(evmtypes.ErrCreateDisabled, "failed to create new contract")
			} else if !params.EnableCall && txData.GetTo() != nil {
				return ctx, errors.Wrap(evmtypes.ErrCallDisabled, "failed to call contract")
			}

			if baseFee == nil && txData.TxType() == ethtypes.DynamicFeeTxType {
				return ctx, errors.Wrap(ethtypes.ErrTxTypeNotSupported, "dynamic fee tx not supported")
			}

			txFee = txFee.Add(sdk.NewCoin(params.EvmDenom, sdkmath.NewIntFromBigInt(txData.Fee())))
		}

		authInfo := protoTx.AuthInfo
		if len(authInfo.SignerInfos) > 0 {
			return ctx, errors.Wrap(sdkerrors.ErrInvalidRequest, "for eth tx AuthInfo SignerInfos should be empty")
		}

		if authInfo.Fee.Payer != "" || authInfo.Fee.Granter != "" {
			return ctx, errors.Wrap(sdkerrors.ErrInvalidRequest, "for eth tx AuthInfo Fee payer and granter should be empty")
		}

		if !authInfo.Fee.Amount.IsEqual(txFee) {
			return ctx, errors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid AuthInfo Fee Amount (%s != %s)", authInfo.Fee.Amount, txFee)
		}

		if authInfo.Fee.GasLimit != txGasLimit {
			return ctx, errors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid AuthInfo Fee GasLimit (%d != %d)", authInfo.Fee.GasLimit, txGasLimit)
		}

		sigs := protoTx.Signatures
		if len(sigs) > 0 {
			return ctx, errors.Wrap(sdkerrors.ErrInvalidRequest, "for eth tx Signatures should be empty")
		}
	}

	return next(ctx, tx, simulate)
}

// EthSetupContextDecorator is adapted from SetUpContextDecorator from cosmos-sdk, it ignores gas consumption
// by setting the gas meter to infinite
type EthSetupContextDecorator struct {
	evmKeeper EVMKeeper
}

func NewEthSetUpContextDecorator(evmKeeper EVMKeeper) EthSetupContextDecorator {
	return EthSetupContextDecorator{
		evmKeeper: evmKeeper,
	}
}

func (esc EthSetupContextDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (newCtx sdk.Context, err error) {
	// all transactions must implement GasTx
	_, ok := tx.(authante.GasTx)
	if !ok {
		return newCtx, errors.Wrap(sdkerrors.ErrTxDecode, "Tx must be GasTx")
	}

	newCtx = ctx.WithGasMeter(sdk.NewInfiniteGasMeter())
	// Reset transient gas used to prepare the execution of current cosmos tx.
	// Transient gas-used is necessary to sum the gas-used of cosmos tx, when it contains multiple eth msgs.
	esc.evmKeeper.ResetTransientGasUsed(ctx)
	return next(newCtx, tx, simulate)
}

// EthMempoolFeeDecorator will check if the transaction's effective fee is at least as large
// as the local validator's minimum gasFee (defined in validator config).
// If fee is too low, decorator returns error and tx is rejected from mempool.
// Note this only applies when ctx.CheckTx = true
// If fee is high enough or not CheckTx, then call next AnteHandler
// CONTRACT: Tx must implement FeeTx to use MempoolFeeDecorator
type EthMempoolFeeDecorator struct {
	evmKeeper EVMKeeper
}

func NewEthMempoolFeeDecorator(ek EVMKeeper) EthMempoolFeeDecorator {
	return EthMempoolFeeDecorator{
		evmKeeper: ek,
	}
}

// AnteHandle ensures that the provided fees meet a minimum threshold for the validator,
// if this is a CheckTx. This is only for local mempool purposes, and thus
// is only ran on check tx.
// It only do the check if london hardfork not enabled or feemarket not enabled, because in that case feemarket will take over the task.
func (mfd EthMempoolFeeDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (newCtx sdk.Context, err error) {
	if ctx.IsCheckTx() && !simulate {
		params := mfd.evmKeeper.GetParams(ctx)
		ethCfg := params.ChainConfig.EthereumConfig()
		evmDenom := params.EvmDenom
		baseFee := mfd.evmKeeper.GetBaseFee(ctx, ethCfg)

		for _, msg := range tx.GetMsgs() {
			ethMsg, ok := msg.(*evmtypes.MsgEthereumTx)
			if !ok {
				return ctx, errors.Wrapf(sdkerrors.ErrUnknownRequest, "invalid message type %T, expected %T", msg, (*evmtypes.MsgEthereumTx)(nil))
			}

			var feeAmt *big.Int
			// if base fee disabled, calculate without effective fee mechanics
			if baseFee == nil {
				feeAmt = ethMsg.GetFee()
			} else {
				feeAmt = ethMsg.GetEffectiveFee(baseFee)
			}
			glDec := sdkmath.LegacyNewDec(int64(ethMsg.GetGas()))
			requiredFee := ctx.MinGasPrices().AmountOf(evmDenom).Mul(glDec)
			if sdkmath.LegacyNewDecFromBigInt(feeAmt).LT(requiredFee) {
				return ctx, errors.Wrapf(sdkerrors.ErrInsufficientFee, "insufficient fees; got: %s required: %s", feeAmt, requiredFee)
			}
		}
	}

	return next(ctx, tx, simulate)
}
