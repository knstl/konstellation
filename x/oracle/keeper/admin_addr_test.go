package keeper

/*
import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/konstellation/konstellation/x/oracle/types"
	"github.com/stretchr/testify/assert"
)

func createNAdminAddr(keeper *Keeper, ctx sdk.Context, n int) []types.AdminAddr {
	items := make([]types.AdminAddr, n)
	for i := range items {
		items[i].Creator = "any"
		items[i].Id = keeper.AppendAdminAddr(ctx, items[i])
	}
	return items
}

func TestAdminAddrGet(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNAdminAddr(keeper, ctx, 10)
	for _, item := range items {
		assert.Equal(t, item, keeper.GetAdminAddr(ctx, item.Id))
	}
}

func TestAdminAddrExist(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNAdminAddr(keeper, ctx, 10)
	for _, item := range items {
		assert.True(t, keeper.HasAdminAddr(ctx, item.Id))
	}
}

func TestAdminAddrRemove(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNAdminAddr(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveAdminAddr(ctx, item.Id)
		assert.False(t, keeper.HasAdminAddr(ctx, item.Id))
	}
}

func TestAdminAddrGetAll(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNAdminAddr(keeper, ctx, 10)
	assert.Equal(t, items, keeper.GetAllAdminAddr(ctx))
}

func TestAdminAddrCount(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNAdminAddr(keeper, ctx, 10)
	count := uint64(len(items))
	assert.Equal(t, count, keeper.GetAdminAddrCount(ctx))
}
*/
