package keeper

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/konstellation/konstellation/x/issue/types"
	"github.com/stretchr/testify/assert"
)

func createNCoinIssueList(keeper *Keeper, ctx sdk.Context, n int) []types.CoinIssueList {
	items := make([]types.CoinIssueList, n)
	for i := range items {
		items[i].Creator = "any"
		items[i].Id = keeper.AppendCoinIssueList(ctx, items[i])
	}
	return items
}

func TestCoinIssueListGet(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNCoinIssueList(keeper, ctx, 10)
	for _, item := range items {
		assert.Equal(t, item, keeper.GetCoinIssueList(ctx, item.Id))
	}
}

func TestCoinIssueListExist(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNCoinIssueList(keeper, ctx, 10)
	for _, item := range items {
		assert.True(t, keeper.HasCoinIssueList(ctx, item.Id))
	}
}

func TestCoinIssueListRemove(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNCoinIssueList(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveCoinIssueList(ctx, item.Id)
		assert.False(t, keeper.HasCoinIssueList(ctx, item.Id))
	}
}

func TestCoinIssueListGetAll(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNCoinIssueList(keeper, ctx, 10)
	assert.Equal(t, items, keeper.GetAllCoinIssueList(ctx))
}

func TestCoinIssueListCount(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNCoinIssueList(keeper, ctx, 10)
	count := uint64(len(items))
	assert.Equal(t, count, keeper.GetCoinIssueListCount(ctx))
}
