package keeper

import (
	db "github.com/cometbft/cometbft-db"

	sdkmath "cosmossdk.io/math"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	stratos "github.com/stratosnet/stratos-chain/types"
	"github.com/stratosnet/stratos-chain/x/pot/types"
)

func (k Keeper) SetTotalMinedTokens(ctx sdk.Context, totalMinedToken sdk.Coin) {
	store := ctx.KVStore(k.storeKey)
	b := k.cdc.MustMarshalLengthPrefixed(&totalMinedToken)
	store.Set(types.TotalMinedTokensKeyPrefix, b)
}

func (k Keeper) GetTotalMinedTokens(ctx sdk.Context) (totalMinedToken sdk.Coin) {
	store := ctx.KVStore(k.storeKey)
	b := store.Get(types.TotalMinedTokensKeyPrefix)
	if b == nil {
		return sdk.NewCoin(k.RewardDenom(ctx), sdkmath.ZeroInt())
	}
	k.cdc.MustUnmarshalLengthPrefixed(b, &totalMinedToken)
	return
}

func (k Keeper) SetLastDistributedEpoch(ctx sdk.Context, epoch sdkmath.Int) {
	store := ctx.KVStore(k.storeKey)
	b := k.cdc.MustMarshalLengthPrefixed(&stratos.Int{Value: &epoch})
	store.Set(types.LastDistributedEpochKeyPrefix, b)
}

func (k Keeper) GetLastDistributedEpoch(ctx sdk.Context) (epoch sdkmath.Int) {
	store := ctx.KVStore(k.storeKey)
	b := store.Get(types.LastDistributedEpochKeyPrefix)
	if b == nil {
		return sdkmath.ZeroInt()
	}
	intValue := stratos.Int{}
	k.cdc.MustUnmarshalLengthPrefixed(b, &intValue)
	epoch = *intValue.Value
	return
}

func (k Keeper) SetIndividualReward(ctx sdk.Context, walletAddress sdk.AccAddress, matureEpoch sdkmath.Int, value types.Reward) {
	store := ctx.KVStore(k.storeKey)
	b := k.cdc.MustMarshalLengthPrefixed(&value)
	store.Set(types.GetIndividualRewardKey(walletAddress, matureEpoch), b)
}

func (k Keeper) GetIndividualReward(ctx sdk.Context, walletAddress sdk.AccAddress, matureEpoch sdkmath.Int) (value types.Reward, found bool) {
	store := ctx.KVStore(k.storeKey)
	b := store.Get(types.GetIndividualRewardKey(walletAddress, matureEpoch))
	if b == nil {
		return value, false
	}
	k.cdc.MustUnmarshalLengthPrefixed(b, &value)
	return value, true
}

func (k Keeper) RemoveIndividualReward(ctx sdk.Context, individualRewardKey []byte) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(individualRewardKey)
}

// IteratorIndividualReward Iteration for getting individual reward of each owner at a specific epoch
func (k Keeper) IteratorIndividualReward(ctx sdk.Context, epoch sdkmath.Int, handler func(walletAddress sdk.AccAddress, individualReward types.Reward) (stop bool)) {
	store := ctx.KVStore(k.storeKey)
	iter := sdk.KVStorePrefixIterator(store, types.GetIndividualRewardIteratorKey(epoch))
	defer iter.Close()
	for ; iter.Valid(); iter.Next() {
		addr := sdk.AccAddress(iter.Key()[len(types.GetIndividualRewardIteratorKey(epoch)):])

		individualReward := types.Reward{}
		k.cdc.MustUnmarshalLengthPrefixed(iter.Value(), &individualReward)
		if handler(addr, individualReward) {
			break
		}
	}
}

func (k Keeper) SetMatureTotalReward(ctx sdk.Context, walletAddress sdk.AccAddress, value sdk.Coins) {
	store := ctx.KVStore(k.storeKey)
	b := k.cdc.MustMarshalLengthPrefixed(&stratos.Coins{Value: value})
	store.Set(types.GetMatureTotalRewardKey(walletAddress), b)
}

func (k Keeper) GetMatureTotalReward(ctx sdk.Context, walletAddress sdk.AccAddress) (value sdk.Coins) {
	store := ctx.KVStore(k.storeKey)
	b := store.Get(types.GetMatureTotalRewardKey(walletAddress))
	if b == nil {
		return sdk.Coins{}
	}
	coinsValue := stratos.Coins{}
	k.cdc.MustUnmarshalLengthPrefixed(b, &coinsValue)
	value = coinsValue.GetValue()
	return
}

// IteratorMatureTotal Iteration for getting total mature reward
func (k Keeper) IteratorMatureTotal(ctx sdk.Context, handler func(walletAddress sdk.AccAddress, matureTotal sdk.Coins) (stop bool)) {
	store := ctx.KVStore(k.storeKey)
	iter := sdk.KVStorePrefixIterator(store, types.MatureTotalRewardKeyPrefix)
	defer iter.Close()
	for ; iter.Valid(); iter.Next() {
		addr := sdk.AccAddress(iter.Key()[len(types.MatureTotalRewardKeyPrefix):])
		coinsValue := stratos.Coins{}
		k.cdc.MustUnmarshalLengthPrefixed(iter.Value(), &coinsValue)
		matureTotal := coinsValue.Value
		if handler(addr, matureTotal) {
			break
		}
	}
}

