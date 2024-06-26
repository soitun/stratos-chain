package keeper

import (
	"fmt"
	"time"

	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"

	stratos "github.com/stratosnet/stratos-chain/types"
	"github.com/stratosnet/stratos-chain/x/register/types"
)

// deprecated
func (k Keeper) SetInitialGenesisDepositTotal(ctx sdk.Context, deposit sdkmath.Int) {
	store := ctx.KVStore(k.storeKey)
	b := k.cdc.MustMarshalLengthPrefixed(&stratos.Int{Value: &deposit})
	store.Set(types.InitialGenesisDepositTotalKey, b)
}

// deprecated
func (k Keeper) GetInitialGenesisDepositTotal(ctx sdk.Context) (deposit sdkmath.Int) {
	store := ctx.KVStore(k.storeKey)
	b := store.Get(types.InitialGenesisDepositTotalKey)
	if b == nil {
		return sdkmath.ZeroInt()
	}
	value := stratos.Int{}
	k.cdc.MustUnmarshalLengthPrefixed(b, &value)
	deposit = *value.Value
	return
}

func (k Keeper) SetRemainingOzoneLimit(ctx sdk.Context, value sdkmath.Int) {
	store := ctx.KVStore(k.storeKey)
	b := k.cdc.MustMarshalLengthPrefixed(&stratos.Int{Value: &value})
	store.Set(types.UpperBoundOfTotalOzoneKey, b)
}

func (k Keeper) GetRemainingOzoneLimit(ctx sdk.Context) (value sdkmath.Int) {
	store := ctx.KVStore(k.storeKey)
	b := store.Get(types.UpperBoundOfTotalOzoneKey)
	if b == nil {
		return sdkmath.ZeroInt()
	}
	intVal := stratos.Int{}
	k.cdc.MustUnmarshalLengthPrefixed(b, &intVal)
	value = *intVal.Value
	return
}

func (k *Keeper) GetPrepayParams(ctx sdk.Context) (St, Pt, Lt sdkmath.Int) {
	St = k.GetEffectiveTotalDeposit(ctx)
	Pt = k.GetTotalUnissuedPrepay(ctx).Amount
	Lt = k.GetRemainingOzoneLimit(ctx)
	return
}

func (k *Keeper) CalculatePurchaseAmount(ctx sdk.Context, amount sdkmath.Int) (sdkmath.Int, sdkmath.Int, error) {
	St, Pt, Lt := k.GetPrepayParams(ctx)

	purchase := Lt.ToLegacyDec().
		Mul(amount.ToLegacyDec()).
		Quo((St.
			Add(Pt).
			Add(amount)).ToLegacyDec()).
		TruncateInt()
	if purchase.GT(Lt) {
		return sdkmath.NewInt(0), sdkmath.NewInt(0), fmt.Errorf("not enough remaining ozone limit to complete prepay")
	}
	remaining := Lt.Sub(purchase)

	return purchase, remaining, nil
}

func (k Keeper) IsUnbondable(ctx sdk.Context, unbondAmt sdkmath.Int) bool {
	remaining := k.GetRemainingOzoneLimit(ctx)
	depositNozRate := k.GetDepositNozRate(ctx)
	return remaining.ToLegacyDec().GTE(unbondAmt.ToLegacyDec().Quo(depositNozRate))
}

// SetUnbondingNode sets the unbonding node
func (k Keeper) SetUnbondingNode(ctx sdk.Context, ubd types.UnbondingNode) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshalLengthPrefixed(&ubd)
	networkAddr, err := stratos.SdsAddressFromBech32(ubd.GetNetworkAddr())
	if err != nil {
		return
	}
	key := types.GetUBDNodeKey(networkAddr)
	store.Set(key, bz)
}

// RemoveUnbondingNode removes the unbonding node object
func (k Keeper) RemoveUnbondingNode(ctx sdk.Context, networkAddr stratos.SdsAddress) {
	store := ctx.KVStore(k.storeKey)
	key := types.GetUBDNodeKey(networkAddr)
	store.Delete(key)
}

// GetUnbondingNode return a unbonding node
func (k Keeper) GetUnbondingNode(ctx sdk.Context, networkAddr stratos.SdsAddress) (ubd types.UnbondingNode, found bool) {
	store := ctx.KVStore(k.storeKey)
	key := types.GetUBDNodeKey(networkAddr)
	value := store.Get(key)
	if value == nil {
		return ubd, false
	}
	k.cdc.MustUnmarshalLengthPrefixed(value, &ubd)
	return ubd, true
}

