package keeper

import (
	"encoding/binary"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/konstellation/konstellation/x/issue/types"
	//"strconv"
)

/*
// GetAllowanceListCount get the total number of TypeName.LowerCamel
func (k Keeper) GetAllowanceListCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AllowanceListCountKey))
	byteKey := types.KeyPrefix(types.AllowanceListCountKey)
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

// SetAllowanceListCount set the total number of allowanceList
func (k Keeper) SetAllowanceListCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AllowanceListCountKey))
	byteKey := types.KeyPrefix(types.AllowanceListCountKey)
	bz := []byte(strconv.FormatUint(count, 10))
	store.Set(byteKey, bz)
}

// AppendAllowanceList appends a allowanceList in the store with a new id and update the count
func (k Keeper) AppendAllowanceList(
	ctx sdk.Context,
	allowanceList types.AllowanceList,
) uint64 {
	// Create the allowanceList
	count := k.GetAllowanceListCount(ctx)

	// Set the ID of the appended value
	allowanceList.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AllowanceListKey))
	appendedValue := k.cdc.MustMarshalBinaryBare(&allowanceList)
	store.Set(GetAllowanceListIDBytes(allowanceList.Id), appendedValue)

	// Update allowanceList count
	k.SetAllowanceListCount(ctx, count+1)

	return count
}
*/

// SetAllowanceList set a specific allowanceList in the store
func (k Keeper) SetAllowanceList(ctx sdk.Context, allowanceList types.AllowanceList) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AllowanceListKey))
	b := k.cdc.MustMarshalBinaryBare(&allowanceList)
	store.Set(GetAllowanceListIDBytes(allowanceList.Id), b)
}

// GetAllowanceList returns a allowanceList from its id
func (k Keeper) GetAllowanceList(ctx sdk.Context, id uint64) types.AllowanceList {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AllowanceListKey))
	var allowanceList types.AllowanceList
	k.cdc.MustUnmarshalBinaryBare(store.Get(GetAllowanceListIDBytes(id)), &allowanceList)
	return allowanceList
}

// HasAllowanceList checks if the allowanceList exists in the store
func (k Keeper) HasAllowanceList(ctx sdk.Context, id uint64) bool {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AllowanceListKey))
	return store.Has(GetAllowanceListIDBytes(id))
}

/*
// GetAllowanceListOwner returns the creator of the
func (k Keeper) GetAllowanceListOwner(ctx sdk.Context, id uint64) string {
	return k.GetAllowanceList(ctx, id).Creator
}
*/

// RemoveAllowanceList removes a allowanceList from the store
func (k Keeper) RemoveAllowanceList(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AllowanceListKey))
	store.Delete(GetAllowanceListIDBytes(id))
}

// GetAllAllowanceList returns all allowanceList
func (k Keeper) GetAllAllowanceList(ctx sdk.Context) (list []types.AllowanceList) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AllowanceListKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.AllowanceList
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetAllowanceListIDBytes returns the byte representation of the ID
func GetAllowanceListIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetAllowanceListIDFromBytes returns ID in uint64 format from a byte array
func GetAllowanceListIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
