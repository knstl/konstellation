package keeper

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/konstellation/konstellation/x/issue/types"
	"github.com/stretchr/testify/assert"
)

func createNIssueParams(keeper *Keeper, ctx sdk.Context, n int) []types.IssueParams {
	items := make([]types.IssueParams, n)
	for i := range items {
		items[i].Creator = "any"
		items[i].Id = keeper.AppendIssueParams(ctx, items[i])
	}
	return items
}

func TestIssueParamsGet(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNIssueParams(keeper, ctx, 10)
	for _, item := range items {
		assert.Equal(t, item, keeper.GetIssueParams(ctx, item.Id))
	}
}

func TestIssueParamsExist(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNIssueParams(keeper, ctx, 10)
	for _, item := range items {
		assert.True(t, keeper.HasIssueParams(ctx, item.Id))
	}
}

func TestIssueParamsRemove(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNIssueParams(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveIssueParams(ctx, item.Id)
		assert.False(t, keeper.HasIssueParams(ctx, item.Id))
	}
}

func TestIssueParamsGetAll(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNIssueParams(keeper, ctx, 10)
	assert.Equal(t, items, keeper.GetAllIssueParams(ctx))
}

func TestIssueParamsCount(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNIssueParams(keeper, ctx, 10)
	count := uint64(len(items))
	assert.Equal(t, count, keeper.GetIssueParamsCount(ctx))
}
