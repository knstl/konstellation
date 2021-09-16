package keeper

import (
	"encoding/binary"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/konstellation/konstellation/x/issue/types"
	//	"strconv"
)

/*
// GetIssuesCount get the total number of TypeName.LowerCamel
func (k Keeper) GetIssuesCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.IssuesCountKey))
	byteKey := types.KeyPrefix(types.IssuesCountKey)
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

// SetIssuesCount set the total number of issues
func (k Keeper) SetIssuesCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.IssuesCountKey))
	byteKey := types.KeyPrefix(types.IssuesCountKey)
	bz := []byte(strconv.FormatUint(count, 10))
	store.Set(byteKey, bz)
}

// AppendIssues appends a issues in the store with a new id and update the count
func (k Keeper) AppendIssues(
	ctx sdk.Context,
	issues types.Issues,
) uint64 {
	// Create the issues
	count := k.GetIssuesCount(ctx)

	// Set the ID of the appended value
	issues.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.IssuesKey))
	appendedValue := k.cdc.MustMarshalBinaryBare(&issues)
	store.Set(GetIssuesIDBytes(issues.Id), appendedValue)

	// Update issues count
	k.SetIssuesCount(ctx, count+1)

	return count
}
*/

// SetIssues set a specific issues in the store
func (k Keeper) SetIssues(ctx sdk.Context, issues types.Issues) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.IssuesKey))
	b := k.cdc.MustMarshalBinaryBare(&issues)
	store.Set(GetIssuesIDBytes(issues.Id), b)
}

// GetIssues returns a issues from its id
func (k Keeper) GetIssues(ctx sdk.Context, id uint64) types.Issues {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.IssuesKey))
	var issues types.Issues
	k.cdc.MustUnmarshalBinaryBare(store.Get(GetIssuesIDBytes(id)), &issues)
	return issues
}

// HasIssues checks if the issues exists in the store
func (k Keeper) HasIssues(ctx sdk.Context, id uint64) bool {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.IssuesKey))
	return store.Has(GetIssuesIDBytes(id))
}

/*
// GetIssuesOwner returns the creator of the
func (k Keeper) GetIssuesOwner(ctx sdk.Context, id uint64) string {
	return k.GetIssues(ctx, id).Creator
}
*/

// RemoveIssues removes a issues from the store
func (k Keeper) RemoveIssues(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.IssuesKey))
	store.Delete(GetIssuesIDBytes(id))
}

// GetAllIssues returns all issues
func (k Keeper) GetAllIssues(ctx sdk.Context) (list []types.Issues) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.IssuesKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Issues
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetIssuesIDBytes returns the byte representation of the ID
func GetIssuesIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetIssuesIDFromBytes returns ID in uint64 format from a byte array
func GetIssuesIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
