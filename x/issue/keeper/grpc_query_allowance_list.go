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

func (k Keeper) AllowanceListAll(c context.Context, req *types.QueryAllAllowanceListRequest) (*types.QueryAllAllowanceListResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var allowanceLists []*types.AllowanceList
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	allowanceListStore := prefix.NewStore(store, types.KeyPrefix(types.AllowanceListKey))

	pageRes, err := query.Paginate(allowanceListStore, req.Pagination, func(key []byte, value []byte) error {
		var allowanceList types.AllowanceList
		if err := k.cdc.UnmarshalBinaryBare(value, &allowanceList); err != nil {
			return err
		}

		allowanceLists = append(allowanceLists, &allowanceList)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllAllowanceListResponse{AllowanceList: allowanceLists, Pagination: pageRes}, nil
}

func (k Keeper) AllowanceList(c context.Context, req *types.QueryGetAllowanceListRequest) (*types.QueryGetAllowanceListResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var allowanceList types.AllowanceList
	ctx := sdk.UnwrapSDKContext(c)

	if !k.HasAllowanceList(ctx, req.Id) {
		return nil, sdkerrors.ErrKeyNotFound
	}

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AllowanceListKey))
	k.cdc.MustUnmarshalBinaryBare(store.Get(GetAllowanceListIDBytes(req.Id)), &allowanceList)

	return &types.QueryGetAllowanceListResponse{AllowanceList: &allowanceList}, nil
}
