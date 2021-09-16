package keeper

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/konstellation/konstellation/x/issue/types"
	"github.com/stretchr/testify/assert"
)

func createNCoins(keeper *Keeper, ctx sdk.Context, n int) []types.Coins {
	items := make([]types.Coins, n)
	for i := range items {
		items[i].Creator = "any"
		items[i].Id = keeper.AppendCoins(ctx, items[i])
	}
	return items
}

func TestCoinsGet(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNCoins(keeper, ctx, 10)
	for _, item := range items {
		assert.Equal(t, item, keeper.GetCoins(ctx, item.Id))
	}
}

func TestCoinsExist(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNCoins(keeper, ctx, 10)
	for _, item := range items {
		assert.True(t, keeper.HasCoins(ctx, item.Id))
	}
}

func TestCoinsRemove(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNCoins(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveCoins(ctx, item.Id)
		assert.False(t, keeper.HasCoins(ctx, item.Id))
	}
}

func TestCoinsGetAll(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNCoins(keeper, ctx, 10)
	assert.Equal(t, items, keeper.GetAllCoins(ctx))
}

func TestCoinsCount(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNCoins(keeper, ctx, 10)
	count := uint64(len(items))
	assert.Equal(t, count, keeper.GetCoinsCount(ctx))
}
