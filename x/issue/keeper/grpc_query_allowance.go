package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/konstellation/konstellation/x/issue/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (q queryServer) Allowances(c context.Context, req *types.QueryAllowancesRequest) (*types.QueryAllowancesResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ownerAddress, err := sdk.AccAddressFromBech32(req.Owner)
	if err != nil {
		return nil, err
	}
	ctx := sdk.UnwrapSDKContext(c)

	allowances := q.keeper.Allowances(ctx, ownerAddress, req.Denom)

	//store := ctx.KVStore(k.storeKey)
	//allowanceStore := prefix.NewStore(store, types.KeyPrefix(types.AllowanceKey))
	//
	//pageRes, err := query.Paginate(allowanceStore, req.Pagination, func(key []byte, value []byte) error {
	//	var allowance types.Allowance
	//	if err := k.cdc.UnmarshalBinaryBare(value, &allowance); err != nil {
	//		return err
	//	}
	//
	//	allowances = append(allowances, &allowance)
	//	return nil
	//})
	//
	//if err != nil {
	//	return nil, status.Error(codes.Internal, err.Error())
	//}

	// todo add pagination
	return &types.QueryAllowancesResponse{Allowances: allowances}, nil
}

func (q queryServer) Allowance(c context.Context, req *types.QueryAllowanceRequest) (*types.QueryAllowanceResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ownerAddress, err := sdk.AccAddressFromBech32(req.Owner)
	if err != nil {
		return nil, err
	}
	spenderAddress, err := sdk.AccAddressFromBech32(req.Spender)
	if err != nil {
		return nil, err
	}

	ctx := sdk.UnwrapSDKContext(c)
	allowance := q.keeper.Allowance(ctx, ownerAddress, spenderAddress, req.Denom)

	return &types.QueryAllowanceResponse{Allowance: &allowance}, nil
}

/*
func (k *Keeper) Allowance(ctx sdk.Context, owner sdk.AccAddress, spender sdk.AccAddress, denom string) sdk.Coin {
	return k.allowance(ctx, owner, spender, denom)
}
*/
