package keeper

import (
	"encoding/binary"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/konstellation/konstellation/x/issue/types"
	//	"strconv"
)

/*
// GetAddressFreezeListCount get the total number of TypeName.LowerCamel
func (k Keeper) GetAddressFreezeListCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AddressFreezeListCountKey))
	byteKey := types.KeyPrefix(types.AddressFreezeListCountKey)
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

// SetAddressFreezeListCount set the total number of addressFreezeList
func (k Keeper) SetAddressFreezeListCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AddressFreezeListCountKey))
	byteKey := types.KeyPrefix(types.AddressFreezeListCountKey)
	bz := []byte(strconv.FormatUint(count, 10))
	store.Set(byteKey, bz)
}

// AppendAddressFreezeList appends a addressFreezeList in the store with a new id and update the count
func (k Keeper) AppendAddressFreezeList(
	ctx sdk.Context,
	addressFreezeList types.AddressFreezeList,
) uint64 {
	// Create the addressFreezeList
	count := k.GetAddressFreezeListCount(ctx)

	// Set the ID of the appended value
	addressFreezeList.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AddressFreezeListKey))
	appendedValue := k.cdc.MustMarshalBinaryBare(&addressFreezeList)
	store.Set(GetAddressFreezeListIDBytes(addressFreezeList.Id), appendedValue)

	// Update addressFreezeList count
	k.SetAddressFreezeListCount(ctx, count+1)

	return count
}
*/

// SetAddressFreezeList set a specific addressFreezeList in the store
func (k Keeper) SetAddressFreezeList(ctx sdk.Context, addressFreezeList types.AddressFreezeList) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AddressFreezeListKey))
	b := k.cdc.MustMarshalBinaryBare(&addressFreezeList)
	store.Set(GetAddressFreezeListIDBytes(addressFreezeList.Id), b)
}

// GetAddressFreezeList returns a addressFreezeList from its id
func (k Keeper) GetAddressFreezeList(ctx sdk.Context, id uint64) types.AddressFreezeList {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AddressFreezeListKey))
	var addressFreezeList types.AddressFreezeList
	k.cdc.MustUnmarshalBinaryBare(store.Get(GetAddressFreezeListIDBytes(id)), &addressFreezeList)
	return addressFreezeList
}

// HasAddressFreezeList checks if the addressFreezeList exists in the store
func (k Keeper) HasAddressFreezeList(ctx sdk.Context, id uint64) bool {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AddressFreezeListKey))
	return store.Has(GetAddressFreezeListIDBytes(id))
}

/*
// GetAddressFreezeListOwner returns the creator of the
func (k Keeper) GetAddressFreezeListOwner(ctx sdk.Context, id uint64) string {
	return k.GetAddressFreezeList(ctx, id).Creator
}
*/

// RemoveAddressFreezeList removes a addressFreezeList from the store
func (k Keeper) RemoveAddressFreezeList(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AddressFreezeListKey))
	store.Delete(GetAddressFreezeListIDBytes(id))
}

// GetAllAddressFreezeList returns all addressFreezeList
func (k Keeper) GetAllAddressFreezeList(ctx sdk.Context) (list []types.AddressFreezeList) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AddressFreezeListKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.AddressFreezeList
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetAddressFreezeListIDBytes returns the byte representation of the ID
func GetAddressFreezeListIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetAddressFreezeListIDFromBytes returns ID in uint64 format from a byte array
func GetAddressFreezeListIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
