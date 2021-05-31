package keeper

import (
	"context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/konstellation/konstellation/x/oracle/types"
)

var _ types.QueryServer = Keeper{}

// Params returns params of the mint module.
func (k Keeper) ExchangeRate(c context.Context, r *types.QueryExchangeRateRequest) (*types.QueryExchangeRateResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	exchangeRate, _ := k.GetExchangeRate(ctx, r.Pair)
	//return nil, status.Errorf(codes.NotFound, "validator %s not found", req.ValidatorAddr)

	return &types.QueryExchangeRateResponse{ExchangeRate: &exchangeRate}, nil
}

// Params returns params of the mint module.
func (k Keeper) AllExchangeRates(c context.Context, _ *types.QueryAllExchangeRatesRequest) (*types.QueryAllExchangeRatesResponse, error) {
	//ctx := sdk.UnwrapSDKContext(c)
	//exchangeRate := k.GetAllExchangeRates(ctx)

	return &types.QueryAllExchangeRatesResponse{}, nil
}
