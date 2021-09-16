package keeper

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/konstellation/konstellation/x/issue/types"
	"github.com/stretchr/testify/assert"
)

func createNAllowanceList(keeper *Keeper, ctx sdk.Context, n int) []types.AllowanceList {
	items := make([]types.AllowanceList, n)
	for i := range items {
		items[i].Creator = "any"
		items[i].Id = keeper.AppendAllowanceList(ctx, items[i])
	}
	return items
}

func TestAllowanceListGet(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNAllowanceList(keeper, ctx, 10)
	for _, item := range items {
		assert.Equal(t, item, keeper.GetAllowanceList(ctx, item.Id))
	}
}

func TestAllowanceListExist(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNAllowanceList(keeper, ctx, 10)
	for _, item := range items {
		assert.True(t, keeper.HasAllowanceList(ctx, item.Id))
	}
}

func TestAllowanceListRemove(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNAllowanceList(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveAllowanceList(ctx, item.Id)
		assert.False(t, keeper.HasAllowanceList(ctx, item.Id))
	}
}

func TestAllowanceListGetAll(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNAllowanceList(keeper, ctx, 10)
	assert.Equal(t, items, keeper.GetAllAllowanceList(ctx))
}

func TestAllowanceListCount(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNAllowanceList(keeper, ctx, 10)
	count := uint64(len(items))
	assert.Equal(t, count, keeper.GetAllowanceListCount(ctx))
}
