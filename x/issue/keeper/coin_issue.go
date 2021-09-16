package keeper

import (
	"encoding/binary"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/konstellation/konstellation/x/issue/types"
	//"strconv"
)

/*
// GetCoinIssueCount get the total number of TypeName.LowerCamel
func (k Keeper) GetCoinIssueCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CoinIssueCountKey))
	byteKey := types.KeyPrefix(types.CoinIssueCountKey)
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

// SetCoinIssueCount set the total number of coinIssue
func (k Keeper) SetCoinIssueCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CoinIssueCountKey))
	byteKey := types.KeyPrefix(types.CoinIssueCountKey)
	bz := []byte(strconv.FormatUint(count, 10))
	store.Set(byteKey, bz)
}

// AppendCoinIssue appends a coinIssue in the store with a new id and update the count
func (k Keeper) AppendCoinIssue(
	ctx sdk.Context,
	coinIssue types.CoinIssue,
) uint64 {
	// Create the coinIssue
	count := k.GetCoinIssueCount(ctx)

	// Set the ID of the appended value
	coinIssue.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CoinIssueKey))
	appendedValue := k.cdc.MustMarshalBinaryBare(&coinIssue)
	store.Set(GetCoinIssueIDBytes(coinIssue.Id), appendedValue)

	// Update coinIssue count
	k.SetCoinIssueCount(ctx, count+1)

	return count
}
*/

// SetCoinIssue set a specific coinIssue in the store
func (k Keeper) SetCoinIssue(ctx sdk.Context, coinIssue types.CoinIssue) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CoinIssueKey))
	b := k.cdc.MustMarshalBinaryBare(&coinIssue)
	store.Set(GetCoinIssueIDBytes(coinIssue.Id), b)
}

// GetCoinIssue returns a coinIssue from its id
func (k Keeper) GetCoinIssue(ctx sdk.Context, id uint64) types.CoinIssue {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CoinIssueKey))
	var coinIssue types.CoinIssue
	k.cdc.MustUnmarshalBinaryBare(store.Get(GetCoinIssueIDBytes(id)), &coinIssue)
	return coinIssue
}

// HasCoinIssue checks if the coinIssue exists in the store
func (k Keeper) HasCoinIssue(ctx sdk.Context, id uint64) bool {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CoinIssueKey))
	return store.Has(GetCoinIssueIDBytes(id))
}

/*
// GetCoinIssueOwner returns the creator of the
func (k Keeper) GetCoinIssueOwner(ctx sdk.Context, id uint64) string {
	return k.GetCoinIssue(ctx, id).Creator
}
*/

// RemoveCoinIssue removes a coinIssue from the store
func (k Keeper) RemoveCoinIssue(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CoinIssueKey))
	store.Delete(GetCoinIssueIDBytes(id))
}

// GetAllCoinIssue returns all coinIssue
func (k Keeper) GetAllCoinIssue(ctx sdk.Context) (list []types.CoinIssue) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CoinIssueKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.CoinIssue
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetCoinIssueIDBytes returns the byte representation of the ID
func GetCoinIssueIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetCoinIssueIDFromBytes returns ID in uint64 format from a byte array
func GetCoinIssueIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
