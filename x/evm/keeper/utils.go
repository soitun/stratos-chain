package keeper

import (
	"math/big"

	"github.com/ethereum/go-ethereum/core"
	ethtypes "github.com/ethereum/go-ethereum/core/types"

	"cosmossdk.io/errors"
	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	authante "github.com/cosmos/cosmos-sdk/x/auth/ante"

	evmtypes "github.com/stratosnet/stratos-chain/x/evm/types"
)

// DeductTxCostsFromUserBalance it calculates the tx costs and deducts the fees
func (k Keeper) DeductTxCostsFromUserBalance(
	ctx sdk.Context,
	msgEthTx evmtypes.MsgEthereumTx,
	txData evmtypes.TxData,
	denom string,
	homestead, istanbul, london bool,
) (sdk.Coins, error) {
	isContractCreation := txData.GetTo() == nil

	// fetch sender account from signature
	signerAcc, err := authante.GetSignerAcc(ctx, k.accountKeeper, msgEthTx.GetFrom())
	if err != nil {
		return nil, errors.Wrapf(err, "account not found for sender %s", msgEthTx.From)
	}

	gasLimit := txData.GetGas()

	var accessList ethtypes.AccessList
	if txData.GetAccessList() != nil {
		accessList = txData.GetAccessList()
	}

	intrinsicGas, err := core.IntrinsicGas(txData.GetData(), accessList, isContractCreation, homestead, istanbul)
	if err != nil {
		return nil, errors.Wrapf(
			err,
			"failed to retrieve intrinsic gas, contract creation = %t; homestead = %t, istanbul = %t",
			isContractCreation, homestead, istanbul,
		)
	}

	// intrinsic gas verification during CheckTx
	if ctx.IsCheckTx() && gasLimit < intrinsicGas {
		return nil, errors.Wrapf(
			sdkerrors.ErrOutOfGas,
			"gas limit too low: %d (gas limit) < %d (intrinsic gas)", gasLimit, intrinsicGas,
		)
	}

	var feeAmt *big.Int

	feeMktParams := k.GetParams(ctx).FeeMarketParams
	if london && !feeMktParams.NoBaseFee && txData.TxType() == ethtypes.DynamicFeeTxType {
		baseFee := k.GetBaseFeeParam(ctx)
		if txData.GetGasFeeCap().Cmp(baseFee) < 0 {
			return nil, errors.Wrapf(sdkerrors.ErrInsufficientFee, "the tx gasFeeCap is lower than the tx baseFee: %s (gasFeeCap), %s (baseFee) ", txData.GetGasFeeCap(), baseFee)
		}
		feeAmt = txData.EffectiveFee(baseFee)
	} else {
		feeAmt = txData.Fee()
	}

	if feeAmt.Sign() == 0 {
		// zero fee, no need to deduct
		return sdk.NewCoins(), nil
	}

	fees := sdk.Coins{sdk.NewCoin(denom, sdkmath.NewIntFromBigInt(feeAmt))}

	// deduct the full gas cost from the user balance
	if err := authante.DeductFees(k.bankKeeper, ctx, signerAcc, fees); err != nil {
		return nil, errors.Wrapf(
			err,
			"failed to deduct full gas cost %s from the user %s balance",
			fees, msgEthTx.From,
		)
	}
	return fees, nil
}

// CheckSenderBalance validates that the tx cost value is positive and that the
// sender has enough funds to pay for the fees and value of the transaction.
func CheckSenderBalance(
	balance sdkmath.Int,
	txData evmtypes.TxData,
) error {
	cost := txData.Cost()

	if cost.Sign() < 0 {
		return errors.Wrapf(
			sdkerrors.ErrInvalidCoins,
			"tx cost (%s) is negative and invalid", cost,
		)
	}

	if balance.IsNegative() || balance.BigInt().Cmp(cost) < 0 {
		return errors.Wrapf(
			sdkerrors.ErrInsufficientFunds,
			"sender balance < tx cost (%s < %s)", balance, txData.Cost(),
		)
	}
	return nil
}
