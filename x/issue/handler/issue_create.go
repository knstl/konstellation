package handler

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/konstellation/kn-sdk/x/issue/keeper"
	"github.com/konstellation/kn-sdk/x/issue/types"
)

func HandleMsgIssueCreate(ctx sdk.Context, k keeper.Keeper, msg types.MsgIssueCreate) sdk.Result {
	// Sub fee from issuer
	fee := k.GetParams(ctx).IssueFee
	if err := k.ChargeFee(ctx, msg.Issuer, fee); err != nil {
		return err.Result()
	}

	params, errr := types.NewIssueParams(msg.IssueParams)
	if errr != nil {
		return types.ErrInvalidIssueParams.Result()
	}

	ci := k.CreateIssue(ctx, msg.Owner, msg.Issuer, params)
	if err := k.Issue(ctx, ci); err != nil {
		return err.Result()
	}

	events := []types.Event{}
	for _, event := range ctx.EventManager().Events() {
		events = append(events, event)
	}
	return sdk.Result{Events: events}
}
