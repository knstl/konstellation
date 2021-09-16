package keeper

import (
	"encoding/binary"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/konstellation/konstellation/x/issue/types"
	//	"strconv"
)

/*
// GetIssueParamsCount get the total number of TypeName.LowerCamel
func (k Keeper) GetIssueParamsCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.IssueParamsCountKey))
	byteKey := types.KeyPrefix(types.IssueParamsCountKey)
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

// SetIssueParamsCount set the total number of issueParams
func (k Keeper) SetIssueParamsCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.IssueParamsCountKey))
	byteKey := types.KeyPrefix(types.IssueParamsCountKey)
	bz := []byte(strconv.FormatUint(count, 10))
	store.Set(byteKey, bz)
}

// AppendIssueParams appends a issueParams in the store with a new id and update the count
func (k Keeper) AppendIssueParams(
	ctx sdk.Context,
	issueParams types.IssueParams,
) uint64 {
	// Create the issueParams
	count := k.GetIssueParamsCount(ctx)

	// Set the ID of the appended value
	issueParams.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.IssueParamsKey))
	appendedValue := k.cdc.MustMarshalBinaryBare(&issueParams)
	store.Set(GetIssueParamsIDBytes(issueParams.Id), appendedValue)

	// Update issueParams count
	k.SetIssueParamsCount(ctx, count+1)

	return count
}
*/

// SetIssueParams set a specific issueParams in the store
func (k Keeper) SetIssueParams(ctx sdk.Context, issueParams types.IssueParams) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.IssueParamsKey))
	b := k.cdc.MustMarshalBinaryBare(&issueParams)
	store.Set(GetIssueParamsIDBytes(issueParams.Id), b)
}

// GetIssueParams returns a issueParams from its id
func (k Keeper) GetIssueParams(ctx sdk.Context, id uint64) types.IssueParams {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.IssueParamsKey))
	var issueParams types.IssueParams
	k.cdc.MustUnmarshalBinaryBare(store.Get(GetIssueParamsIDBytes(id)), &issueParams)
	return issueParams
}

// HasIssueParams checks if the issueParams exists in the store
func (k Keeper) HasIssueParams(ctx sdk.Context, id uint64) bool {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.IssueParamsKey))
	return store.Has(GetIssueParamsIDBytes(id))
}

/*
// GetIssueParamsOwner returns the creator of the
func (k Keeper) GetIssueParamsOwner(ctx sdk.Context, id uint64) string {
	return k.GetIssueParams(ctx, id).Creator
}
*/

// RemoveIssueParams removes a issueParams from the store
func (k Keeper) RemoveIssueParams(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.IssueParamsKey))
	store.Delete(GetIssueParamsIDBytes(id))
}

// GetAllIssueParams returns all issueParams
func (k Keeper) GetAllIssueParams(ctx sdk.Context) (list []types.IssueParams) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.IssueParamsKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.IssueParams
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetIssueParamsIDBytes returns the byte representation of the ID
func GetIssueParamsIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetIssueParamsIDFromBytes returns ID in uint64 format from a byte array
func GetIssueParamsIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
