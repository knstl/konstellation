package keeper

//func createNCoinIssueDenom(keeper *Keeper, ctx sdk.Context, n int) []types.CoinIssueDenom {
//	items := make([]types.CoinIssueDenom, n)
//	for i := range items {
//		items[i].Creator = "any"
//		items[i].Id = keeper.AppendCoinIssueDenom(ctx, items[i])
//	}
//	return items
//}
//
//func TestCoinIssueDenomGet(t *testing.T) {
//	keeper, ctx := setupKeeper(t)
//	items := createNCoinIssueDenom(keeper, ctx, 10)
//	for _, item := range items {
//		assert.Equal(t, item, keeper.GetCoinIssueDenom(ctx, item.Id))
//	}
//}
//
//func TestCoinIssueDenomExist(t *testing.T) {
//	keeper, ctx := setupKeeper(t)
//	items := createNCoinIssueDenom(keeper, ctx, 10)
//	for _, item := range items {
//		assert.True(t, keeper.HasCoinIssueDenom(ctx, item.Id))
//	}
//}
//
//func TestCoinIssueDenomRemove(t *testing.T) {
//	keeper, ctx := setupKeeper(t)
//	items := createNCoinIssueDenom(keeper, ctx, 10)
//	for _, item := range items {
//		keeper.RemoveCoinIssueDenom(ctx, item.Id)
//		assert.False(t, keeper.HasCoinIssueDenom(ctx, item.Id))
//	}
//}
//
//func TestCoinIssueDenomGetAll(t *testing.T) {
//	keeper, ctx := setupKeeper(t)
//	items := createNCoinIssueDenom(keeper, ctx, 10)
//	assert.Equal(t, items, keeper.GetAllCoinIssueDenom(ctx))
//}
//
//func TestCoinIssueDenomCount(t *testing.T) {
//	keeper, ctx := setupKeeper(t)
//	items := createNCoinIssueDenom(keeper, ctx, 10)
//	count := uint64(len(items))
//	assert.Equal(t, count, keeper.GetCoinIssueDenomCount(ctx))
//}
