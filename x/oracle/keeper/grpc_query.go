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

	return &types.QueryExchangeRateResponse{ExchangeRate: &exchangeRate}, nil
}

// Params returns params of the mint module.
func (k Keeper) AllExchangeRates(c context.Context, _ *types.QueryAllExchangeRatesRequest) (*types.QueryAllExchangeRatesResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	rates := k.GetAllExchangeRates(ctx)

	// cast to pointer, protobuf requires
	exrates := make([]*types.ExchangeRate, len(rates))
	for _, r := range rates {
		rr := r
		exrates = append(exrates, &rr)
	}

	return &types.QueryAllExchangeRatesResponse{Pairs: exrates}, nil
}
