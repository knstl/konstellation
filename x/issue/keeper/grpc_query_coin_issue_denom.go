package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/konstellation/konstellation/x/issue/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) CoinIssueDenomAll(c context.Context, req *types.QueryAllCoinIssueDenomRequest) (*types.QueryAllCoinIssueDenomResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var coinIssueDenoms []*types.CoinIssueDenom
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	coinIssueDenomStore := prefix.NewStore(store, types.KeyPrefix(types.CoinIssueDenomKey))

	pageRes, err := query.Paginate(coinIssueDenomStore, req.Pagination, func(key []byte, value []byte) error {
		var coinIssueDenom types.CoinIssueDenom
		if err := k.cdc.UnmarshalBinaryBare(value, &coinIssueDenom); err != nil {
			return err
		}

		coinIssueDenoms = append(coinIssueDenoms, &coinIssueDenom)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllCoinIssueDenomResponse{CoinIssueDenom: coinIssueDenoms, Pagination: pageRes}, nil
}

func (k Keeper) CoinIssueDenom(c context.Context, req *types.QueryGetCoinIssueDenomRequest) (*types.QueryGetCoinIssueDenomResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var coinIssueDenom types.CoinIssueDenom
	ctx := sdk.UnwrapSDKContext(c)

	if !k.HasCoinIssueDenom(ctx, req.Id) {
		return nil, sdkerrors.ErrKeyNotFound
	}

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CoinIssueDenomKey))
	k.cdc.MustUnmarshalBinaryBare(store.Get(GetCoinIssueDenomIDBytes(req.Id)), &coinIssueDenom)

	return &types.QueryGetCoinIssueDenomResponse{CoinIssueDenom: &coinIssueDenom}, nil
}
