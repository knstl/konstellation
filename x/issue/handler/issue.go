package handler

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/konstellation/konstellation/x/issue/keeper"
	"github.com/konstellation/konstellation/x/issue/types"
)

func HandleMsgIssueCreate(ctx sdk.Context, k keeper.Keeper, msg types.MsgIssueCreate) sdk.Result {
	// Sub fee from sender
	//fee := keeper.GetParams(ctx).IssueFee
	//if err := keeper.Fee(ctx, msg.Sender, fee); err != nil {
	//	return err.Result()
	//}

	params, errr := types.NewIssueParams(msg.IssueParams)
	if errr != nil {
		return types.ErrInvalidIssueParams().Result()
	}

	ci := k.CreateIssue(ctx, msg.Owner, msg.Issuer, params)
	if err := k.Issue(ctx, ci); err != nil {
		return err.Result()
	}

	return sdk.Result{Events: ctx.EventManager().Events()}
}
