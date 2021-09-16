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

func (k Keeper) AddressFreezeListAll(c context.Context, req *types.QueryAllAddressFreezeListRequest) (*types.QueryAllAddressFreezeListResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var addressFreezeLists []*types.AddressFreezeList
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	addressFreezeListStore := prefix.NewStore(store, types.KeyPrefix(types.AddressFreezeListKey))

	pageRes, err := query.Paginate(addressFreezeListStore, req.Pagination, func(key []byte, value []byte) error {
		var addressFreezeList types.AddressFreezeList
		if err := k.cdc.UnmarshalBinaryBare(value, &addressFreezeList); err != nil {
			return err
		}

		addressFreezeLists = append(addressFreezeLists, &addressFreezeList)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllAddressFreezeListResponse{AddressFreezeList: addressFreezeLists, Pagination: pageRes}, nil
}

func (k Keeper) AddressFreezeList(c context.Context, req *types.QueryGetAddressFreezeListRequest) (*types.QueryGetAddressFreezeListResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var addressFreezeList types.AddressFreezeList
	ctx := sdk.UnwrapSDKContext(c)

	if !k.HasAddressFreezeList(ctx, req.Id) {
		return nil, sdkerrors.ErrKeyNotFound
	}

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AddressFreezeListKey))
	k.cdc.MustUnmarshalBinaryBare(store.Get(GetAddressFreezeListIDBytes(req.Id)), &addressFreezeList)

	return &types.QueryGetAddressFreezeListResponse{AddressFreezeList: &addressFreezeList}, nil
}
