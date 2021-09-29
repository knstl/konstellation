package keeper

import (
	"encoding/binary"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/konstellation/konstellation/x/issue/types"
	//	"strconv"
)

/*
// GetIssueFeaturesCount get the total number of TypeName.LowerCamel
func (k Keeper) GetIssueFeaturesCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.IssueFeaturesCountKey))
	byteKey := types.KeyPrefix(types.IssueFeaturesCountKey)
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

// SetIssueFeaturesCount set the total number of issueFeatures
func (k Keeper) SetIssueFeaturesCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.IssueFeaturesCountKey))
	byteKey := types.KeyPrefix(types.IssueFeaturesCountKey)
	bz := []byte(strconv.FormatUint(count, 10))
	store.Set(byteKey, bz)
}

// AppendIssueFeatures appends a issueFeatures in the store with a new id and update the count
func (k Keeper) AppendIssueFeatures(
	ctx sdk.Context,
	issueFeatures types.IssueFeatures,
) uint64 {
	// Create the issueFeatures
	count := k.GetIssueFeaturesCount(ctx)

	// Set the ID of the appended value
	issueFeatures.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.IssueFeaturesKey))
	appendedValue := k.cdc.MustMarshalBinaryBare(&issueFeatures)
	store.Set(GetIssueFeaturesIDBytes(issueFeatures.Id), appendedValue)

	// Update issueFeatures count
	k.SetIssueFeaturesCount(ctx, count+1)

	return count
}
*/

// SetIssueFeatures set a specific issueFeatures in the store
//func (k Keeper) SetIssueFeatures(ctx sdk.Context, issueFeatures types.IssueFeatures) {
//	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.IssueFeaturesKey))
//	b := k.cdc.MustMarshalBinaryBare(&issueFeatures)
//	store.Set(GetIssueFeaturesIDBytes(issueFeatures.Id), b)
//}
//
//// GetIssueFeatures returns a issueFeatures from its id
//func (k Keeper) GetIssueFeatures(ctx sdk.Context, id uint64) types.IssueFeatures {
//	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.IssueFeaturesKey))
//	var issueFeatures types.IssueFeatures
//	k.cdc.MustUnmarshalBinaryBare(store.Get(GetIssueFeaturesIDBytes(id)), &issueFeatures)
//	return issueFeatures
//}
//
//// HasIssueFeatures checks if the issueFeatures exists in the store
//func (k Keeper) HasIssueFeatures(ctx sdk.Context, id uint64) bool {
//	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.IssueFeaturesKey))
//	return store.Has(GetIssueFeaturesIDBytes(id))
//}

/*
// GetIssueFeaturesOwner returns the creator of the
func (k Keeper) GetIssueFeaturesOwner(ctx sdk.Context, id uint64) string {
	return k.GetIssueFeatures(ctx, id).Creator
}
*/

// RemoveIssueFeatures removes a issueFeatures from the store
func (k Keeper) RemoveIssueFeatures(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.IssueFeaturesKey))
	store.Delete(GetIssueFeaturesIDBytes(id))
}

// GetAllIssueFeatures returns all issueFeatures
func (k Keeper) GetAllIssueFeatures(ctx sdk.Context) (list []types.IssueFeatures) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.IssueFeaturesKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.IssueFeatures
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetIssueFeaturesIDBytes returns the byte representation of the ID
func GetIssueFeaturesIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetIssueFeaturesIDFromBytes returns ID in uint64 format from a byte array
func GetIssueFeaturesIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