func (k Keeper) GetAllUnbondingNode(ctx sdk.Context) (unbondingNodes []types.UnbondingNode) {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.UBDNodeKey)
	defer iterator.Close()

	unbondingNodes = make([]types.UnbondingNode, 0)
	for ; iterator.Valid(); iterator.Next() {
		node := types.UnbondingNode{}
		k.cdc.MustUnmarshalLengthPrefixed(iterator.Value(), &node)
		unbondingNodes = append(unbondingNodes, node)
	}
	return
}

// SetUnbondingNodeQueueTimeSlice sets a specific unbonding queue timeslice.
func (k Keeper) SetUnbondingNodeQueueTimeSlice(ctx sdk.Context, timestamp time.Time, networkAddrs []string) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshalLengthPrefixed(&stratos.SdsAddresses{Addresses: networkAddrs})
	store.Set(types.GetUBDTimeKey(timestamp), bz)
}

// GetUnbondingNodeQueueTimeSlice gets a specific unbonding queue timeslice. A timeslice is a slice of DVPairs
// corresponding to unbonding delegations that expire at a certain time.
func (k Keeper) GetUnbondingNodeQueueTimeSlice(ctx sdk.Context, timestamp time.Time) (networkAddrs []string) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.GetUBDTimeKey(timestamp))
	if bz == nil {
		return make([]string, 0)
	}

	addrValue := stratos.SdsAddresses{}
	k.cdc.MustUnmarshalLengthPrefixed(bz, &addrValue)
	networkAddrs = addrValue.GetAddresses()
	return networkAddrs
}

// UnbondingNodeQueueIterator returns all the unbonding queue timeslices from time 0 until endTime
func (k Keeper) UnbondingNodeQueueIterator(ctx sdk.Context, endTime time.Time) sdk.Iterator {
	store := ctx.KVStore(k.storeKey)
	return store.Iterator(types.UBDNodeQueueKey, sdk.InclusiveEndBytes(types.GetUBDTimeKey(endTime)))
}

func (k Keeper) SetBondedResourceNodeCnt(ctx sdk.Context, count sdkmath.Int) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshalLengthPrefixed(&stratos.Int{Value: &count})
	store.Set(types.ResourceNodeCntKey, bz)
}

func (k Keeper) GetBondedResourceNodeCnt(ctx sdk.Context) (balance sdkmath.Int) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.ResourceNodeCntKey)
	if bz == nil {
		return sdkmath.ZeroInt()
	}
	intValue := stratos.Int{}
	k.cdc.MustUnmarshalLengthPrefixed(bz, &intValue)
	balance = *intValue.Value
	return
}

func (k Keeper) SetBondedMetaNodeCnt(ctx sdk.Context, count sdkmath.Int) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshalLengthPrefixed(&stratos.Int{Value: &count})
	store.Set(types.MetaNodeCntKey, bz)
}

func (k Keeper) GetBondedMetaNodeCnt(ctx sdk.Context) (balance sdkmath.Int) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.MetaNodeCntKey)
	if bz == nil {
		return sdkmath.ZeroInt()
	}
	intValue := stratos.Int{}
	k.cdc.MustUnmarshalLengthPrefixed(bz, &intValue)
	balance = *intValue.Value
	return
}

func (k Keeper) DeleteMetaNodeRegistrationVotePool(ctx sdk.Context, nodeAddr stratos.SdsAddress) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.GetMetaNodeRegistrationVotesKey(nodeAddr))
}

func (k Keeper) SetMetaNodeRegistrationVotePool(ctx sdk.Context, votePool types.MetaNodeRegistrationVotePool) {
	nodeAddr := votePool.GetNetworkAddress()
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshalLengthPrefixed(&votePool)
	node, _ := stratos.SdsAddressFromBech32(nodeAddr)
	store.Set(types.GetMetaNodeRegistrationVotesKey(node), bz)
}
func (k Keeper) GetMetaNodeRegistrationVotePool(ctx sdk.Context, nodeAddr stratos.SdsAddress) (votePool types.MetaNodeRegistrationVotePool, found bool) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.GetMetaNodeRegistrationVotesKey(nodeAddr))
	if bz == nil {
		return votePool, false
	}
	k.cdc.MustUnmarshalLengthPrefixed(bz, &votePool)
	return votePool, true
}

func (k Keeper) GetAllMetaNodeRegVotePool(ctx sdk.Context) (votePools []types.MetaNodeRegistrationVotePool) {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.MetaNodeRegistrationVotesKey)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		voteInfo := types.MetaNodeRegistrationVotePool{}
		k.cdc.MustUnmarshalLengthPrefixed(iterator.Value(), &voteInfo)
		votePools = append(votePools, voteInfo)
	}
	return
}

