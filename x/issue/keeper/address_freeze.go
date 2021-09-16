package keeper

import (
	"encoding/binary"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/konstellation/konstellation/x/issue/types"
	//	"strconv"
)

/*
// GetAddressFreezeCount get the total number of TypeName.LowerCamel
func (k Keeper) GetAddressFreezeCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AddressFreezeCountKey))
	byteKey := types.KeyPrefix(types.AddressFreezeCountKey)
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

// SetAddressFreezeCount set the total number of addressFreeze
func (k Keeper) SetAddressFreezeCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AddressFreezeCountKey))
	byteKey := types.KeyPrefix(types.AddressFreezeCountKey)
	bz := []byte(strconv.FormatUint(count, 10))
	store.Set(byteKey, bz)
}

// AppendAddressFreeze appends a addressFreeze in the store with a new id and update the count
func (k Keeper) AppendAddressFreeze(
	ctx sdk.Context,
	addressFreeze types.AddressFreeze,
) uint64 {
	// Create the addressFreeze
	count := k.GetAddressFreezeCount(ctx)

	// Set the ID of the appended value
	addressFreeze.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AddressFreezeKey))
	appendedValue := k.cdc.MustMarshalBinaryBare(&addressFreeze)
	store.Set(GetAddressFreezeIDBytes(addressFreeze.Id), appendedValue)

	// Update addressFreeze count
	k.SetAddressFreezeCount(ctx, count+1)

	return count
}
*/

// SetAddressFreeze set a specific addressFreeze in the store
func (k Keeper) SetAddressFreeze(ctx sdk.Context, addressFreeze types.AddressFreeze) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AddressFreezeKey))
	b := k.cdc.MustMarshalBinaryBare(&addressFreeze)
	store.Set(GetAddressFreezeIDBytes(addressFreeze.Id), b)
}

// GetAddressFreeze returns a addressFreeze from its id
func (k Keeper) GetAddressFreeze(ctx sdk.Context, id uint64) types.AddressFreeze {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AddressFreezeKey))
	var addressFreeze types.AddressFreeze
	k.cdc.MustUnmarshalBinaryBare(store.Get(GetAddressFreezeIDBytes(id)), &addressFreeze)
	return addressFreeze
}

// HasAddressFreeze checks if the addressFreeze exists in the store
func (k Keeper) HasAddressFreeze(ctx sdk.Context, id uint64) bool {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AddressFreezeKey))
	return store.Has(GetAddressFreezeIDBytes(id))
}

/*
// GetAddressFreezeOwner returns the creator of the
func (k Keeper) GetAddressFreezeOwner(ctx sdk.Context, id uint64) string {
	return k.GetAddressFreeze(ctx, id).Creator
}
*/

// RemoveAddressFreeze removes a addressFreeze from the store
func (k Keeper) RemoveAddressFreeze(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AddressFreezeKey))
	store.Delete(GetAddressFreezeIDBytes(id))
}

// GetAllAddressFreeze returns all addressFreeze
func (k Keeper) GetAllAddressFreeze(ctx sdk.Context) (list []types.AddressFreeze) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AddressFreezeKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.AddressFreeze
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetAddressFreezeIDBytes returns the byte representation of the ID
func GetAddressFreezeIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetAddressFreezeIDFromBytes returns ID in uint64 format from a byte array
func GetAddressFreezeIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
