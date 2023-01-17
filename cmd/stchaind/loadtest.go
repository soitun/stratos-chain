package main

//import (
//	"bufio"
//	"encoding/json"
//	"fmt"
//	"github.com/cosmos/cosmos-sdk/client/context"
//	"github.com/cosmos/cosmos-sdk/client/flags"
//	"github.com/cosmos/cosmos-sdk/codec"
//	"github.com/cosmos/cosmos-sdk/server"
//	sdk "github.com/cosmos/cosmos-sdk/types"
//	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
//	"github.com/cosmos/cosmos-sdk/x/auth"
//	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"
//	"github.com/cosmos/cosmos-sdk/x/auth/exported"
//	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
//	"github.com/cosmos/cosmos-sdk/x/bank"
//	"github.com/spf13/cobra"
//	"github.com/spf13/viper"
//	"github.com/tendermint/tendermint/libs/cli"
//	tmtypes "github.com/tendermint/tendermint/types"
//	"os"
//	"os/signal"
//	"strconv"
//	"strings"
//	"sync"
//	"time"
//)
//
//const (
//	flagThreads    = "threads"
//	flagInterval   = "interval"
//	flagRandomRecv = "random-recv"
//	flagMaxTx      = "max-tx"
//	flagAddr       = "addr"
//	flagShowTxHash = "show-txhash"
//	flagSkipInit   = "skip-init"
//
//	defaultOutputFlag = "text"
//
//	defaultNodeURI        = "tcp://127.0.0.1:26657"
//	defaultKeyringBackend = "test"
//	defaultHome           = "build/node/stchaincli"
//	defaultDenom          = "wei"
//	defaultChainId        = "test-chain"
//)
//
////var ModuleCdc *codec.Codec
//
//// global to load command line args
//var loadTestArgs = LoadTestArgs{}
//
//// struct to hold the command-line args
//type LoadTestArgs struct {
//	threads    int  // no. of threads in the load test; for concurrency
//	interval   int  // interval (in milliseconds) between two successive send transactions on a thread
//	randomRecv bool // whether to send tokens to a random address every time or no, the default is false
//	maxTx      int  // max transactions after which the load test should stop, default is 10000(10k)
//	address    []byte
//}
//
//func LoadTestCommands(ctx *server.Context, cdc *codec.Codec, defaultNodeHome, defaultClientHome string) *cobra.Command {
//	cmd := &cobra.Command{
//		Use:   "load",
//		Short: "Run a load test",
//		Long:  `Run a load test with fixed senders or random senders`,
//	}
//	cmd.AddCommand(
//		AddFixedLoadTestCmd(ctx, cdc, defaultNodeHome, defaultClientHome),
//		AddRandomLoadTestCmd(ctx, cdc, defaultNodeHome, defaultClientHome),
//	)
//	return cmd
//}
//
//// AddLoadTestCmd returns load test cobra Command.
//func AddFixedLoadTestCmd(
//	ctx *server.Context, cdc *codec.Codec, defaultNodeHome, defaultClientHome string,
//) *cobra.Command {
//
//	cmd := &cobra.Command{
//		Use:   "fixed",
//		Short: "Run a load test with fixed senders and receivers",
//		Args:  cobra.RangeArgs(0, 5),
//		RunE: func(cmd *cobra.Command, args []string) (err error) {
//			config := ctx.Config
//			config.SetRoot(viper.GetString(cli.HomeFlag))
//
//			loadTestArgs.threads = viper.GetInt(flagThreads)
//			loadTestArgs.interval = viper.GetInt(flagInterval)
//			loadTestArgs.maxTx = viper.GetInt(flagMaxTx)
//			loadTestArgs.randomRecv = viper.GetBool(flagRandomRecv)
//			senderAddr := viper.GetString(flagAddr)
//			senderAddrBytes, err := sdk.AccAddressFromBech32(senderAddr)
//			if err != nil {
//				return fmt.Errorf("failed to parse bech32 address: %w", err)
//			}
//			loadTestArgs.address = senderAddrBytes
//
//			if loadTestArgs.address == nil {
//				genesis := ctx.Config.GenesisFile()
//				loadTestArgs.address, err = getFirstAccAddressFromGenesis(cdc, genesis)
//				if err != nil {
//					return fmt.Errorf("failed to parse genesis: %w", err)
//				}
//				fmt.Printf("No sender account specified, using accounts in genesis for load test\n")
//			}
//
//			ctx.Logger.Info("Starting load test with fixed senders/receivers...")
//
//			// create a channel to catch os.Interrupt from a SIGTERM or similar kill signal
//			c := make(chan os.Signal, 1)
//			signal.Notify(c, os.Interrupt)
//
//			// create vars for concurrency management
//			stopChan := make(chan bool, loadTestArgs.threads)
//			waiter := sync.WaitGroup{}
//			counterChan := make(chan int, loadTestArgs.threads)
//
//			// spawn a goroutine to handle sigterm and max transactions
//			waiter.Add(1)
//			counter := 0
//			go handleSigTerm(c, counterChan, stopChan, loadTestArgs.threads, loadTestArgs.maxTx, &waiter, &counter)
//
//			genesis := ctx.Config.GenesisFile()
//			accsFromGenesis, err := getAllAccAddressFromGenesis(cdc, genesis)
//			if err != nil {
//				return fmt.Errorf("failed to parse bech32 address: %w", err)
//			}
//
//			if loadTestArgs.threads > len(accsFromGenesis) {
//				loadTestArgs.threads = len(accsFromGenesis)
//				fmt.Printf("Total available accounts: %d, max threads num set to %d", len(accsFromGenesis), len(accsFromGenesis))
//			}
//			seqStart := make(map[int]uint64)
//			// start threads
//			for i := 0; i < loadTestArgs.threads; i++ {
//				inBuf := bufio.NewReader(cmd.InOrStdin())
//				if !viper.IsSet(flags.FlagChainID) {
//					viper.Set(flags.FlagChainID, defaultChainId)
//				}
//				txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc)).WithChainID(viper.GetString(flags.FlagChainID))
//				viper.Set(flags.FlagBroadcastMode, "async")
//				viper.Set(flags.FlagSkipConfirmation, true)
//				if !viper.IsSet(flags.FlagKeyringBackend) {
//					viper.Set(flags.FlagKeyringBackend, defaultKeyringBackend)
//				}
//				if viper.GetBool(flagShowTxHash) {
//					viper.Set(cli.OutputFlag, defaultOutputFlag)
//				}
//				if !viper.IsSet(flags.FlagNode) {
//					viper.Set(flags.FlagNode, defaultNodeURI)
//				}
//				if !viper.IsSet(flags.FlagHome) {
//					viper.Set(flags.FlagHome, defaultHome)
//				}
//				viper.Set(flags.FlagTrustNode, true)
//				from := accsFromGenesis[i]
//				if len(loadTestArgs.address) != 0 {
//					from = loadTestArgs.address
//				}
//				cliCtx := context.NewCLIContextWithInputAndFrom(inBuf, from.String()).WithCodec(cdc)
//				to := accsFromGenesis[0]
//				if len(accsFromGenesis)-1 > i && (i != loadTestArgs.threads-1 || i == 0) {
//					to = accsFromGenesis[i+1]
//				}
//				ctx.Logger.Info(fmt.Sprintf("thread: %d, from addr %s, to addr %s", i, from, to))
//				txBldr, _ = utils.PrepareTxBuilder(txBldr, cliCtx)
//				ctx.Logger.Info(fmt.Sprintf("thread: %d, first sequence in this thread: %d\n", i, int(txBldr.Sequence())))
//				seqStart[i] = txBldr.Sequence()
//
//				// single-threading start --------
//				//for j := 0; j < loadTestArgs.maxTx; j++ {
//				//	ctx.Logger.Info(fmt.Sprintf("current sequence: %d\n", int(firstSeqUint64+uint64(j))))
//				//	doSendTransaction(cliCtx, txBldr.WithSequence(seqStart[i]+uint64(j)), i, to, from, loadTestArgs.randomRecv, sdk.Coin{Amount: sdk.NewInt(10), Denom: defaultDenom}, seqStart[i]) // send coin to temp account
//				//	counter += 1
//				//}
//				// single-threading end --------
//
//				// multi-threading start --------
//				threadIndex := i
//				threadTo := to
//				threadFrom := from
//				threadTxBldr := txBldr
//				threadCliCtx := cliCtx
//				// start a thread to keep sending transactions after some interval
//				go func(stop chan bool) {
//					waitDuration := getWaitDuration(loadTestArgs.interval)
//					cliCtx.SkipConfirm = true
//					iter := 0
//					for true {
//						currSeqInt := int(seqStart[threadIndex] + uint64(iter))
//						ctx.Logger.Info(fmt.Sprintf("thread: %d, sending tx with sequence: %d\n", threadIndex, currSeqInt))
//						doSendTransaction(threadCliCtx, threadTxBldr.WithSequence(seqStart[threadIndex]+uint64(iter)).WithMemo(strconv.Itoa(currSeqInt)).WithGas(uint64(400000)), threadIndex, threadTo, threadFrom, loadTestArgs.randomRecv, sdk.Coin{Amount: sdk.NewInt(1), Denom: defaultDenom}, seqStart[threadIndex]) // send coin to temp account
//						iter += 1
//						counterChan <- 1
//
//						select {
//						case <-stop:
//							waiter.Done()
//							return
//						default:
//							time.Sleep(waitDuration)
//						}
//					}
//				}(stopChan)
//			}
//			//// wait for all threads to close through sigterm; indefinitely
//			waiter.Wait()
//			// multi-threading end --------
//
//			// print stats
//			fmt.Println("####################################################################")
//			fmt.Println("################        Terminating load test        ###############")
//			fmt.Println("####################################################################")
//			fmt.Printf("################       Messages sent: % 9d      ###############\n", counter)
//			fmt.Println("####################################################################")
//			return nil
//		},
//	}
//
//	cmd.Flags().String(cli.HomeFlag, defaultNodeHome, "node's home directory")
//	cmd.Flags().String(flagClientHome, defaultClientHome, "client's home directory")
//	cmd.Flags().String(flags.FlagKeyringBackend, flags.DefaultKeyringBackend, "Select keyring's backend (os|file|test)")
//	cmd.Flags().Int(flagThreads, 1, "no. of threads in the load test; for concurrency")
//	cmd.Flags().Int(flagInterval, 10, "interval (in milliseconds) between two successive send transactions on a thread")
//	cmd.Flags().Bool(flagRandomRecv, false, "whether to send tokens to a random address every time or no, the default is false")
//	cmd.Flags().Bool(flagShowTxHash, false, "whether to show tx hash after sending it")
//	cmd.Flags().Int(flagMaxTx, 10000, "max transactions after which the load test should stop, default is 10000(10k)")
//	cmd.Flags().String(flagAddr, "", "fund address that load test uses")
//	cmd.Flags().String(flags.FlagChainID, "", "chain id")
//
//	return cmd
//}
//
//// AddLoadTestCmd returns load test cobra Command.
//func AddRandomLoadTestCmd(
//	ctx *server.Context, cdc *codec.Codec, defaultNodeHome, defaultClientHome string,
//) *cobra.Command {
//
//	cmd := &cobra.Command{
//		Use:   "random",
//		Short: "Run a load test with random senders and receivers",
//		Args:  cobra.RangeArgs(0, 1),
//		RunE: func(cmd *cobra.Command, args []string) (err error) {
//			config := ctx.Config
//			config.SetRoot(viper.GetString(cli.HomeFlag))
//
//			loadTestArgs.threads = viper.GetInt(flagThreads)
//			loadTestArgs.interval = viper.GetInt(flagInterval)
//			loadTestArgs.maxTx = viper.GetInt(flagMaxTx)
//			loadTestArgs.randomRecv = viper.GetBool(flagRandomRecv)
//			senderAddr := viper.GetString(flagAddr)
//			senderAddrBytes, err := sdk.AccAddressFromBech32(senderAddr)
//			if err != nil {
//				return fmt.Errorf("failed to parse bech32 address: %w", err)
//			}
//			loadTestArgs.address = senderAddrBytes
//
//			if loadTestArgs.address == nil {
//				genesis := ctx.Config.GenesisFile()
//				loadTestArgs.address, err = getFirstAccAddressFromGenesis(cdc, genesis)
//				if err != nil {
//					return fmt.Errorf("failed to parse genesis: %w", err)
//				}
//				fmt.Printf("No sender account specified, using accounts in genesis for load test\n")
//			}
//
//			ctx.Logger.Info("Starting load test with random senders/receivers...")
//
//			// create a channel to catch os.Interrupt from a SIGTERM or similar kill signal
//			c := make(chan os.Signal, 1)
//			signal.Notify(c, os.Interrupt)
//
//			// create vars for concurrency management
//			stopChan := make(chan bool, loadTestArgs.threads)
//			waiter := sync.WaitGroup{}
//			counterChan := make(chan int, loadTestArgs.threads)
//
//			// spawn a goroutine to handle sigterm and max transactions
//			waiter.Add(1)
//			counter := 0
//			go handleSigTerm(c, counterChan, stopChan, loadTestArgs.threads, loadTestArgs.maxTx, &waiter, &counter)
//
//			genesis := ctx.Config.GenesisFile()
//			accsFromGenesis, err := getAllAccAddressFromGenesis(cdc, genesis)
//			if err != nil {
//				return fmt.Errorf("failed to parse bech32 address: %w", err)
//			}
//
//			fundFrom := accsFromGenesis[0]
//			if len(loadTestArgs.address) != 0 {
//				fundFrom = loadTestArgs.address
//			}
//
//			inBuf := bufio.NewReader(cmd.InOrStdin())
//			if !viper.IsSet(flags.FlagChainID) {
//				viper.Set(flags.FlagChainID, defaultChainId)
//			}
//
//			if !viper.GetBool(flagSkipInit) {
//				initDistribution(ctx, inBuf, cdc, fundFrom, accsFromGenesis)
//			}
//
//			if loadTestArgs.threads > len(accsFromGenesis) {
//				loadTestArgs.threads = len(accsFromGenesis)
//				fmt.Printf("Total available accounts: %d, max threads num set to %d", len(accsFromGenesis), len(accsFromGenesis))
//			}
//
//			fmt.Printf("Start testing in multiple threads\n")
//			seqStart := make(map[int]uint64)
//			// start threads
//			for i := 0; i < loadTestArgs.threads; i++ {
//				//waiter.Add(1)
//				inBuf := bufio.NewReader(cmd.InOrStdin())
//				if !viper.IsSet(flags.FlagChainID) {
//					viper.Set(flags.FlagChainID, defaultChainId)
//				}
//				//if !viper.IsSet("gas") {
//				//	viper.Set("gas", flags.GasFlagAuto)
//				//}
//				txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc)).WithChainID(viper.GetString(flags.FlagChainID))
//				viper.Set(flags.FlagBroadcastMode, "block")
//				viper.Set(flags.FlagSkipConfirmation, true)
//				if !viper.IsSet(flags.FlagKeyringBackend) {
//					viper.Set(flags.FlagKeyringBackend, defaultKeyringBackend)
//				}
//				if viper.GetBool(flagShowTxHash) {
//					viper.Set(cli.OutputFlag, defaultOutputFlag)
//				}
//				if !viper.IsSet(flags.FlagNode) {
//					viper.Set(flags.FlagNode, defaultNodeURI)
//				}
//				if !viper.IsSet(flags.FlagHome) {
//					viper.Set(flags.FlagHome, defaultHome)
//				}
//				viper.Set(flags.FlagTrustNode, true)
//
//				//ctx.Logger.Info(fmt.Sprintf("From addr: %s, chain-id: %s, keyring-backend: %s, home: %s", from.String(),
//				cliCtx := context.NewCLIContextWithInputAndFrom(inBuf, accsFromGenesis[i].String()).WithCodec(cdc)
//				from := accsFromGenesis[i]
//				to := accsFromGenesis[0]
//				if len(accsFromGenesis) > i && (i != loadTestArgs.threads-1 || i == 0) {
//					to = accsFromGenesis[i+1]
//				}
//				ctx.Logger.Info(fmt.Sprintf("thread: %d, from addr %s, to addr %s", i, accsFromGenesis[i], to))
//				txBldr, _ = utils.PrepareTxBuilder(txBldr, cliCtx)
//				ctx.Logger.Info(fmt.Sprintf("thread: %d, first sequence in this thread: %d\n", i, int(txBldr.Sequence())))
//				seqStart[i] = txBldr.Sequence()
//
//				// multi-threading start --------
//				threadIndex := i
//				threadTo := to
//				threadFrom := from
//				threadTxBldr := txBldr
//				threadCliCtx := cliCtx
//				// start a thread to keep sending transactions after some interval
//				go func(stop chan bool) {
//					waitDuration := getWaitDuration(loadTestArgs.interval)
//					cliCtx.SkipConfirm = true
//					iter := 0
//					for true {
//						currSeqInt := int(seqStart[threadIndex] + uint64(iter))
//						ctx.Logger.Info(fmt.Sprintf("thread: %d, sending tx with sequence: %d\n", threadIndex, currSeqInt))
//						doSendTransaction(threadCliCtx, threadTxBldr.WithSequence(seqStart[threadIndex]+uint64(iter)).WithMemo(strconv.Itoa(currSeqInt)).WithGas(uint64(400000)), threadIndex, threadTo, threadFrom, loadTestArgs.randomRecv, sdk.Coin{Amount: sdk.NewInt(1), Denom: defaultDenom}, seqStart[threadIndex]) // send coin to temp account
//						iter += 1
//						counterChan <- 1
//
//						select {
//						case <-stop:
//							waiter.Done()
//							return
//						default:
//							time.Sleep(waitDuration)
//						}
//					}
//				}(stopChan)
//			}
//			//// wait for all threads to close through sigterm; indefinitely
//			waiter.Wait()
//			// multi-threading end --------
//
//			// print stats
//			fmt.Println("####################################################################")
//			fmt.Println("################        Terminating load test        ###############")
//			fmt.Println("####################################################################")
//			fmt.Printf("################       Messages sent: % 9d      ###############\n", counter)
//			fmt.Println("####################################################################")
//			return nil
//		},
//	}
//
//	cmd.Flags().String(cli.HomeFlag, defaultNodeHome, "node's home directory")
//	cmd.Flags().String(flagClientHome, defaultClientHome, "client's home directory")
//	cmd.Flags().String(flags.FlagKeyringBackend, flags.DefaultKeyringBackend, "Select keyring's backend (os|file|test)")
//	cmd.Flags().Int(flagThreads, 1, "no. of threads in the load test; for concurrency")
//	cmd.Flags().Int(flagInterval, 10, "interval (in milliseconds) between two successive send transactions on a thread")
//	cmd.Flags().Bool(flagRandomRecv, true, "whether to send tokens to a random address every time or no, the default is false")
//	cmd.Flags().Bool(flagShowTxHash, false, "whether to show tx hash after sending it")
//	cmd.Flags().Bool(flagSkipInit, false, "whether to skip init distribution")
//	cmd.Flags().Int(flagMaxTx, 10000, "max transactions after which the load test should stop, default is 10000(10k)")
//	cmd.Flags().String(flagAddr, "", "fund address that load test uses")
//	cmd.Flags().String(flags.FlagChainID, "", "chain id")
//
//	return cmd
//}
//
//func initDistribution(ctx *server.Context, inBuf *bufio.Reader, cdc *codec.Codec, fundFrom sdk.AccAddress, accsFromGenesis []sdk.AccAddress) {
//	txBldrFund := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc)).WithChainID(viper.GetString(flags.FlagChainID))
//	viper.Set(flags.FlagBroadcastMode, "async")
//	viper.Set(flags.FlagSkipConfirmation, true)
//	if !viper.IsSet(flags.FlagKeyringBackend) {
//		viper.Set(flags.FlagKeyringBackend, defaultKeyringBackend)
//	}
//	if viper.GetBool(flagShowTxHash) {
//		viper.Set(cli.OutputFlag, defaultOutputFlag)
//	}
//	if !viper.IsSet(flags.FlagNode) {
//		viper.Set(flags.FlagNode, defaultNodeURI)
//	}
//	if !viper.IsSet(flags.FlagHome) {
//		viper.Set(flags.FlagHome, defaultHome)
//	}
//	//if !viper.IsSet("gas") {
//	//	viper.Set("gas", flags.GasFlagAuto)
//	//}
//	viper.Set(flags.FlagTrustNode, true)
//
//	cliCtxFund := context.NewCLIContextWithInputAndFrom(inBuf, fundFrom.String()).WithCodec(cdc)
//	txBldrFund, _ = utils.PrepareTxBuilder(txBldrFund, cliCtxFund)
//	seqStartFund := txBldrFund.Sequence()
//	ctx.Logger.Info(fmt.Sprintf("Preparing funds for each test account, seq starts from %d", int(seqStartFund)))
//	doInitDistribution(ctx, cliCtxFund, txBldrFund, seqStartFund, fundFrom, accsFromGenesis, sdk.Coin{Amount: sdk.NewInt(1000000000), Denom: defaultDenom})
//}
//
//func doInitDistribution(ctx *server.Context, cliCtx context.CLIContext, txBldr authtypes.TxBuilder, seqStartFund uint64, fundAcc sdk.AccAddress, subAccs []sdk.AccAddress, coin sdk.Coin) {
//	for i := 0; i < len(subAccs); i++ {
//		msg := bank.NewMsgSend(fundAcc, subAccs[i], sdk.Coins{coin})
//		currSeqInt := seqStartFund + uint64(i)
//		//// build and sign the transaction, then broadcast to Tendermint
//		err := utils.GenerateOrBroadcastMsgs(cliCtx, txBldr.WithSequence(currSeqInt).WithGas(uint64(400000)), []sdk.Msg{msg})
//		ctx.Logger.Info(fmt.Sprintf("Transfered fund to account %s (index: %d, seq: %d)", subAccs[i].String(), i, int(currSeqInt)))
//		if err != nil {
//			fmt.Println(err)
//		}
//		time.Sleep(getWaitDuration(20)) // 20ms
//	}
//}
//
//// doSendTransaction takes in an account and currency object and sends random amounts of coin from the
//// node account. It prints any errors to ctx.logger and returns
//func doSendTransaction(cliCtx context.CLIContext, txBldr authtypes.TxBuilder, threadNo int, to sdk.AccAddress, from sdk.AccAddress, randomRev bool, coin sdk.Coin, firstSeq uint64) {
//	msg := bank.NewMsgSend(from, to, sdk.Coins{coin})
//	//// build and sign the transaction, then broadcast to Tendermint
//	err := utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
//	if err != nil {
//		fmt.Println(err)
//	}
//}
//
//// handleSigTerm keeps a count of messages sent and if the maximum number of transactions is reached it stops
//// all threads and proceeds to shut down the main thread. If it catches a SIGTERM or a CTRL C it similarly shuts down
//// gracefully. This function is blocking and is called as a go routine.
//func handleSigTerm(c chan os.Signal, counterChan chan int, stopChan chan bool,
//	n int, maxTx int, waiter *sync.WaitGroup, cnt *int) {
//
//	// indefinite loop listens over the counter and os.Signal for interrupt signal
//	for true {
//		select {
//		case <-c:
//			// signal the goroutines to stop
//			for i := 0; i < n; i++ {
//				stopChan <- true
//			}
//			// wait for the goroutines to stop
//			time.Sleep(time.Second)
//
//			waiter.Done()
//
//		case <-counterChan:
//			// increment counter
//			*cnt++
//
//			// send shutdown signal if max no. of transactions is reached
//			if *cnt >= maxTx {
//				c <- os.Interrupt
//			}
//		}
//	}
//}
//
//func getAllAccAddressFromGenesis(cdc *codec.Codec, genesisFilePath string) (accAddrs []sdk.AccAddress, err error) {
//	var genDoc *tmtypes.GenesisDoc
//	if genDoc, err = tmtypes.GenesisDocFromFile(strings.ReplaceAll(genesisFilePath, "cli", "d")); err != nil {
//		return nil, fmt.Errorf("error loading genesis doc from %s: %s", genesisFilePath, err.Error())
//	}
//	var genState map[string]json.RawMessage
//	if err = cdc.UnmarshalJSON(genDoc.AppState, &genState); err != nil {
//		return nil, fmt.Errorf("error unmarshalling genesis doc %s: %s", genesisFilePath, err.Error())
//	}
//	var addresses []sdk.AccAddress
//	auth.GenesisAccountIterator{}.IterateGenesisAccounts(
//		cdc, genState, func(acc exported.Account) (stop bool) {
//			addresses = append(addresses, acc.GetAddress())
//			return false
//		},
//	)
//	if len(addresses) > 0 {
//		return addresses, nil
//	}
//	return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownAddress, "No account initiated in genesis")
//}
//
//func getWaitDuration(interval int) time.Duration {
//	return time.Millisecond * time.Duration(interval)
//}
//
//func getFirstAccAddressFromGenesis(cdc *codec.Codec, genesisFilePath string) (accAddr sdk.AccAddress, err error) {
//	var genDoc *tmtypes.GenesisDoc
//	if genDoc, err = tmtypes.GenesisDocFromFile(strings.ReplaceAll(genesisFilePath, "cli", "d")); err != nil {
//		return nil, fmt.Errorf("error loading genesis doc from %s: %s", genesisFilePath, err.Error())
//	}
//	var genState map[string]json.RawMessage
//	if err = cdc.UnmarshalJSON(genDoc.AppState, &genState); err != nil {
//		return nil, fmt.Errorf("error unmarshalling genesis doc %s: %s", genesisFilePath, err.Error())
//	}
//	var addresses []sdk.AccAddress
//	auth.GenesisAccountIterator{}.IterateGenesisAccounts(
//		cdc, genState, func(acc exported.Account) (stop bool) {
//			addresses = append(addresses, acc.GetAddress())
//			return false
//		},
//	)
//	if len(addresses) > 0 {
//		return addresses[0], nil
//	}
//	return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownAddress, "No account initiated in genesis")
//}
