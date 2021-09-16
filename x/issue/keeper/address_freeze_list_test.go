package keeper

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/konstellation/konstellation/x/issue/types"
	"github.com/stretchr/testify/assert"
)

func createNAddressFreezeList(keeper *Keeper, ctx sdk.Context, n int) []types.AddressFreezeList {
	items := make([]types.AddressFreezeList, n)
	for i := range items {
		items[i].Creator = "any"
		items[i].Id = keeper.AppendAddressFreezeList(ctx, items[i])
	}
	return items
}

func TestAddressFreezeListGet(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNAddressFreezeList(keeper, ctx, 10)
	for _, item := range items {
		assert.Equal(t, item, keeper.GetAddressFreezeList(ctx, item.Id))
	}
}

func TestAddressFreezeListExist(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNAddressFreezeList(keeper, ctx, 10)
	for _, item := range items {
		assert.True(t, keeper.HasAddressFreezeList(ctx, item.Id))
	}
}

func TestAddressFreezeListRemove(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNAddressFreezeList(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveAddressFreezeList(ctx, item.Id)
		assert.False(t, keeper.HasAddressFreezeList(ctx, item.Id))
	}
}

func TestAddressFreezeListGetAll(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNAddressFreezeList(keeper, ctx, 10)
	assert.Equal(t, items, keeper.GetAllAddressFreezeList(ctx))
}

func TestAddressFreezeListCount(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNAddressFreezeList(keeper, ctx, 10)
	count := uint64(len(items))
	assert.Equal(t, count, keeper.GetAddressFreezeListCount(ctx))
}
