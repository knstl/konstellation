package keeper

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/konstellation/konstellation/x/issue/types"
	"github.com/stretchr/testify/assert"
)

func createNAddressFreeze(keeper *Keeper, ctx sdk.Context, n int) []types.AddressFreeze {
	items := make([]types.AddressFreeze, n)
	for i := range items {
		items[i].Creator = "any"
		items[i].Id = keeper.AppendAddressFreeze(ctx, items[i])
	}
	return items
}

func TestAddressFreezeGet(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNAddressFreeze(keeper, ctx, 10)
	for _, item := range items {
		assert.Equal(t, item, keeper.GetAddressFreeze(ctx, item.Id))
	}
}

func TestAddressFreezeExist(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNAddressFreeze(keeper, ctx, 10)
	for _, item := range items {
		assert.True(t, keeper.HasAddressFreeze(ctx, item.Id))
	}
}

func TestAddressFreezeRemove(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNAddressFreeze(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveAddressFreeze(ctx, item.Id)
		assert.False(t, keeper.HasAddressFreeze(ctx, item.Id))
	}
}

func TestAddressFreezeGetAll(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNAddressFreeze(keeper, ctx, 10)
	assert.Equal(t, items, keeper.GetAllAddressFreeze(ctx))
}

func TestAddressFreezeCount(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNAddressFreeze(keeper, ctx, 10)
	count := uint64(len(items))
	assert.Equal(t, count, keeper.GetAddressFreezeCount(ctx))
}
