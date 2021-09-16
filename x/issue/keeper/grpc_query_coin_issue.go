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

func (k Keeper) CoinIssueAll(c context.Context, req *types.QueryAllCoinIssueRequest) (*types.QueryAllCoinIssueResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var coinIssues []*types.CoinIssue
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	coinIssueStore := prefix.NewStore(store, types.KeyPrefix(types.CoinIssueKey))

	pageRes, err := query.Paginate(coinIssueStore, req.Pagination, func(key []byte, value []byte) error {
		var coinIssue types.CoinIssue
		if err := k.cdc.UnmarshalBinaryBare(value, &coinIssue); err != nil {
			return err
		}

		coinIssues = append(coinIssues, &coinIssue)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllCoinIssueResponse{CoinIssue: coinIssues, Pagination: pageRes}, nil
}

func (k Keeper) CoinIssue(c context.Context, req *types.QueryGetCoinIssueRequest) (*types.QueryGetCoinIssueResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var coinIssue types.CoinIssue
	ctx := sdk.UnwrapSDKContext(c)

	if !k.HasCoinIssue(ctx, req.Id) {
		return nil, sdkerrors.ErrKeyNotFound
	}

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CoinIssueKey))
	k.cdc.MustUnmarshalBinaryBare(store.Get(GetCoinIssueIDBytes(req.Id)), &coinIssue)

	return &types.QueryGetCoinIssueResponse{CoinIssue: &coinIssue}, nil
}