func (k Keeper) SetImmatureTotalReward(ctx sdk.Context, walletAddress sdk.AccAddress, value sdk.Coins) {
	store := ctx.KVStore(k.storeKey)
	b := k.cdc.MustMarshalLengthPrefixed(&stratos.Coins{Value: value})
	store.Set(types.GetImmatureTotalRewardKey(walletAddress), b)
}

func (k Keeper) GetImmatureTotalReward(ctx sdk.Context, walletAddress sdk.AccAddress) (value sdk.Coins) {
	store := ctx.KVStore(k.storeKey)
	b := store.Get(types.GetImmatureTotalRewardKey(walletAddress))
	if b == nil {
		return sdk.Coins{}
	}
	coinsValue := stratos.Coins{}
	k.cdc.MustUnmarshalLengthPrefixed(b, &coinsValue)
	value = coinsValue.GetValue()
	return
}

// IteratorImmatureTotal Iteration for getting total immature reward
func (k Keeper) IteratorImmatureTotal(ctx sdk.Context, handler func(walletAddress sdk.AccAddress, immatureTotal sdk.Coins) (stop bool)) {
	store := ctx.KVStore(k.storeKey)
	iter := sdk.KVStorePrefixIterator(store, types.ImmatureTotalRewardKeyPrefix)
	defer iter.Close()
	for ; iter.Valid(); iter.Next() {
		addr := sdk.AccAddress(iter.Key()[len(types.ImmatureTotalRewardKeyPrefix):])
		coinsValue := stratos.Coins{}
		k.cdc.MustUnmarshalLengthPrefixed(iter.Value(), &coinsValue)
		immatureTotal := coinsValue.Value
		if handler(addr, immatureTotal) {
			break
		}
	}
}

func (k Keeper) GetVolumeReport(ctx sdk.Context, epoch sdkmath.Int) (res types.VolumeReportRecord) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.VolumeReportStoreKey(epoch))
	if bz == nil {
		return types.VolumeReportRecord{}
	}
	k.cdc.MustUnmarshalLengthPrefixed(bz, &res)
	return res
}

func (k Keeper) SetVolumeReport(ctx sdk.Context, epoch sdkmath.Int, reportRecord types.VolumeReportRecord) {
	store := ctx.KVStore(k.storeKey)
	storeKey := types.VolumeReportStoreKey(epoch)
	bz := k.cdc.MustMarshalLengthPrefixed(&reportRecord)
	store.Set(storeKey, bz)
}

func (k Keeper) GetMaturedEpoch(ctx sdk.Context) (epoch sdkmath.Int) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.MaturedEpochKeyPrefix)
	if bz == nil {
		return sdkmath.ZeroInt()
	}
	intValue := stratos.Int{}
	k.cdc.MustUnmarshalLengthPrefixed(bz, &intValue)
	epoch = *intValue.Value
	return
}

func (k Keeper) SetMaturedEpoch(ctx sdk.Context, epoch sdkmath.Int) {
	store := ctx.KVStore(k.storeKey)
	b := k.cdc.MustMarshalLengthPrefixed(&stratos.Int{Value: &epoch})
	store.Set(types.MaturedEpochKeyPrefix, b)
}

func GetIterator(prefixStore storetypes.KVStore, start []byte, reverse bool) db.Iterator {
	if reverse {
		var end []byte
		if start != nil {
			itr := prefixStore.Iterator(start, nil)
			defer itr.Close()
			if itr.Valid() {
				itr.Next()
				end = itr.Key()
			}
		}
		return prefixStore.ReverseIterator(nil, end)
	}
	return prefixStore.Iterator(start, nil)
}

func (k Keeper) SetTotalReward(ctx sdk.Context, epoch sdkmath.Int, totalReward types.TotalReward) {
	store := ctx.KVStore(k.storeKey)
	storeKey := types.GetTotalRewardKey(epoch)
	bz := k.cdc.MustMarshalLengthPrefixed(&totalReward)
	store.Set(storeKey, bz)
}

func (k Keeper) GetTotalReward(ctx sdk.Context, epoch sdkmath.Int) (totalReward types.TotalReward) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.GetTotalRewardKey(epoch))
	if bz == nil {
		return types.TotalReward{}
	}
	k.cdc.MustUnmarshalLengthPrefixed(bz, &totalReward)
	return
}

// IteratorTotalReward Iteration for getting total mature reward
func (k Keeper) IteratorTotalReward(ctx sdk.Context, handler func(epoch sdkmath.Int, totalReward types.TotalReward) (stop bool)) {
	store := ctx.KVStore(k.storeKey)
	iter := sdk.KVStorePrefixIterator(store, types.TotalRewardKeyPrefix)
	defer iter.Close()
	for ; iter.Valid(); iter.Next() {
		epochBytes := iter.Key()[len(types.TotalRewardKeyPrefix):]
		epoch, _ := sdkmath.NewIntFromString(string(epochBytes))
		var totalReward types.TotalReward
		k.cdc.MustUnmarshalLengthPrefixed(iter.Value(), &totalReward)
		if handler(epoch, totalReward) {
			break
		}
	}
}
