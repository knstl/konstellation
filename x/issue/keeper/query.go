package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

//func QueryParams(ctx sdk.Context, keeper Keeper) ([]byte, sdk.Error) {
//	params := keeper.GetParams(ctx)
//	bz, err := codec.MarshalJSONIndent(keeper.Getcdc(), params)
//	if err != nil {
//		return nil, sdk.ErrInternal(sdk.AppendMsgToErr("could not marshal result to JSON", err.Error()))
//	}
//	return bz, nil
//}
func QueryIssue(ctx sdk.Context, k Keeper, issueID string) ([]byte, sdk.Error) {
	//issue := keeper.GetIssue(ctx, issueID)
	//if issue == nil {
	//	return nil, errors.ErrUnknownIssue(issueID)
	//}

	//coins, err := keeper.CreateIssue(ctx)
	//if err != nil {
	//	return nil, err
	//}
	//
	//bz, errr := codec.MarshalJSONIndent(keeper.cdc, &coins)
	//if errr != nil {
	//	return nil, sdk.ErrInternal(sdk.AppendMsgToErr("could not marshal result to JSON", err.Error()))
	//}
	return nil, nil
}
