package keeper

/*
// GetParamsCount get the total number of TypeName.LowerCamel
func (k Keeper) GetParamsCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ParamsCountKey))
	byteKey := types.KeyPrefix(types.ParamsCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	count, err := strconv.ParseUint(string(bz), 10, 64)
	if err != nil {
		// Panic because the count should be always formattable to uint64
		panic("cannot decode count")
	}

	return count
}

// SetParamsCount set the total number of params
func (k Keeper) SetParamsCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ParamsCountKey))
	byteKey := types.KeyPrefix(types.ParamsCountKey)
	bz := []byte(strconv.FormatUint(count, 10))
	store.Set(byteKey, bz)
}

// AppendParams appends a params in the store with a new id and update the count
func (k Keeper) AppendParams(
	ctx sdk.Context,
	params types.Params,
) uint64 {
	// Create the params
	count := k.GetParamsCount(ctx)

	// Set the ID of the appended value
	params.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ParamsKey))
	appendedValue := k.cdc.MustMarshalBinaryBare(&params)
	store.Set(GetParamsIDBytes(params.Id), appendedValue)

	// Update params count
	k.SetParamsCount(ctx, count+1)

	return count
}
*/
//
//// SetParams set a specific params in the store
//func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
//	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ParamsKey))
//	b := k.cdc.MustMarshalBinaryBare(&params)
//	store.Set(GetParamsIDBytes(params.Id), b)
//}
//
//// GetParams returns a params from its id
//func (k Keeper) GetParams(ctx sdk.Context, id uint64) types.Params {
//	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ParamsKey))
//	var params types.Params
//	k.cdc.MustUnmarshalBinaryBare(store.Get(GetParamsIDBytes(id)), &params)
//	return params
//}
//
//// HasParams checks if the params exists in the store
//func (k Keeper) HasParams(ctx sdk.Context, id uint64) bool {
//	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ParamsKey))
//	return store.Has(GetParamsIDBytes(id))
//}
//
//// GetParamsOwner returns the creator of the
//func (k Keeper) GetParamsOwner(ctx sdk.Context, id uint64) string {
//	return k.GetParams(ctx, id).Creator
//}
//
//// RemoveParams removes a params from the store
//func (k Keeper) RemoveParams(ctx sdk.Context, id uint64) {
//	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ParamsKey))
//	store.Delete(GetParamsIDBytes(id))
//}
//
//// GetAllParams returns all params
//func (k Keeper) GetAllParams(ctx sdk.Context) (list []types.Params) {
//	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ParamsKey))
//	iterator := sdk.KVStorePrefixIterator(store, []byte{})
//
//	defer iterator.Close()
//
//	for ; iterator.Valid(); iterator.Next() {
//		var val types.Params
//		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &val)
//		list = append(list, val)
//	}
//
//	return
//}
//
//// GetParamsIDBytes returns the byte representation of the ID
//func GetParamsIDBytes(id uint64) []byte {
//	bz := make([]byte, 8)
//	binary.BigEndian.PutUint64(bz, id)
//	return bz
//}
//
//// GetParamsIDFromBytes returns ID in uint64 format from a byte array
//func GetParamsIDFromBytes(bz []byte) uint64 {
//	return binary.BigEndian.Uint64(bz)
//}
