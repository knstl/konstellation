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

func (k Keeper) CoinsAll(c context.Context, req *types.QueryAllCoinsRequest) (*types.QueryAllCoinsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var coinss []*types.Coins
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	coinsStore := prefix.NewStore(store, types.KeyPrefix(types.CoinsKey))

	pageRes, err := query.Paginate(coinsStore, req.Pagination, func(key []byte, value []byte) error {
		var coins types.Coins
		if err := k.cdc.UnmarshalBinaryBare(value, &coins); err != nil {
			return err
		}

		coinss = append(coinss, &coins)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllCoinsResponse{Coins: coinss, Pagination: pageRes}, nil
}

func (k Keeper) Coins(c context.Context, req *types.QueryGetCoinsRequest) (*types.QueryGetCoinsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var coins types.Coins
	ctx := sdk.UnwrapSDKContext(c)

	if !k.HasCoins(ctx, req.Id) {
		return nil, sdkerrors.ErrKeyNotFound
	}

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CoinsKey))
	k.cdc.MustUnmarshalBinaryBare(store.Get(GetCoinsIDBytes(req.Id)), &coins)

	return &types.QueryGetCoinsResponse{Coins: &coins}, nil
}
