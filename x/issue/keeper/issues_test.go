package keeper

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/konstellation/konstellation/x/issue/types"
	"github.com/stretchr/testify/assert"
)

func createNIssues(keeper *Keeper, ctx sdk.Context, n int) []types.Issues {
	items := make([]types.Issues, n)
	//for i := range items {
	//	//items[i].Creator = "any"
	//	//items[i].Id = keeper.AppendIssues(ctx, items[i])
	//}
	return items
}

func TestIssuesGet(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNIssues(keeper, ctx, 10)
	for _, item := range items {
		assert.Equal(t, item, keeper.GetIssues(ctx, item.Id))
	}
}

func TestIssuesExist(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNIssues(keeper, ctx, 10)
	for _, item := range items {
		assert.True(t, keeper.HasIssues(ctx, item.Id))
	}
}

func TestIssuesRemove(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNIssues(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveIssues(ctx, item.Id)
		assert.False(t, keeper.HasIssues(ctx, item.Id))
	}
}

func TestIssuesGetAll(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNIssues(keeper, ctx, 10)
	assert.Equal(t, items, keeper.GetAllIssues(ctx))
}

func TestIssuesCount(t *testing.T) {
	//keeper, ctx := setupKeeper(t)
	//items := createNIssues(keeper, ctx, 10)
	//count := uint64(len(items))
	//assert.Equal(t, count, keeper.GetIssuesCount(ctx))
}
