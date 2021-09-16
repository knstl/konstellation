package keeper

//func createNIssueFeatures(keeper *Keeper, ctx sdk.Context, n int) []types.IssueFeatures {
//	items := make([]types.IssueFeatures, n)
//	for i := range items {
//		items[i].Creator = "any"
//		items[i].Id = keeper.AppendIssueFeatures(ctx, items[i])
//	}
//	return items
//}
//
//func TestIssueFeaturesGet(t *testing.T) {
//	keeper, ctx := setupKeeper(t)
//	items := createNIssueFeatures(keeper, ctx, 10)
//	for _, item := range items {
//		assert.Equal(t, item, keeper.GetIssueFeatures(ctx, item.Id))
//	}
//}
//
//func TestIssueFeaturesExist(t *testing.T) {
//	keeper, ctx := setupKeeper(t)
//	items := createNIssueFeatures(keeper, ctx, 10)
//	for _, item := range items {
//		assert.True(t, keeper.HasIssueFeatures(ctx, item.Id))
//	}
//}
//
//func TestIssueFeaturesRemove(t *testing.T) {
//	keeper, ctx := setupKeeper(t)
//	items := createNIssueFeatures(keeper, ctx, 10)
//	for _, item := range items {
//		keeper.RemoveIssueFeatures(ctx, item.Id)
//		assert.False(t, keeper.HasIssueFeatures(ctx, item.Id))
//	}
//}
//
//func TestIssueFeaturesGetAll(t *testing.T) {
//	keeper, ctx := setupKeeper(t)
//	items := createNIssueFeatures(keeper, ctx, 10)
//	assert.Equal(t, items, keeper.GetAllIssueFeatures(ctx))
//}
//
//func TestIssueFeaturesCount(t *testing.T) {
//	keeper, ctx := setupKeeper(t)
//	items := createNIssueFeatures(keeper, ctx, 10)
//	count := uint64(len(items))
//	assert.Equal(t, count, keeper.GetIssueFeaturesCount(ctx))
//}