func (k Keeper) SetKickMetaNodeVotePool(ctx sdk.Context, votePool types.KickMetaNodeVotePool) {
	targetNetworkAddr := votePool.GetTargetNetworkAddress()
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshalLengthPrefixed(&votePool)
	node, _ := stratos.SdsAddressFromBech32(targetNetworkAddr)
	store.Set(types.GetKickMetaNodeVotesKey(node), bz)
}

func (k Keeper) GetKickMetaNodeVotePool(ctx sdk.Context, targetNetworkAddr stratos.SdsAddress) (votePool types.KickMetaNodeVotePool, found bool) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.GetKickMetaNodeVotesKey(targetNetworkAddr))
	if bz == nil {
		return votePool, false
	}
	k.cdc.MustUnmarshalLengthPrefixed(bz, &votePool)
	return votePool, true
}

func (k Keeper) GetAllKickMetaNodeVotePool(ctx sdk.Context) (votePools []types.KickMetaNodeVotePool) {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.KickMetaNodeVotesKey)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		voteInfo := types.KickMetaNodeVotePool{}
		k.cdc.MustUnmarshalLengthPrefixed(iterator.Value(), &voteInfo)
		votePools = append(votePools, voteInfo)
	}
	return
}

func (k Keeper) GetAllExpiredKickMetaNodeVotePool(ctx sdk.Context) (votePools []types.KickMetaNodeVotePool) {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.KickMetaNodeVotesKey)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		votePool := types.KickMetaNodeVotePool{}
		k.cdc.MustUnmarshalLengthPrefixed(iterator.Value(), &votePool)
		if votePool.GetExpireTime().Before(ctx.BlockTime()) {
			votePools = append(votePools, votePool)
		}
	}
	return
}

func (k Keeper) DeleteKickMetaNodeVotePool(ctx sdk.Context, targetNetworkAddr stratos.SdsAddress) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.GetKickMetaNodeVotesKey(targetNetworkAddr))
}

func (k Keeper) SetEffectiveTotalDeposit(ctx sdk.Context, deposit sdkmath.Int) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshalLengthPrefixed(&stratos.Int{Value: &deposit})
	store.Set(types.EffectiveGenesisDepositTotalKey, bz)
}

func (k Keeper) GetEffectiveTotalDeposit(ctx sdk.Context) (deposit sdkmath.Int) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.EffectiveGenesisDepositTotalKey)
	if bz == nil {
		return sdkmath.ZeroInt()
	}
	intValue := stratos.Int{}
	k.cdc.MustUnmarshalLengthPrefixed(bz, &intValue)
	deposit = *intValue.Value
	return
}

func (k Keeper) SetDepositNozRate(ctx sdk.Context, depositNozRate sdkmath.LegacyDec) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshalLengthPrefixed(&stratos.Dec{Value: &depositNozRate})
	store.Set(types.DepositNozRateKey, bz)
}

func (k Keeper) GetDepositNozRate(ctx sdk.Context) (depositNozRate sdkmath.LegacyDec) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.DepositNozRateKey)
	if bz == nil {
		panic("Stored deposit noz rate should not be nil")
	}
	decValue := stratos.Dec{}
	k.cdc.MustUnmarshalLengthPrefixed(bz, &decValue)
	depositNozRate = *decValue.Value
	return
}

// IteratorSlashingInfo Iteration for each slashing
func (k Keeper) IteratorSlashingInfo(ctx sdk.Context, handler func(walletAddress sdk.AccAddress, slashing sdkmath.Int) (stop bool)) {
	store := ctx.KVStore(k.storeKey)
	iter := sdk.KVStorePrefixIterator(store, types.SlashingPrefix)
	defer iter.Close()
	for ; iter.Valid(); iter.Next() {
		walletAddress := sdk.AccAddress(iter.Key()[len(types.SlashingPrefix):])
		intValue := stratos.Int{}
		k.cdc.MustUnmarshalLengthPrefixed(iter.Value(), &intValue)
		slashing := *intValue.Value
		if handler(walletAddress, slashing) {
			break
		}
	}
}

func (k Keeper) SetSlashing(ctx sdk.Context, walletAddress sdk.AccAddress, slashing sdkmath.Int) {
	store := ctx.KVStore(k.storeKey)
	storeKey := types.GetSlashingKey(walletAddress)
	bz := k.cdc.MustMarshalLengthPrefixed(&stratos.Int{Value: &slashing})
	store.Set(storeKey, bz)
}

func (k Keeper) GetSlashing(ctx sdk.Context, walletAddress sdk.AccAddress) (res sdkmath.Int) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.GetSlashingKey(walletAddress))
	if bz == nil {
		return sdkmath.ZeroInt()
	}
	intValue := stratos.Int{}
	k.cdc.MustUnmarshalLengthPrefixed(bz, &intValue)
	res = *intValue.Value
	return
}
