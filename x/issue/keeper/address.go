package keeper

import (
	"encoding/binary"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/konstellation/konstellation/x/issue/types"
	//	"strconv"
)

/*
// GetAddressCount get the total number of TypeName.LowerCamel
func (k Keeper) GetAddressCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AddressCountKey))
	byteKey := types.KeyPrefix(types.AddressCountKey)
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

// SetAddressCount set the total number of address
func (k Keeper) SetAddressCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AddressCountKey))
	byteKey := types.KeyPrefix(types.AddressCountKey)
	bz := []byte(strconv.FormatUint(count, 10))
	store.Set(byteKey, bz)
}

// AppendAddress appends a address in the store with a new id and update the count
func (k Keeper) AppendAddress(
	ctx sdk.Context,
	address types.Address,
) uint64 {
	// Create the address
	count := k.GetAddressCount(ctx)

	// Set the ID of the appended value
	address.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AddressKey))
	appendedValue := k.cdc.MustMarshalBinaryBare(&address)
	store.Set(GetAddressIDBytes(address.Id), appendedValue)

	// Update address count
	k.SetAddressCount(ctx, count+1)

	return count
}
*/

// SetAddress set a specific address in the store
func (k Keeper) SetAddress(ctx sdk.Context, address types.Address) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AddressKey))
	b := k.cdc.MustMarshalBinaryBare(&address)
	store.Set(GetAddressIDBytes(address.Id), b)
}

// GetAddress returns a address from its id
func (k Keeper) GetAddress(ctx sdk.Context, id uint64) types.Address {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AddressKey))
	var address types.Address
	k.cdc.MustUnmarshalBinaryBare(store.Get(GetAddressIDBytes(id)), &address)
	return address
}

// HasAddress checks if the address exists in the store
func (k Keeper) HasAddress(ctx sdk.Context, id uint64) bool {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AddressKey))
	return store.Has(GetAddressIDBytes(id))
}

/*
// GetAddressOwner returns the creator of the
func (k Keeper) GetAddressOwner(ctx sdk.Context, id uint64) string {
	return k.GetAddress(ctx, id).Creator
}
*/

// RemoveAddress removes a address from the store
func (k Keeper) RemoveAddress(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AddressKey))
	store.Delete(GetAddressIDBytes(id))
}

// GetAllAddress returns all address
func (k Keeper) GetAllAddress(ctx sdk.Context) (list []types.Address) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AddressKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Address
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetAddressIDBytes returns the byte representation of the ID
func GetAddressIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetAddressIDFromBytes returns ID in uint64 format from a byte array
func GetAddressIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
