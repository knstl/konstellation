package keeper

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/konstellation/konstellation/x/issue/types"
	"github.com/stretchr/testify/assert"
)

func createNAddress(keeper *Keeper, ctx sdk.Context, n int) []types.Address {
	items := make([]types.Address, n)
	for i := range items {
		items[i].Creator = "any"
		items[i].Id = keeper.AppendAddress(ctx, items[i])
	}
	return items
}

func TestAddressGet(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNAddress(keeper, ctx, 10)
	for _, item := range items {
		assert.Equal(t, item, keeper.GetAddress(ctx, item.Id))
	}
}

func TestAddressExist(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNAddress(keeper, ctx, 10)
	for _, item := range items {
		assert.True(t, keeper.HasAddress(ctx, item.Id))
	}
}

func TestAddressRemove(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNAddress(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveAddress(ctx, item.Id)
		assert.False(t, keeper.HasAddress(ctx, item.Id))
	}
}

func TestAddressGetAll(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNAddress(keeper, ctx, 10)
	assert.Equal(t, items, keeper.GetAllAddress(ctx))
}

func TestAddressCount(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNAddress(keeper, ctx, 10)
	count := uint64(len(items))
	assert.Equal(t, count, keeper.GetAddressCount(ctx))
}
