package keeper

import (
	"encoding/binary"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/konstellation/konstellation/x/issue/types"
	//	"strconv"
)

/*
// GetCoinIssueDenomCount get the total number of TypeName.LowerCamel
func (k Keeper) GetCoinIssueDenomCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CoinIssueDenomCountKey))
	byteKey := types.KeyPrefix(types.CoinIssueDenomCountKey)
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

// SetCoinIssueDenomCount set the total number of coinIssueDenom
func (k Keeper) SetCoinIssueDenomCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CoinIssueDenomCountKey))
	byteKey := types.KeyPrefix(types.CoinIssueDenomCountKey)
	bz := []byte(strconv.FormatUint(count, 10))
	store.Set(byteKey, bz)
}

// AppendCoinIssueDenom appends a coinIssueDenom in the store with a new id and update the count
func (k Keeper) AppendCoinIssueDenom(
	ctx sdk.Context,
	coinIssueDenom types.CoinIssueDenom,
) uint64 {
	// Create the coinIssueDenom
	count := k.GetCoinIssueDenomCount(ctx)

	// Set the ID of the appended value
	coinIssueDenom.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CoinIssueDenomKey))
	appendedValue := k.cdc.MustMarshalBinaryBare(&coinIssueDenom)
	store.Set(GetCoinIssueDenomIDBytes(coinIssueDenom.Id), appendedValue)

	// Update coinIssueDenom count
	k.SetCoinIssueDenomCount(ctx, count+1)

	return count
}
*/

// SetCoinIssueDenom set a specific coinIssueDenom in the store
func (k Keeper) SetCoinIssueDenom(ctx sdk.Context, coinIssueDenom types.CoinIssueDenom) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CoinIssueDenomKey))
	b := k.cdc.MustMarshalBinaryBare(&coinIssueDenom)
	store.Set(GetCoinIssueDenomIDBytes(coinIssueDenom.Id), b)
}

// GetCoinIssueDenom returns a coinIssueDenom from its id
func (k Keeper) GetCoinIssueDenom(ctx sdk.Context, id uint64) types.CoinIssueDenom {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CoinIssueDenomKey))
	var coinIssueDenom types.CoinIssueDenom
	k.cdc.MustUnmarshalBinaryBare(store.Get(GetCoinIssueDenomIDBytes(id)), &coinIssueDenom)
	return coinIssueDenom
}

// HasCoinIssueDenom checks if the coinIssueDenom exists in the store
func (k Keeper) HasCoinIssueDenom(ctx sdk.Context, id uint64) bool {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CoinIssueDenomKey))
	return store.Has(GetCoinIssueDenomIDBytes(id))
}

/*
// GetCoinIssueDenomOwner returns the creator of the
func (k Keeper) GetCoinIssueDenomOwner(ctx sdk.Context, id uint64) string {
	return k.GetCoinIssueDenom(ctx, id).Creator
}
*/

// RemoveCoinIssueDenom removes a coinIssueDenom from the store
func (k Keeper) RemoveCoinIssueDenom(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CoinIssueDenomKey))
	store.Delete(GetCoinIssueDenomIDBytes(id))
}

// GetAllCoinIssueDenom returns all coinIssueDenom
func (k Keeper) GetAllCoinIssueDenom(ctx sdk.Context) (list []types.CoinIssueDenom) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CoinIssueDenomKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.CoinIssueDenom
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetCoinIssueDenomIDBytes returns the byte representation of the ID
func GetCoinIssueDenomIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetCoinIssueDenomIDFromBytes returns ID in uint64 format from a byte array
func GetCoinIssueDenomIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
