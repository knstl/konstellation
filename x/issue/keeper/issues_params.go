package keeper

import (
	"encoding/binary"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/konstellation/konstellation/x/issue/types"
	//	"strconv"
)

/*
// GetIssuesParamsCount get the total number of TypeName.LowerCamel
func (k Keeper) GetIssuesParamsCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.IssuesParamsCountKey))
	byteKey := types.KeyPrefix(types.IssuesParamsCountKey)
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

// SetIssuesParamsCount set the total number of issuesParams
func (k Keeper) SetIssuesParamsCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.IssuesParamsCountKey))
	byteKey := types.KeyPrefix(types.IssuesParamsCountKey)
	bz := []byte(strconv.FormatUint(count, 10))
	store.Set(byteKey, bz)
}

// AppendIssuesParams appends a issuesParams in the store with a new id and update the count
func (k Keeper) AppendIssuesParams(
	ctx sdk.Context,
	issuesParams types.IssuesParams,
) uint64 {
	// Create the issuesParams
	count := k.GetIssuesParamsCount(ctx)

	// Set the ID of the appended value
	issuesParams.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.IssuesParamsKey))
	appendedValue := k.cdc.MustMarshalBinaryBare(&issuesParams)
	store.Set(GetIssuesParamsIDBytes(issuesParams.Id), appendedValue)

	// Update issuesParams count
	k.SetIssuesParamsCount(ctx, count+1)

	return count
}
*/

// SetIssuesParams set a specific issuesParams in the store
func (k Keeper) SetIssuesParams(ctx sdk.Context, issuesParams types.IssuesParams) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.IssuesParamsKey))
	b := k.cdc.MustMarshalBinaryBare(&issuesParams)
	store.Set(GetIssuesParamsIDBytes(issuesParams.Id), b)
}

// GetIssuesParams returns a issuesParams from its id
func (k Keeper) GetIssuesParams(ctx sdk.Context, id uint64) types.IssuesParams {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.IssuesParamsKey))
	var issuesParams types.IssuesParams
	k.cdc.MustUnmarshalBinaryBare(store.Get(GetIssuesParamsIDBytes(id)), &issuesParams)
	return issuesParams
}

// HasIssuesParams checks if the issuesParams exists in the store
func (k Keeper) HasIssuesParams(ctx sdk.Context, id uint64) bool {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.IssuesParamsKey))
	return store.Has(GetIssuesParamsIDBytes(id))
}

/*
// GetIssuesParamsOwner returns the creator of the
func (k Keeper) GetIssuesParamsOwner(ctx sdk.Context, id uint64) string {
	return k.GetIssuesParams(ctx, id).Creator
}
*/

// RemoveIssuesParams removes a issuesParams from the store
func (k Keeper) RemoveIssuesParams(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.IssuesParamsKey))
	store.Delete(GetIssuesParamsIDBytes(id))
}

// GetAllIssuesParams returns all issuesParams
func (k Keeper) GetAllIssuesParams(ctx sdk.Context) (list []types.IssuesParams) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.IssuesParamsKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.IssuesParams
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetIssuesParamsIDBytes returns the byte representation of the ID
func GetIssuesParamsIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetIssuesParamsIDFromBytes returns ID in uint64 format from a byte array
func GetIssuesParamsIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
