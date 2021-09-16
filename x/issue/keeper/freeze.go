package keeper

import (
	"encoding/binary"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	//	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/konstellation/konstellation/x/issue/types"
	//	"strconv"
)

/*
// GetFreezeCount get the total number of TypeName.LowerCamel
func (k Keeper) GetFreezeCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.FreezeCountKey))
	byteKey := types.KeyPrefix(types.FreezeCountKey)
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

// SetFreezeCount set the total number of freeze
func (k Keeper) SetFreezeCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.FreezeCountKey))
	byteKey := types.KeyPrefix(types.FreezeCountKey)
	bz := []byte(strconv.FormatUint(count, 10))
	store.Set(byteKey, bz)
}

// AppendFreeze appends a freeze in the store with a new id and update the count
func (k Keeper) AppendFreeze(
	ctx sdk.Context,
	freeze types.Freeze,
) uint64 {
	// Create the freeze
	count := k.GetFreezeCount(ctx)

	// Set the ID of the appended value
	freeze.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.FreezeKey))
	appendedValue := k.cdc.MustMarshalBinaryBare(&freeze)
	store.Set(GetFreezeIDBytes(freeze.Id), appendedValue)

	// Update freeze count
	k.SetFreezeCount(ctx, count+1)

	return count
}

// SetFreeze set a specific freeze in the store
func (k Keeper) SetFreeze(ctx sdk.Context, freeze types.Freeze) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.FreezeKey))
	b := k.cdc.MustMarshalBinaryBare(&freeze)
	store.Set(GetFreezeIDBytes(freeze.Id), b)
}

// GetFreeze returns a freeze from its id
func (k Keeper) GetFreeze(ctx sdk.Context, id uint64) types.Freeze {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.FreezeKey))
	var freeze types.Freeze
	k.cdc.MustUnmarshalBinaryBare(store.Get(GetFreezeIDBytes(id)), &freeze)
	return freeze
}
*/

func (k *Keeper) GetFreeze(ctx sdk.Context, denom string, holder sdk.AccAddress) *types.Freeze {
	return k.getFreeze(ctx, denom, holder)
}

// HasFreeze checks if the freeze exists in the store
func (k Keeper) HasFreeze(ctx sdk.Context, id uint64) bool {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.FreezeKey))
	return store.Has(GetFreezeIDBytes(id))
}

/*
// GetFreezeOwner returns the creator of the
func (k Keeper) GetFreezeOwner(ctx sdk.Context, id uint64) string {
	return k.GetFreeze(ctx, id).Creator
}
*/

// RemoveFreeze removes a freeze from the store
func (k Keeper) RemoveFreeze(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.FreezeKey))
	store.Delete(GetFreezeIDBytes(id))
}

// GetAllFreeze returns all freeze
func (k Keeper) GetAllFreeze(ctx sdk.Context) (list []types.Freeze) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.FreezeKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Freeze
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetFreezeIDBytes returns the byte representation of the ID
func GetFreezeIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetFreezeIDFromBytes returns ID in uint64 format from a byte array
func GetFreezeIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
