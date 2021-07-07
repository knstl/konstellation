package handler

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/konstellation/kn-sdk/x/issue/keeper"
	"github.com/konstellation/kn-sdk/x/issue/types"
)

func HandleMsgApprove(ctx sdk.Context, k keeper.Keeper, msg types.MsgApprove) sdk.Result {
	if err := k.Approve(ctx, msg.Owner, msg.Spender, msg.Amount); err != nil {
		return err.Result()
	}

	events := []types.Event{}
	for _, event := range ctx.EventManager().Events() {
		events = append(events, event)
	}
	return sdk.Result{Events: events}
}
