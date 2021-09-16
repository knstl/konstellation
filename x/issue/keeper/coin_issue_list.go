package keeper

import (
	"encoding/binary"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/konstellation/konstellation/x/issue/types"
	//	"strconv"
)

/*
// GetCoinIssueListCount get the total number of TypeName.LowerCamel
func (k Keeper) GetCoinIssueListCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CoinIssueListCountKey))
	byteKey := types.KeyPrefix(types.CoinIssueListCountKey)
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

// SetCoinIssueListCount set the total number of coinIssueList
func (k Keeper) SetCoinIssueListCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CoinIssueListCountKey))
	byteKey := types.KeyPrefix(types.CoinIssueListCountKey)
	bz := []byte(strconv.FormatUint(count, 10))
	store.Set(byteKey, bz)
}

// AppendCoinIssueList appends a coinIssueList in the store with a new id and update the count
func (k Keeper) AppendCoinIssueList(
	ctx sdk.Context,
	coinIssueList types.CoinIssueList,
) uint64 {
	// Create the coinIssueList
	count := k.GetCoinIssueListCount(ctx)

	// Set the ID of the appended value
	coinIssueList.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CoinIssueListKey))
	appendedValue := k.cdc.MustMarshalBinaryBare(&coinIssueList)
	store.Set(GetCoinIssueListIDBytes(coinIssueList.Id), appendedValue)

	// Update coinIssueList count
	k.SetCoinIssueListCount(ctx, count+1)

	return count
}
*/

// SetCoinIssueList set a specific coinIssueList in the store
func (k Keeper) SetCoinIssueList(ctx sdk.Context, coinIssueList types.CoinIssueList) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CoinIssueListKey))
	b := k.cdc.MustMarshalBinaryBare(&coinIssueList)
	store.Set(GetCoinIssueListIDBytes(coinIssueList.Id), b)
}

// GetCoinIssueList returns a coinIssueList from its id
func (k Keeper) GetCoinIssueList(ctx sdk.Context, id uint64) types.CoinIssueList {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CoinIssueListKey))
	var coinIssueList types.CoinIssueList
	k.cdc.MustUnmarshalBinaryBare(store.Get(GetCoinIssueListIDBytes(id)), &coinIssueList)
	return coinIssueList
}

// HasCoinIssueList checks if the coinIssueList exists in the store
func (k Keeper) HasCoinIssueList(ctx sdk.Context, id uint64) bool {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CoinIssueListKey))
	return store.Has(GetCoinIssueListIDBytes(id))
}

/*
// GetCoinIssueListOwner returns the creator of the
func (k Keeper) GetCoinIssueListOwner(ctx sdk.Context, id uint64) string {
	return k.GetCoinIssueList(ctx, id).Creator
}
*/

// RemoveCoinIssueList removes a coinIssueList from the store
func (k Keeper) RemoveCoinIssueList(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CoinIssueListKey))
	store.Delete(GetCoinIssueListIDBytes(id))
}

// GetAllCoinIssueList returns all coinIssueList
func (k Keeper) GetAllCoinIssueList(ctx sdk.Context) (list []types.CoinIssueList) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CoinIssueListKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.CoinIssueList
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetCoinIssueListIDBytes returns the byte representation of the ID
func GetCoinIssueListIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetCoinIssueListIDFromBytes returns ID in uint64 format from a byte array
func GetCoinIssueListIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
