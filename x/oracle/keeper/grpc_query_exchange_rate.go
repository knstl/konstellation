package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/konstellation/konstellation/x/oracle/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) AllExchangeRates(c context.Context, req *types.QueryAllExchangeRatesRequest) (*types.QueryAllExchangeRatesResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var exchangeRates []*types.ExchangeRate
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	exchangeRateStore := prefix.NewStore(store, types.ExchangeRateKeyValue) //types.KeyPrefix(types.ExchangeRateKey))

	pageRes, err := query.Paginate(exchangeRateStore, req.Pagination, func(key []byte, value []byte) error {
		var exchangeRate types.ExchangeRate
		if err := k.cdc.UnmarshalBinaryBare(value, &exchangeRate); err != nil {
			return err
		}

		exchangeRates = append(exchangeRates, &exchangeRate)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllExchangeRatesResponse{ExchangeRate: exchangeRates, Pagination: pageRes}, nil
}

func (k Keeper) ExchangeRate(c context.Context, req *types.QueryExchangeRateRequest) (*types.QueryExchangeRateResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)

	ex, found := k.GetExchangeRate(ctx, req.Pair)
	if !found {
		return &types.QueryExchangeRateResponse{}, nil
	}

	return &types.QueryExchangeRateResponse{ExchangeRate: &ex}, nil
}
