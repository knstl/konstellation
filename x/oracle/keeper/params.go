package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/user/nameservice/x/nameservice/types"
	"time"
)

// GetParams returns the total set of nameservice parameters.
func (k Keeper) GetParams(ctx sdk.Context) (params types.Params) {
	k.paramspace.GetParamSet(ctx, &params)
	return params
}

// SetParams sets the nameservice parameters to the param space.
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
	k.paramspace.SetParamSet(ctx, &params)
}
