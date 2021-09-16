package keeper

/*
import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/konstellation/konstellation/x/oracle/types"
	"github.com/stretchr/testify/assert"
)

func createNExchangeRate(keeper *Keeper, ctx sdk.Context, n int) []types.ExchangeRate {
	items := make([]types.ExchangeRate, n)
	for i := range items {
		items[i].Creator = "any"
		items[i].Id = keeper.AppendExchangeRate(ctx, items[i])
	}
	return items
}

func TestExchangeRateGet(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNExchangeRate(keeper, ctx, 10)
	for _, item := range items {
		assert.Equal(t, item, keeper.GetExchangeRate(ctx, item.Id))
	}
}

func TestExchangeRateExist(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNExchangeRate(keeper, ctx, 10)
	for _, item := range items {
		assert.True(t, keeper.HasExchangeRate(ctx, item.Id))
	}
}

func TestExchangeRateRemove(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNExchangeRate(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveExchangeRate(ctx, item.Id)
		assert.False(t, keeper.HasExchangeRate(ctx, item.Id))
	}
}

func TestExchangeRateGetAll(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNExchangeRate(keeper, ctx, 10)
	assert.Equal(t, items, keeper.GetAllExchangeRate(ctx))
}

func TestExchangeRateCount(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNExchangeRate(keeper, ctx, 10)
	count := uint64(len(items))
	assert.Equal(t, count, keeper.GetExchangeRateCount(ctx))
}
*/
