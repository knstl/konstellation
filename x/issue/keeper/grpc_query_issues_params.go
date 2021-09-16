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

func (k Keeper) IssuesParamsAll(c context.Context, req *types.QueryAllIssuesParamsRequest) (*types.QueryAllIssuesParamsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var issuesParamss []*types.IssuesParams
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	issuesParamsStore := prefix.NewStore(store, types.KeyPrefix(types.IssuesParamsKey))

	pageRes, err := query.Paginate(issuesParamsStore, req.Pagination, func(key []byte, value []byte) error {
		var issuesParams types.IssuesParams
		if err := k.cdc.UnmarshalBinaryBare(value, &issuesParams); err != nil {
			return err
		}

		issuesParamss = append(issuesParamss, &issuesParams)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllIssuesParamsResponse{IssuesParams: issuesParamss, Pagination: pageRes}, nil
}

func (k Keeper) IssuesParams(c context.Context, req *types.QueryGetIssuesParamsRequest) (*types.QueryGetIssuesParamsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var issuesParams types.IssuesParams
	ctx := sdk.UnwrapSDKContext(c)

	if !k.HasIssuesParams(ctx, req.Id) {
		return nil, sdkerrors.ErrKeyNotFound
	}

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.IssuesParamsKey))
	k.cdc.MustUnmarshalBinaryBare(store.Get(GetIssuesParamsIDBytes(req.Id)), &issuesParams)

	return &types.QueryGetIssuesParamsResponse{IssuesParams: &issuesParams}, nil
}
