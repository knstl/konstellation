package keeper

import (
	"encoding/binary"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/konstellation/konstellation/x/issue/types"
	//	"strconv"
)

/*
// GetCoinsCount get the total number of TypeName.LowerCamel
func (k Keeper) GetCoinsCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CoinsCountKey))
	byteKey := types.KeyPrefix(types.CoinsCountKey)
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

// SetCoinsCount set the total number of coins
func (k Keeper) SetCoinsCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CoinsCountKey))
	byteKey := types.KeyPrefix(types.CoinsCountKey)
	bz := []byte(strconv.FormatUint(count, 10))
	store.Set(byteKey, bz)
}

// AppendCoins appends a coins in the store with a new id and update the count
func (k Keeper) AppendCoins(
	ctx sdk.Context,
	coins types.Coins,
) uint64 {
	// Create the coins
	count := k.GetCoinsCount(ctx)

	// Set the ID of the appended value
	coins.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CoinsKey))
	appendedValue := k.cdc.MustMarshalBinaryBare(&coins)
	store.Set(GetCoinsIDBytes(coins.Id), appendedValue)

	// Update coins count
	k.SetCoinsCount(ctx, count+1)

	return count
}
*/

// SetCoins set a specific coins in the store
func (k Keeper) SetCoins(ctx sdk.Context, coins types.Coins) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CoinsKey))
	b := k.cdc.MustMarshalBinaryBare(&coins)
	store.Set(GetCoinsIDBytes(coins.Id), b)
}

// GetCoins returns a coins from its id
func (k Keeper) GetCoins(ctx sdk.Context, id uint64) types.Coins {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CoinsKey))
	var coins types.Coins
	k.cdc.MustUnmarshalBinaryBare(store.Get(GetCoinsIDBytes(id)), &coins)
	return coins
}

// HasCoins checks if the coins exists in the store
func (k Keeper) HasCoins(ctx sdk.Context, id uint64) bool {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CoinsKey))
	return store.Has(GetCoinsIDBytes(id))
}

/*
// GetCoinsOwner returns the creator of the
func (k Keeper) GetCoinsOwner(ctx sdk.Context, id uint64) string {
	return k.GetCoins(ctx, id).Creator
}
*/

// RemoveCoins removes a coins from the store
func (k Keeper) RemoveCoins(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CoinsKey))
	store.Delete(GetCoinsIDBytes(id))
}

// GetAllCoins returns all coins
func (k Keeper) GetAllCoins(ctx sdk.Context) (list []types.Coins) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CoinsKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Coins
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetCoinsIDBytes returns the byte representation of the ID
func GetCoinsIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetCoinsIDFromBytes returns ID in uint64 format from a byte array
func GetCoinsIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
