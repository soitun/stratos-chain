package register

//
//import (
//	"os"
//	"testing"
//
//	sdk "github.com/cosmos/cosmos-sdk/types"
//	"github.com/cosmos/cosmos-sdk/x/auth"
//	"github.com/cosmos/cosmos-sdk/x/auth/ante"
//	"github.com/cosmos/cosmos-sdk/x/bank"
//	"github.com/cosmos/cosmos-sdk/x/mock"
//	"github.com/cosmos/cosmos-sdk/x/staking"
//	"github.com/cosmos/cosmos-sdk/x/supply"
//	supplyexported "github.com/cosmos/cosmos-sdk/x/supply/exported"
//	"github.com/stratosnet/stratos-chain/helpers"
//	stratos "github.com/stratosnet/stratos-chain/types"
//	"github.com/stratosnet/stratos-chain/x/register/types"
//	"github.com/stretchr/testify/require"
//	abci "github.com/tendermint/tendermint/abci/types"
//)
//
//func TestMain(m *testing.M) {
//	config := stratos.GetConfig()
//	config.Seal()
//	exitVal := m.Run()
//	os.Exit(exitVal)
//}
//
//func Test(t *testing.T) {
//
//	/********************* initialize mock app *********************/
//	//mApp, k, accountKeeper, bankKeeper, stakingKeeper, registerKeeper := getMockApp(t)
//	mApp, k, _, _ := getMockApp(t)
//	accounts := setupAccounts(mApp)
//	mock.SetGenesis(mApp, accounts)
//
//	header := abci.Header{}
//	ctx := mApp.BaseApp.NewContext(true, header)
//
//	//2 bonded resource node, 1 bonded indexing node, 1 unBonded indexing node initialized by genesis
//	resBondedToken := k.GetResourceNodeBondedToken(ctx)
//	require.EqualValues(t, resBondedToken, sdk.NewCoin(k.BondDenom(ctx), resNodeInitStake.Add(resNodeInitStake)))
//	resNotBondedToken := k.GetResourceNodeNotBondedToken(ctx)
//	require.EqualValues(t, resNotBondedToken, sdk.NewCoin(k.BondDenom(ctx), sdk.ZeroInt()))
//	idxBondedToken := k.GetIndexingNodeBondedToken(ctx)
//	require.EqualValues(t, idxBondedToken, sdk.NewCoin(k.BondDenom(ctx), idxNodeInitStake))
//	idxNotBondedToken := k.GetIndexingNodeNotBondedToken(ctx)
//	require.EqualValues(t, idxNotBondedToken, sdk.NewCoin(k.BondDenom(ctx), idxNodeInitStake))
//
//	/********************* send register resource node msg *********************/
//	header = abci.Header{Height: mApp.LastBlockHeight() + 1}
//	ctx = mApp.BaseApp.NewContext(true, header)
//	registerResNodeMsg := types.NewMsgCreateResourceNode(resNodeNetworkId2, resNodePubKey2, sdk.NewCoin(k.BondDenom(ctx), resNodeInitStake), resOwnerAddr2, NewDescription("sds://resourceNode2", "", "", "", ""), 4)
//	resNodeOwnerAcc2 := mApp.AccountKeeper.GetAccount(ctx, resOwnerAddr2)
//	accNumOwner := resNodeOwnerAcc2.GetAccountNumber()
//	accSeqOwner := resNodeOwnerAcc2.GetSequence()
//
//	resNodeAcc2 := mApp.AccountKeeper.GetAccount(ctx, resNodeAddr2)
//	accNumNode := resNodeAcc2.GetAccountNumber()
//	accSeqNode := resNodeAcc2.GetSequence()
//
//	mock.SignCheckDeliver(t, mApp.Cdc, mApp.BaseApp, header, []sdk.Msg{registerResNodeMsg}, []uint64{accNumOwner, accNumNode}, []uint64{accSeqOwner, accSeqNode}, true, true, resOwnerPrivKey2)
//
//	/*-------------------- commit & check result --------------------*/
//	header = abci.Header{Height: mApp.LastBlockHeight() + 1}
//	ctx = mApp.BaseApp.NewContext(true, header)
//
//	resBondedToken = k.GetResourceNodeBondedToken(ctx)
//	require.EqualValues(t, resBondedToken, sdk.NewCoin(k.BondDenom(ctx), resNodeInitStake.Add(resNodeInitStake).Add(resNodeInitStake)))
//	resNotBondedToken = k.GetResourceNodeNotBondedToken(ctx)
//	require.EqualValues(t, resNotBondedToken, sdk.NewCoin(k.BondDenom(ctx), sdk.ZeroInt()))
//	idxBondedToken = k.GetIndexingNodeBondedToken(ctx)
//	require.EqualValues(t, idxBondedToken, sdk.NewCoin(k.BondDenom(ctx), idxNodeInitStake))
//	idxNotBondedToken = k.GetIndexingNodeNotBondedToken(ctx)
//	require.EqualValues(t, idxNotBondedToken, sdk.NewCoin(k.BondDenom(ctx), idxNodeInitStake))
//
//	/********************* send register indexing node msg *********************/
//	header = abci.Header{Height: mApp.LastBlockHeight() + 1}
//	ctx = mApp.BaseApp.NewContext(true, header)
//	registerIdxNodeMsg := types.NewMsgCreateIndexingNode(idxNodeNetworkId3, idxNodePubKey3, sdk.NewCoin(k.BondDenom(ctx), idxNodeInitStake), idxOwnerAddr3, NewDescription("sds://indexingNode3", "", "", "", ""))
//	idxOwnerAcc3 := mApp.AccountKeeper.GetAccount(ctx, idxOwnerAddr3)
//	accNumOwner = idxOwnerAcc3.GetAccountNumber()
//	accSeqOwner = idxOwnerAcc3.GetSequence()
//
//	idxNodeAcc3 := mApp.AccountKeeper.GetAccount(ctx, idxNodeAddr3)
//	accNumNode = idxNodeAcc3.GetAccountNumber()
//	accSeqNode = idxNodeAcc3.GetSequence()
//
//	mock.SignCheckDeliver(t, mApp.Cdc, mApp.BaseApp, header, []sdk.Msg{registerIdxNodeMsg}, []uint64{accNumOwner, accNumNode}, []uint64{accSeqOwner, accSeqNode}, true, true, idxOwnerPrivKey3)
//
//	/*-------------------- commit & check result, stake should be stored in the not bonded pool --------------------*/
//	header = abci.Header{Height: mApp.LastBlockHeight() + 1}
//	ctx = mApp.BaseApp.NewContext(true, header)
//
//	resBondedToken = k.GetResourceNodeBondedToken(ctx)
//	require.EqualValues(t, resBondedToken, sdk.NewCoin(k.BondDenom(ctx), resNodeInitStake.Add(resNodeInitStake).Add(resNodeInitStake)))
//	resNotBondedToken = k.GetResourceNodeNotBondedToken(ctx)
//	require.EqualValues(t, resNotBondedToken, sdk.NewCoin(k.BondDenom(ctx), sdk.ZeroInt()))
//	idxBondedToken = k.GetIndexingNodeBondedToken(ctx)
//	require.EqualValues(t, idxBondedToken, sdk.NewCoin(k.BondDenom(ctx), idxNodeInitStake))
//	idxNotBondedToken = k.GetIndexingNodeNotBondedToken(ctx)
//	require.EqualValues(t, idxNotBondedToken, sdk.NewCoin(k.BondDenom(ctx), idxNodeInitStake.Add(idxNodeInitStake)))
//
//	/********************* deliver tx to vote *********************/
//	header = abci.Header{Height: mApp.LastBlockHeight() + 1}
//	ctx = mApp.BaseApp.NewContext(true, header)
//	voteMsg := types.NewMsgIndexingNodeRegistrationVote(idxNodeNetworkId3, idxOwnerAddr3, types.Approve, idxNodeNetworkId1, idxOwnerAddr1)
//	idxOwnerAcc1 := mApp.AccountKeeper.GetAccount(ctx, idxOwnerAddr1)
//	accNumOwner = idxOwnerAcc1.GetAccountNumber()
//	accSeqOwner = idxOwnerAcc1.GetSequence()
//
//	mock.SignCheckDeliver(t, mApp.Cdc, mApp.BaseApp, header, []sdk.Msg{voteMsg}, []uint64{accNumOwner}, []uint64{accSeqOwner}, true, true, idxOwnerPrivKey1)
//
//	/*-------------------- commit & check result, stake should be transferred to the bonded pool --------------------*/
//	header = abci.Header{Height: mApp.LastBlockHeight() + 1}
//	ctx = mApp.BaseApp.NewContext(true, header)
//
//	resBondedToken = k.GetResourceNodeBondedToken(ctx)
//	require.EqualValues(t, resBondedToken, sdk.NewCoin(k.BondDenom(ctx), resNodeInitStake.Add(resNodeInitStake).Add(resNodeInitStake)))
//	resNotBondedToken = k.GetResourceNodeNotBondedToken(ctx)
//	require.EqualValues(t, resNotBondedToken, sdk.NewCoin(k.BondDenom(ctx), sdk.ZeroInt()))
//	idxBondedToken = k.GetIndexingNodeBondedToken(ctx)
//	require.EqualValues(t, idxBondedToken, sdk.NewCoin(k.BondDenom(ctx), idxNodeInitStake.Add(idxNodeInitStake)))
//	idxNotBondedToken = k.GetIndexingNodeNotBondedToken(ctx)
//	require.EqualValues(t, idxNotBondedToken, sdk.NewCoin(k.BondDenom(ctx), idxNodeInitStake))
//
//}
//
//func getMockApp(t *testing.T) (*mock.App, Keeper, bank.Keeper, supply.Keeper) {
//	mApp := mock.NewApp()
//
//	RegisterCodec(mApp.Cdc)
//	bank.RegisterCodec(mApp.Cdc)
//	supply.RegisterCodec(mApp.Cdc)
//	staking.RegisterCodec(mApp.Cdc)
//
//	keySupply := sdk.NewKVStoreKey(supply.StoreKey)
//	keyStaking := sdk.NewKVStoreKey(staking.StoreKey)
//	keyRegister := sdk.NewKVStoreKey(StoreKey)
//
//	feeCollector := supply.NewEmptyModuleAccount(auth.FeeCollectorName)
//	notBondedPool := supply.NewEmptyModuleAccount(staking.NotBondedPoolName, supply.Burner, supply.Staking)
//	bondPool := supply.NewEmptyModuleAccount(staking.BondedPoolName, supply.Burner, supply.Staking)
//
//	blacklistedAddrs := make(map[string]bool)
//	blacklistedAddrs[feeCollector.GetAddress().String()] = true
//	blacklistedAddrs[notBondedPool.GetAddress().String()] = true
//	blacklistedAddrs[bondPool.GetAddress().String()] = true
//
//	bankKeeper := bank.NewBaseKeeper(mApp.AccountKeeper, mApp.ParamsKeeper.Subspace(bank.DefaultParamspace), blacklistedAddrs)
//	maccPerms := map[string][]string{
//		auth.FeeCollectorName:     {"fee_collector"},
//		staking.NotBondedPoolName: {supply.Burner, supply.Staking},
//		staking.BondedPoolName:    {supply.Burner, supply.Staking},
//	}
//	supplyKeeper := supply.NewKeeper(mApp.Cdc, keySupply, mApp.AccountKeeper, bankKeeper, maccPerms)
//	stakingKeeper := staking.NewKeeper(mApp.Cdc, keyStaking, supplyKeeper, mApp.ParamsKeeper.Subspace(staking.DefaultParamspace))
//	keeper := NewKeeper(mApp.Cdc, keyRegister, mApp.ParamsKeeper.Subspace(DefaultParamSpace), mApp.AccountKeeper, bankKeeper)
//
//	anteHandler := ante.NewAnteHandler(mApp.AccountKeeper, supplyKeeper, helpers.StSigVerificationGasConsumer)
//	mApp.SetAnteHandler(anteHandler)
//
//	mApp.Router().AddRoute(bank.RouterKey, bank.NewHandler(bankKeeper))
//	mApp.Router().AddRoute(staking.RouterKey, staking.NewHandler(stakingKeeper))
//	mApp.Router().AddRoute(RouterKey, NewHandler(keeper))
//	mApp.SetEndBlocker(getEndBlocker(keeper))
//	mApp.SetInitChainer(getInitChainer(mApp, keeper, mApp.AccountKeeper, supplyKeeper,
//		[]supplyexported.ModuleAccountI{feeCollector, notBondedPool, bondPool}, stakingKeeper, bankKeeper))
//
//	err := mApp.CompleteSetup(keyStaking, keySupply, keyRegister)
//	require.NoError(t, err)
//
//	return mApp, keeper, bankKeeper, supplyKeeper
//}
//
//// getInitChainer initializes the chainer of the mock app and sets the genesis
//// state. It returns an empty ResponseInitChain.
//func getInitChainer(mapp *mock.App, keeper Keeper, accountKeeper auth.AccountKeeper, supplyKeeper supply.Keeper,
//	blacklistedAddrs []supplyexported.ModuleAccountI, stakingKeeper staking.Keeper, bankKeeper bank.Keeper) sdk.InitChainer {
//	return func(ctx sdk.Context, req abci.RequestInitChain) abci.ResponseInitChain {
//		// set module accounts
//		for _, macc := range blacklistedAddrs {
//			supplyKeeper.SetModuleAccount(ctx, macc)
//		}
//		mapp.InitChainer(ctx, req)
//
//		// set module accounts
//		for _, macc := range blacklistedAddrs {
//			supplyKeeper.SetModuleAccount(ctx, macc)
//		}
//		stakingGenesis := staking.NewGenesisState(staking.NewParams(staking.DefaultUnbondingTime, staking.DefaultMaxValidators, staking.DefaultMaxEntries, 0, "wei"), nil, nil)
//		totalSupply := sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(100000000000000)))
//		supplyKeeper.SetSupply(ctx, supply.NewSupply(totalSupply))
//
//		// set module accounts
//		for _, macc := range blacklistedAddrs {
//			supplyKeeper.SetModuleAccount(ctx, macc)
//		}
//		validators := staking.InitGenesis(ctx, stakingKeeper, accountKeeper, supplyKeeper, stakingGenesis)
//		bankGenesis := bank.NewGenesisState(true)
//		bank.InitGenesis(ctx, bankKeeper, bankGenesis)
//
//		resourceNodes := setupAllResourceNodes()
//		indexingNodes := setupAllIndexingNodes()
//
//		registerGenesis := NewGenesisState(
//			DefaultParams(),
//			resourceNodes,
//			indexingNodes,
//			initialNOzonePrice,
//			sdk.ZeroInt(),
//			make([]Slashing, 0),
//		)
//
//		InitGenesis(ctx, keeper, registerGenesis)
//
//		return abci.ResponseInitChain{
//			Validators: validators,
//		}
//	}
//
//}
//
//// getEndBlocker returns a staking endblocker.
//func getEndBlocker(keeper Keeper) sdk.EndBlocker {
//	//return func(ctx sdk.Context, req abci.RequestEndBlock) abci.ResponseEndBlock {
//	//	validatorUpdates := keeper.StakingKeeper.BlockValidatorUpdates(ctx)
//	//
//	//	return abci.ResponseEndBlock{
//	//		ValidatorUpdates: validatorUpdates,
//	//	}
//	//}
//	return nil
//}
