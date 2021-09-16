package keeper

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/konstellation/konstellation/x/issue/types"
	"github.com/stretchr/testify/assert"
)

func createNCoinIssue(keeper *Keeper, ctx sdk.Context, n int) []types.CoinIssue {
	items := make([]types.CoinIssue, n)
	for i := range items {
		items[i].Creator = "any"
		items[i].Id = keeper.AppendCoinIssue(ctx, items[i])
	}
	return items
}

func TestCoinIssueGet(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNCoinIssue(keeper, ctx, 10)
	for _, item := range items {
		assert.Equal(t, item, keeper.GetCoinIssue(ctx, item.Id))
	}
}

func TestCoinIssueExist(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNCoinIssue(keeper, ctx, 10)
	for _, item := range items {
		assert.True(t, keeper.HasCoinIssue(ctx, item.Id))
	}
}

func TestCoinIssueRemove(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNCoinIssue(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveCoinIssue(ctx, item.Id)
		assert.False(t, keeper.HasCoinIssue(ctx, item.Id))
	}
}

func TestCoinIssueGetAll(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNCoinIssue(keeper, ctx, 10)
	assert.Equal(t, items, keeper.GetAllCoinIssue(ctx))
}

func TestCoinIssueCount(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNCoinIssue(keeper, ctx, 10)
	count := uint64(len(items))
	assert.Equal(t, count, keeper.GetCoinIssueCount(ctx))
}
