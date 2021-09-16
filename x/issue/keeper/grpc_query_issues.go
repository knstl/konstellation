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

func (k Keeper) IssuesAll(c context.Context, req *types.QueryAllIssuesRequest) (*types.QueryAllIssuesResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var issuess []*types.Issues
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	issuesStore := prefix.NewStore(store, types.KeyPrefix(types.IssuesKey))

	pageRes, err := query.Paginate(issuesStore, req.Pagination, func(key []byte, value []byte) error {
		var issues types.Issues
		if err := k.cdc.UnmarshalBinaryBare(value, &issues); err != nil {
			return err
		}

		issuess = append(issuess, &issues)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllIssuesResponse{Issues: issuess, Pagination: pageRes}, nil
}

func (k Keeper) Issues(c context.Context, req *types.QueryGetIssuesRequest) (*types.QueryGetIssuesResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var issues types.Issues
	ctx := sdk.UnwrapSDKContext(c)

	if !k.HasIssues(ctx, req.Id) {
		return nil, sdkerrors.ErrKeyNotFound
	}

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.IssuesKey))
	k.cdc.MustUnmarshalBinaryBare(store.Get(GetIssuesIDBytes(req.Id)), &issues)

	return &types.QueryGetIssuesResponse{Issues: &issues}, nil
}
