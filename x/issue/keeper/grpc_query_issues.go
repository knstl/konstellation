package keeper

import (
	"context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/konstellation/konstellation/x/issue/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (q queryServer) IssuesAll(c context.Context, req *types.QueryAllIssuesRequest) (*types.QueryAllIssuesResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	issues := q.keeper.ListAll(ctx)

	// todo pageresponse
	return &types.QueryAllIssuesResponse{Issues: issues}, nil
}

func (q queryServer) Issues(c context.Context, req *types.QueryIssuesRequest) (*types.QueryIssuesResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)

	params := types.IssuesParams{
		Owner: req.Owner,
		Limit: req.Pagination.Limit,
	}

	issues := q.keeper.List(ctx, params)

	// todo pageresponse
	return &types.QueryIssuesResponse{Issues: issues}, nil
}
