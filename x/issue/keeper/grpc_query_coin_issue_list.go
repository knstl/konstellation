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

func (k Keeper) CoinIssueListAll(c context.Context, req *types.QueryAllCoinIssueListRequest) (*types.QueryAllCoinIssueListResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var coinIssueLists []*types.CoinIssueList
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	coinIssueListStore := prefix.NewStore(store, types.KeyPrefix(types.CoinIssueListKey))

	pageRes, err := query.Paginate(coinIssueListStore, req.Pagination, func(key []byte, value []byte) error {
		var coinIssueList types.CoinIssueList
		if err := k.cdc.UnmarshalBinaryBare(value, &coinIssueList); err != nil {
			return err
		}

		coinIssueLists = append(coinIssueLists, &coinIssueList)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllCoinIssueListResponse{CoinIssueList: coinIssueLists, Pagination: pageRes}, nil
}

func (k Keeper) CoinIssueList(c context.Context, req *types.QueryGetCoinIssueListRequest) (*types.QueryGetCoinIssueListResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var coinIssueList types.CoinIssueList
	ctx := sdk.UnwrapSDKContext(c)

	if !k.HasCoinIssueList(ctx, req.Id) {
		return nil, sdkerrors.ErrKeyNotFound
	}

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CoinIssueListKey))
	k.cdc.MustUnmarshalBinaryBare(store.Get(GetCoinIssueListIDBytes(req.Id)), &coinIssueList)

	return &types.QueryGetCoinIssueListResponse{CoinIssueList: &coinIssueList}, nil
}
