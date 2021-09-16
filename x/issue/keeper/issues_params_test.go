package keeper

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/konstellation/konstellation/x/issue/types"
	"github.com/stretchr/testify/assert"
)

func createNIssuesParams(keeper *Keeper, ctx sdk.Context, n int) []types.IssuesParams {
	items := make([]types.IssuesParams, n)
	for i := range items {
		items[i].Creator = "any"
		items[i].Id = keeper.AppendIssuesParams(ctx, items[i])
	}
	return items
}

func TestIssuesParamsGet(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNIssuesParams(keeper, ctx, 10)
	for _, item := range items {
		assert.Equal(t, item, keeper.GetIssuesParams(ctx, item.Id))
	}
}

func TestIssuesParamsExist(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNIssuesParams(keeper, ctx, 10)
	for _, item := range items {
		assert.True(t, keeper.HasIssuesParams(ctx, item.Id))
	}
}

func TestIssuesParamsRemove(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNIssuesParams(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveIssuesParams(ctx, item.Id)
		assert.False(t, keeper.HasIssuesParams(ctx, item.Id))
	}
}

func TestIssuesParamsGetAll(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNIssuesParams(keeper, ctx, 10)
	assert.Equal(t, items, keeper.GetAllIssuesParams(ctx))
}

func TestIssuesParamsCount(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNIssuesParams(keeper, ctx, 10)
	count := uint64(len(items))
	assert.Equal(t, count, keeper.GetIssuesParamsCount(ctx))
}
