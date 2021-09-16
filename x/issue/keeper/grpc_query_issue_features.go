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

func (k Keeper) IssueFeaturesAll(c context.Context, req *types.QueryAllIssueFeaturesRequest) (*types.QueryAllIssueFeaturesResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var issueFeaturess []*types.IssueFeatures
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	issueFeaturesStore := prefix.NewStore(store, types.KeyPrefix(types.IssueFeaturesKey))

	pageRes, err := query.Paginate(issueFeaturesStore, req.Pagination, func(key []byte, value []byte) error {
		var issueFeatures types.IssueFeatures
		if err := k.cdc.UnmarshalBinaryBare(value, &issueFeatures); err != nil {
			return err
		}

		issueFeaturess = append(issueFeaturess, &issueFeatures)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllIssueFeaturesResponse{IssueFeatures: issueFeaturess, Pagination: pageRes}, nil
}

func (k Keeper) IssueFeatures(c context.Context, req *types.QueryGetIssueFeaturesRequest) (*types.QueryGetIssueFeaturesResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var issueFeatures types.IssueFeatures
	ctx := sdk.UnwrapSDKContext(c)

	if !k.HasIssueFeatures(ctx, req.Id) {
		return nil, sdkerrors.ErrKeyNotFound
	}

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.IssueFeaturesKey))
	k.cdc.MustUnmarshalBinaryBare(store.Get(GetIssueFeaturesIDBytes(req.Id)), &issueFeatures)

	return &types.QueryGetIssueFeaturesResponse{IssueFeatures: &issueFeatures}, nil
}
