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

func (k Keeper) AllowanceAll(c context.Context, req *types.QueryAllAllowanceRequest) (*types.QueryAllAllowanceResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var allowances []*types.Allowance
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	allowanceStore := prefix.NewStore(store, types.KeyPrefix(types.AllowanceKey))

	pageRes, err := query.Paginate(allowanceStore, req.Pagination, func(key []byte, value []byte) error {
		var allowance types.Allowance
		if err := k.cdc.UnmarshalBinaryBare(value, &allowance); err != nil {
			return err
		}

		allowances = append(allowances, &allowance)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllAllowanceResponse{Allowance: allowances, Pagination: pageRes}, nil
}

func (k Keeper) Allowance(c context.Context, req *types.QueryGetAllowanceRequest) (*types.QueryGetAllowanceResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var allowance types.Allowance
	ctx := sdk.UnwrapSDKContext(c)

	if !k.HasAllowance(ctx, req.Id) {
		return nil, sdkerrors.ErrKeyNotFound
	}

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AllowanceKey))
	k.cdc.MustUnmarshalBinaryBare(store.Get(GetAllowanceIDBytes(req.Id)), &allowance)

	return &types.QueryGetAllowanceResponse{Allowance: &allowance}, nil
}

/*
func (k *Keeper) Allowance(ctx sdk.Context, owner sdk.AccAddress, spender sdk.AccAddress, denom string) sdk.Coin {
	return k.allowance(ctx, owner, spender, denom)
}
*/
