package keeper

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/konstellation/konstellation/x/issue/types"
	"github.com/stretchr/testify/assert"
)

func createNCoinIssueDenoms(keeper *Keeper, ctx sdk.Context, n int) []types.CoinIssueDenoms {
	items := make([]types.CoinIssueDenoms, n)
	for i := range items {
		items[i].Creator = "any"
		items[i].Id = keeper.AppendCoinIssueDenoms(ctx, items[i])
	}
	return items
}

func TestCoinIssueDenomsGet(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNCoinIssueDenoms(keeper, ctx, 10)
	for _, item := range items {
		assert.Equal(t, item, keeper.GetCoinIssueDenoms(ctx, item.Id))
	}
}

func TestCoinIssueDenomsExist(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNCoinIssueDenoms(keeper, ctx, 10)
	for _, item := range items {
		assert.True(t, keeper.HasCoinIssueDenoms(ctx, item.Id))
	}
}

func TestCoinIssueDenomsRemove(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNCoinIssueDenoms(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveCoinIssueDenoms(ctx, item.Id)
		assert.False(t, keeper.HasCoinIssueDenoms(ctx, item.Id))
	}
}

func TestCoinIssueDenomsGetAll(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNCoinIssueDenoms(keeper, ctx, 10)
	assert.Equal(t, items, keeper.GetAllCoinIssueDenoms(ctx))
}

func TestCoinIssueDenomsCount(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNCoinIssueDenoms(keeper, ctx, 10)
	count := uint64(len(items))
	assert.Equal(t, count, keeper.GetCoinIssueDenomsCount(ctx))
}
