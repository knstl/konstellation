package keeper

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/konstellation/konstellation/x/issue/types"
	"github.com/stretchr/testify/assert"
)

func createNAllowance(keeper *Keeper, ctx sdk.Context, n int) []types.Allowance {
	items := make([]types.Allowance, n)
	for i := range items {
		items[i].Creator = "any"
		items[i].Id = keeper.AppendAllowance(ctx, items[i])
	}
	return items
}

func TestAllowanceGet(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNAllowance(keeper, ctx, 10)
	for _, item := range items {
		assert.Equal(t, item, keeper.GetAllowance(ctx, item.Id))
	}
}

func TestAllowanceExist(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNAllowance(keeper, ctx, 10)
	for _, item := range items {
		assert.True(t, keeper.HasAllowance(ctx, item.Id))
	}
}

func TestAllowanceRemove(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNAllowance(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveAllowance(ctx, item.Id)
		assert.False(t, keeper.HasAllowance(ctx, item.Id))
	}
}

func TestAllowanceGetAll(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNAllowance(keeper, ctx, 10)
	assert.Equal(t, items, keeper.GetAllAllowance(ctx))
}

func TestAllowanceCount(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNAllowance(keeper, ctx, 10)
	count := uint64(len(items))
	assert.Equal(t, count, keeper.GetAllowanceCount(ctx))
}
