package keeper

/*
func (k Keeper) ParamsAll(c context.Context, req *types.QueryAllParamsRequest) (*types.QueryAllParamsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var paramss []*types.Params
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	paramsStore := prefix.NewStore(store, types.KeyPrefix(types.ParamsKey))

	pageRes, err := query.Paginate(paramsStore, req.Pagination, func(key []byte, value []byte) error {
		var params types.Params
		if err := k.cdc.UnmarshalBinaryBare(value, &params); err != nil {
			return err
		}

		paramss = append(paramss, &params)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllParamsResponse{Params: paramss, Pagination: pageRes}, nil
}

func (k Keeper) Params(c context.Context, req *types.QueryGetParamsRequest) (*types.QueryGetParamsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var params types.Params
	ctx := sdk.UnwrapSDKContext(c)

	if !k.HasParams(ctx, req.Id) {
		return nil, sdkerrors.ErrKeyNotFound
	}

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ParamsKey))
	k.cdc.MustUnmarshalBinaryBare(store.Get(GetParamsIDBytes(req.Id)), &params)

	return &types.QueryGetParamsResponse{Params: &params}, nil
}
*/
