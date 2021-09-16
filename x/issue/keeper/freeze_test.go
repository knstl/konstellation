package keeper

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/konstellation/konstellation/x/issue/types"
	"github.com/stretchr/testify/assert"
)

func createNFreeze(keeper *Keeper, ctx sdk.Context, n int) []types.Freeze {
	items := make([]types.Freeze, n)
	for i := range items {
		items[i].Creator = "any"
		items[i].Id = keeper.AppendFreeze(ctx, items[i])
	}
	return items
}

func TestFreezeGet(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNFreeze(keeper, ctx, 10)
	for _, item := range items {
		assert.Equal(t, item, keeper.GetFreeze(ctx, item.Id))
	}
}

func TestFreezeExist(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNFreeze(keeper, ctx, 10)
	for _, item := range items {
		assert.True(t, keeper.HasFreeze(ctx, item.Id))
	}
}

func TestFreezeRemove(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNFreeze(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveFreeze(ctx, item.Id)
		assert.False(t, keeper.HasFreeze(ctx, item.Id))
	}
}

func TestFreezeGetAll(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNFreeze(keeper, ctx, 10)
	assert.Equal(t, items, keeper.GetAllFreeze(ctx))
}

func TestFreezeCount(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNFreeze(keeper, ctx, 10)
	count := uint64(len(items))
	assert.Equal(t, count, keeper.GetFreezeCount(ctx))
}
