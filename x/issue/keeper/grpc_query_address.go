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

func (k Keeper) AddressAll(c context.Context, req *types.QueryAllAddressRequest) (*types.QueryAllAddressResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var addresss []*types.Address
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	addressStore := prefix.NewStore(store, types.KeyPrefix(types.AddressKey))

	pageRes, err := query.Paginate(addressStore, req.Pagination, func(key []byte, value []byte) error {
		var address types.Address
		if err := k.cdc.UnmarshalBinaryBare(value, &address); err != nil {
			return err
		}

		addresss = append(addresss, &address)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllAddressResponse{Address: addresss, Pagination: pageRes}, nil
}

func (k Keeper) Address(c context.Context, req *types.QueryGetAddressRequest) (*types.QueryGetAddressResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var address types.Address
	ctx := sdk.UnwrapSDKContext(c)

	if !k.HasAddress(ctx, req.Id) {
		return nil, sdkerrors.ErrKeyNotFound
	}

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AddressKey))
	k.cdc.MustUnmarshalBinaryBare(store.Get(GetAddressIDBytes(req.Id)), &address)

	return &types.QueryGetAddressResponse{Address: &address}, nil
}
