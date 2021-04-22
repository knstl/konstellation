package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/konstellation/konstellation/x/oracle/types"
)

var _ types.QueryServer = Keeper{}

// Params returns params of the mint module.
func (k Keeper) ExchangeRate(c context.Context, _ *types.QueryExchangeRateRequest) (*types.QueryExchangeRateResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	exchangeRate := k.GetExchangeRate(ctx)

	return &types.QueryExchangeRateResponse{ExchangeRate: exchangeRate}, nil
}
