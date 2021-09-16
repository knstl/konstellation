package keeper

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/konstellation/konstellation/x/issue/types"
	"github.com/stretchr/testify/assert"
)

func createNParams(keeper *Keeper, ctx sdk.Context, n int) []types.Params {
	items := make([]types.Params, n)
	for i := range items {
		items[i].Creator = "any"
		items[i].Id = keeper.AppendParams(ctx, items[i])
	}
	return items
}

func TestParamsGet(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNParams(keeper, ctx, 10)
	for _, item := range items {
		assert.Equal(t, item, keeper.GetParams(ctx, item.Id))
	}
}

func TestParamsExist(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNParams(keeper, ctx, 10)
	for _, item := range items {
		assert.True(t, keeper.HasParams(ctx, item.Id))
	}
}

func TestParamsRemove(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNParams(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveParams(ctx, item.Id)
		assert.False(t, keeper.HasParams(ctx, item.Id))
	}
}

func TestParamsGetAll(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNParams(keeper, ctx, 10)
	assert.Equal(t, items, keeper.GetAllParams(ctx))
}

func TestParamsCount(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNParams(keeper, ctx, 10)
	count := uint64(len(items))
	assert.Equal(t, count, keeper.GetParamsCount(ctx))
}
