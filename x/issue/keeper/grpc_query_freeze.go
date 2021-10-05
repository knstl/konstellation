package keeper

//func (k Keeper) FreezeAll(c context.Context, req *types.QueryAllFreezeRequest) (*types.QueryAllFreezeResponse, error) {
//	if req == nil {
//		return nil, status.Error(codes.InvalidArgument, "invalid request")
//	}
//
//	var freezes []*types.Freeze
//	ctx := sdk.UnwrapSDKContext(c)
//
//	store := ctx.KVStore(k.storeKey)
//	freezeStore := prefix.NewStore(store, types.KeyPrefix(types.FreezeKey))
//
//	pageRes, err := query.Paginate(freezeStore, req.Pagination, func(key []byte, value []byte) error {
//		var freeze types.Freeze
//		if err := k.cdc.UnmarshalBinaryBare(value, &freeze); err != nil {
//			return err
//		}
//
//		freezes = append(freezes, &freeze)
//		return nil
//	})
//
//	if err != nil {
//		return nil, status.Error(codes.Internal, err.Error())
//	}
//
//	return &types.QueryAllFreezeResponse{Freeze: freezes, Pagination: pageRes}, nil
//}
//
//func (k Keeper) Freeze(c context.Context, req *types.QueryGetFreezeRequest) (*types.QueryGetFreezeResponse, error) {
//	if req == nil {
//		return nil, status.Error(codes.InvalidArgument, "invalid request")
//	}
//
//	var freeze types.Freeze
//	ctx := sdk.UnwrapSDKContext(c)
//
//	if !k.HasFreeze(ctx, req.Id) {
//		return nil, sdkerrors.ErrKeyNotFound
//	}
//
//	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.FreezeKey))
//	k.cdc.MustUnmarshalBinaryBare(store.Get(GetFreezeIDBytes(req.Id)), &freeze)
//
//	return &types.QueryGetFreezeResponse{Freeze: &freeze}, nil
//}

/*
func (k *Keeper) Freeze(ctx sdk.Context, freezer, holder sdk.AccAddress, denom, op string) *sdkerrors.Error {
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeFreeze,
			sdk.NewAttribute(types.AttributeKeyFreezer, freezer.String()),
			sdk.NewAttribute(types.AttributeKeyHolder, holder.String()),
			sdk.NewAttribute(types.AttributeKeyDenom, denom),
			sdk.NewAttribute(types.AttributeKeyOp, op),
		),
	)

	issue, err := k.getIssueIfOwner(ctx, denom, freezer)
	if err != nil {
		return err
	}
	if issue.FreezeDisabled {
		return types.ErrCanNotFreeze(denom)
	}

	return k.freeze(ctx, holder, denom, op, true)
}
*/
