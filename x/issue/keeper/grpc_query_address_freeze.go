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

func (k Keeper) AddressFreezeAll(c context.Context, req *types.QueryAllAddressFreezeRequest) (*types.QueryAllAddressFreezeResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var addressFreezes []*types.AddressFreeze
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	addressFreezeStore := prefix.NewStore(store, types.KeyPrefix(types.AddressFreezeKey))

	pageRes, err := query.Paginate(addressFreezeStore, req.Pagination, func(key []byte, value []byte) error {
		var addressFreeze types.AddressFreeze
		if err := k.cdc.UnmarshalBinaryBare(value, &addressFreeze); err != nil {
			return err
		}

		addressFreezes = append(addressFreezes, &addressFreeze)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllAddressFreezeResponse{AddressFreeze: addressFreezes, Pagination: pageRes}, nil
}

func (k Keeper) AddressFreeze(c context.Context, req *types.QueryGetAddressFreezeRequest) (*types.QueryGetAddressFreezeResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var addressFreeze types.AddressFreeze
	ctx := sdk.UnwrapSDKContext(c)

	if !k.HasAddressFreeze(ctx, req.Id) {
		return nil, sdkerrors.ErrKeyNotFound
	}

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AddressFreezeKey))
	k.cdc.MustUnmarshalBinaryBare(store.Get(GetAddressFreezeIDBytes(req.Id)), &addressFreeze)

	return &types.QueryGetAddressFreezeResponse{AddressFreeze: &addressFreeze}, nil
}
