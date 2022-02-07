package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/konstellation/konstellation/x/oracle/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) AdminAddrAll(c context.Context, req *types.QueryAllAdminAddrRequest) (*types.QueryAllAdminAddrResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var adminAddrs []*types.AdminAddr
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	adminAddrStore := prefix.NewStore(store, types.AllowedAddressKey)

	pageRes, err := query.Paginate(adminAddrStore, req.Pagination, func(key []byte, value []byte) error {
		var adminAddr types.AdminAddr
		if err := k.cdc.Unmarshal(value, &adminAddr); err != nil {
			return err
		}

		adminAddrs = append(adminAddrs, &adminAddr)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllAdminAddrResponse{AdminAddr: adminAddrs, Pagination: pageRes}, nil
}

//
//func (k Keeper) AdminAddr(c context.Context, req *types.QueryGetAdminAddrRequest) (*types.QueryGetAdminAddrResponse, error) {
//	if req == nil {
//		return nil, status.Error(codes.InvalidArgument, "invalid request")
//	}
//
//	var adminAddr types.AdminAddr
//	ctx := sdk.UnwrapSDKContext(c)
//
//	if !k.HasAdminAddr(ctx, req.Id) {
//		return nil, sdkerrors.ErrKeyNotFound
//	}
//
//	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AdminAddrKey))
//	k.cdc. MustUnmarshal(store.Get(GetAdminAddrIDBytes(req.Id)), &adminAddr)
//
//	return &types.QueryGetAdminAddrResponse{AdminAddr: &adminAddr}, nil
//}
