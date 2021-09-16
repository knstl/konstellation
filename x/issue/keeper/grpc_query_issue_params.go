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

func (k Keeper) IssueParamsAll(c context.Context, req *types.QueryAllIssueParamsRequest) (*types.QueryAllIssueParamsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var issueParamss []*types.IssueParams
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	issueParamsStore := prefix.NewStore(store, types.KeyPrefix(types.IssueParamsKey))

	pageRes, err := query.Paginate(issueParamsStore, req.Pagination, func(key []byte, value []byte) error {
		var issueParams types.IssueParams
		if err := k.cdc.UnmarshalBinaryBare(value, &issueParams); err != nil {
			return err
		}

		issueParamss = append(issueParamss, &issueParams)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllIssueParamsResponse{IssueParams: issueParamss, Pagination: pageRes}, nil
}

func (k Keeper) IssueParams(c context.Context, req *types.QueryGetIssueParamsRequest) (*types.QueryGetIssueParamsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var issueParams types.IssueParams
	ctx := sdk.UnwrapSDKContext(c)

	if !k.HasIssueParams(ctx, req.Id) {
		return nil, sdkerrors.ErrKeyNotFound
	}

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.IssueParamsKey))
	k.cdc.MustUnmarshalBinaryBare(store.Get(GetIssueParamsIDBytes(req.Id)), &issueParams)

	return &types.QueryGetIssueParamsResponse{IssueParams: &issueParams}, nil
}
