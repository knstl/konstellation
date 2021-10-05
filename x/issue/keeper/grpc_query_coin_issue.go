package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/konstellation/konstellation/x/issue/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (q queryServer) Issue(c context.Context, req *types.QueryIssueRequest) (*types.QueryIssueResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	if req.Denom == "" {
		return nil, status.Error(codes.InvalidArgument, "denom cannot be empty")
	}

	ctx := sdk.UnwrapSDKContext(c)

	coinIssue, err := q.keeper.GetIssue(ctx, req.Denom)
	if err != nil {
		return nil, err
	}
	return &types.QueryIssueResponse{Issue: coinIssue}, nil
}
