package keeper

import (
	"encoding/binary"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/konstellation/konstellation/x/issue/types"
	//	"strconv"
)

/*
// GetCoinIssueDenomsCount get the total number of TypeName.LowerCamel
func (k Keeper) GetCoinIssueDenomsCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CoinIssueDenomsCountKey))
	byteKey := types.KeyPrefix(types.CoinIssueDenomsCountKey)
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

// SetCoinIssueDenomsCount set the total number of coinIssueDenoms
func (k Keeper) SetCoinIssueDenomsCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CoinIssueDenomsCountKey))
	byteKey := types.KeyPrefix(types.CoinIssueDenomsCountKey)
	bz := []byte(strconv.FormatUint(count, 10))
	store.Set(byteKey, bz)
}

// AppendCoinIssueDenoms appends a coinIssueDenoms in the store with a new id and update the count
func (k Keeper) AppendCoinIssueDenoms(
	ctx sdk.Context,
	coinIssueDenoms types.CoinIssueDenoms,
) uint64 {
	// Create the coinIssueDenoms
	count := k.GetCoinIssueDenomsCount(ctx)

	// Set the ID of the appended value
	coinIssueDenoms.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CoinIssueDenomsKey))
	appendedValue := k.cdc.MustMarshalBinaryBare(&coinIssueDenoms)
	store.Set(GetCoinIssueDenomsIDBytes(coinIssueDenoms.Id), appendedValue)

	// Update coinIssueDenoms count
	k.SetCoinIssueDenomsCount(ctx, count+1)

	return count
}
*/

// SetCoinIssueDenoms set a specific coinIssueDenoms in the store
func (k Keeper) SetCoinIssueDenoms(ctx sdk.Context, coinIssueDenoms types.CoinIssueDenoms) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CoinIssueDenomsKey))
	b := k.cdc.MustMarshalBinaryBare(&coinIssueDenoms)
	store.Set(GetCoinIssueDenomsIDBytes(coinIssueDenoms.Id), b)
}

// GetCoinIssueDenoms returns a coinIssueDenoms from its id
func (k Keeper) GetCoinIssueDenoms(ctx sdk.Context, id uint64) types.CoinIssueDenoms {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CoinIssueDenomsKey))
	var coinIssueDenoms types.CoinIssueDenoms
	k.cdc.MustUnmarshalBinaryBare(store.Get(GetCoinIssueDenomsIDBytes(id)), &coinIssueDenoms)
	return coinIssueDenoms
}

// HasCoinIssueDenoms checks if the coinIssueDenoms exists in the store
func (k Keeper) HasCoinIssueDenoms(ctx sdk.Context, id uint64) bool {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CoinIssueDenomsKey))
	return store.Has(GetCoinIssueDenomsIDBytes(id))
}

/*
// GetCoinIssueDenomsOwner returns the creator of the
func (k Keeper) GetCoinIssueDenomsOwner(ctx sdk.Context, id uint64) string {
	return k.GetCoinIssueDenoms(ctx, id).Creator
}
*/

// RemoveCoinIssueDenoms removes a coinIssueDenoms from the store
func (k Keeper) RemoveCoinIssueDenoms(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CoinIssueDenomsKey))
	store.Delete(GetCoinIssueDenomsIDBytes(id))
}

// GetAllCoinIssueDenoms returns all coinIssueDenoms
func (k Keeper) GetAllCoinIssueDenoms(ctx sdk.Context) (list []types.CoinIssueDenoms) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CoinIssueDenomsKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.CoinIssueDenoms
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetCoinIssueDenomsIDBytes returns the byte representation of the ID
func GetCoinIssueDenomsIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetCoinIssueDenomsIDFromBytes returns ID in uint64 format from a byte array
func GetCoinIssueDenomsIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
