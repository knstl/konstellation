package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/params"
	"github.com/konstellation/kn-sdk/x/issue/types"
)

// type declaration for parameters
func ParamKeyTable() params.KeyTable {
	return params.NewKeyTable().RegisterParamSet(&types.Params{})
}

func (k Keeper) IssueFee(ctx sdk.Context) (res sdk.Coin) {
	k.paramSubspace.Get(ctx, types.KeyIssueFee, &res)
	return
}
func (k Keeper) MintFee(ctx sdk.Context) (res sdk.Coin) {
	k.paramSubspace.Get(ctx, types.KeyMintFee, &res)
	return
}
func (k Keeper) BurnFee(ctx sdk.Context) (res sdk.Coin) {
	k.paramSubspace.Get(ctx, types.KeyBurnFee, &res)
	return
}
func (k Keeper) BurnFromFee(ctx sdk.Context) (res sdk.Coin) {
	k.paramSubspace.Get(ctx, types.KeyBurnFromFee, &res)
	return
}
func (k Keeper) FreezeFee(ctx sdk.Context) (res sdk.Coin) {
	k.paramSubspace.Get(ctx, types.KeyFreezeFee, &res)
	return
}
func (k Keeper) UnfreezeFee(ctx sdk.Context) (res sdk.Coin) {
	k.paramSubspace.Get(ctx, types.KeyUnFreezeFee, &res)
	return
}

func (k Keeper) TransferOwnerFee(ctx sdk.Context) (res sdk.Coin) {
	k.paramSubspace.Get(ctx, types.KeyTransferOwnerFee, &res)
	return
}

func (k Keeper) GetParams(ctx sdk.Context) (params types.Params) {
	k.paramSubspace.GetParamSet(ctx, &params)
	return
	//return types.NewParams(
	//	k.IssueFee(ctx),
	//	k.MintFee(ctx),
	//	k.BurnFee(ctx),
	//	k.BurnFromFee(ctx),
	//	k.FreezeFee(ctx),
	//	k.UnfreezeFee(ctx),
	//	k.TransferOwnerFee(ctx),
	//)
}

// set the params
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
	k.paramSubspace.SetParamSet(ctx, &params)
}
