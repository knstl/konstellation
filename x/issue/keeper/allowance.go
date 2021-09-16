package keeper

import (
	"encoding/binary"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/konstellation/konstellation/x/issue/types"
	//	"strconv"
)

/*
// GetAllowanceCount get the total number of TypeName.LowerCamel
func (k Keeper) GetAllowanceCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AllowanceCountKey))
	byteKey := types.KeyPrefix(types.AllowanceCountKey)
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

// SetAllowanceCount set the total number of allowance
func (k Keeper) SetAllowanceCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AllowanceCountKey))
	byteKey := types.KeyPrefix(types.AllowanceCountKey)
	bz := []byte(strconv.FormatUint(count, 10))
	store.Set(byteKey, bz)
}

// AppendAllowance appends a allowance in the store with a new id and update the count
func (k Keeper) AppendAllowance(
	ctx sdk.Context,
	allowance types.Allowance,
) uint64 {
	// Create the allowance
	count := k.GetAllowanceCount(ctx)

	// Set the ID of the appended value
	allowance.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AllowanceKey))
	appendedValue := k.cdc.MustMarshalBinaryBare(&allowance)
	store.Set(GetAllowanceIDBytes(allowance.Id), appendedValue)

	// Update allowance count
	k.SetAllowanceCount(ctx, count+1)

	return count
}
*/

// SetAllowance set a specific allowance in the store
func (k Keeper) SetAllowance(ctx sdk.Context, allowance types.Allowance) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AllowanceKey))
	b := k.cdc.MustMarshalBinaryBare(&allowance)
	store.Set(GetAllowanceIDBytes(allowance.Id), b)
}

// GetAllowance returns a allowance from its id
func (k Keeper) GetAllowance(ctx sdk.Context, id uint64) types.Allowance {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AllowanceKey))
	var allowance types.Allowance
	k.cdc.MustUnmarshalBinaryBare(store.Get(GetAllowanceIDBytes(id)), &allowance)
	return allowance
}

// HasAllowance checks if the allowance exists in the store
func (k Keeper) HasAllowance(ctx sdk.Context, id uint64) bool {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AllowanceKey))
	return store.Has(GetAllowanceIDBytes(id))
}

/*
// GetAllowanceOwner returns the creator of the
func (k Keeper) GetAllowanceOwner(ctx sdk.Context, id uint64) string {
	return k.GetAllowance(ctx, id).Creator
}
*/

// RemoveAllowance removes a allowance from the store
func (k Keeper) RemoveAllowance(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AllowanceKey))
	store.Delete(GetAllowanceIDBytes(id))
}

// GetAllAllowance returns all allowance
func (k Keeper) GetAllAllowance(ctx sdk.Context) (list []types.Allowance) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AllowanceKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Allowance
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetAllowanceIDBytes returns the byte representation of the ID
func GetAllowanceIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetAllowanceIDFromBytes returns ID in uint64 format from a byte array
func GetAllowanceIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
