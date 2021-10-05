package keeper

import (
	"context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/konstellation/konstellation/x/issue/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (q queryServer) FreezesAll(c context.Context, req *types.QueryAllFreezesRequest) (*types.QueryAllFreezesResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	freezes := q.keeper.GetFreezesOfDenom(ctx, req.Denom)

	return &types.QueryAllFreezesResponse{Freezes: freezes}, nil
}

func (q queryServer) Freezes(c context.Context, req *types.QueryFreezesRequest) (*types.QueryFreezesResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	holder, err := sdk.AccAddressFromBech32(req.Holder)
	if err != nil {
		return nil, err
	}

	ctx := sdk.UnwrapSDKContext(c)
	freezes := q.keeper.GetFreezesOfHolder(ctx, holder)

	return &types.QueryFreezesResponse{Freezes: freezes}, nil
}

func (q queryServer) Freeze(c context.Context, req *types.QueryFreezeRequest) (*types.QueryFreezeResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	holder, err := sdk.AccAddressFromBech32(req.Holder)
	if err != nil {
		return nil, err
	}

	ctx := sdk.UnwrapSDKContext(c)
	freeze := q.keeper.GetFreeze(ctx, req.Denom, holder)

	return &types.QueryFreezeResponse{Freeze: freeze}, nil
}
