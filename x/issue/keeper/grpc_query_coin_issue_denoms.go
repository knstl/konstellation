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

func (k Keeper) CoinIssueDenomsAll(c context.Context, req *types.QueryAllCoinIssueDenomsRequest) (*types.QueryAllCoinIssueDenomsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var coinIssueDenomss []*types.CoinIssueDenoms
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	coinIssueDenomsStore := prefix.NewStore(store, types.KeyPrefix(types.CoinIssueDenomsKey))

	pageRes, err := query.Paginate(coinIssueDenomsStore, req.Pagination, func(key []byte, value []byte) error {
		var coinIssueDenoms types.CoinIssueDenoms
		if err := k.cdc.UnmarshalBinaryBare(value, &coinIssueDenoms); err != nil {
			return err
		}

		coinIssueDenomss = append(coinIssueDenomss, &coinIssueDenoms)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllCoinIssueDenomsResponse{CoinIssueDenoms: coinIssueDenomss, Pagination: pageRes}, nil
}

func (k Keeper) CoinIssueDenoms(c context.Context, req *types.QueryGetCoinIssueDenomsRequest) (*types.QueryGetCoinIssueDenomsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var coinIssueDenoms types.CoinIssueDenoms
	ctx := sdk.UnwrapSDKContext(c)

	if !k.HasCoinIssueDenoms(ctx, req.Id) {
		return nil, sdkerrors.ErrKeyNotFound
	}

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CoinIssueDenomsKey))
	k.cdc.MustUnmarshalBinaryBare(store.Get(GetCoinIssueDenomsIDBytes(req.Id)), &coinIssueDenoms)

	return &types.QueryGetCoinIssueDenomsResponse{CoinIssueDenoms: &coinIssueDenoms}, nil
}
