package handler

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/konstellation/kn-sdk/x/issue/keeper"
	"github.com/konstellation/kn-sdk/x/issue/types"
)

func HandleMsgFeatures(ctx sdk.Context, k keeper.Keeper, msg types.MsgFeatures) sdk.Result {
	if err := k.ChangeFeatures(ctx, msg.Owner, msg.Denom, msg.IssueFeatures); err != nil {
		return err.Result()
	}

	events := []types.Event{}
	for _, event := range ctx.EventManager().Events() {
		events = append(events, event)
	}
	return sdk.Result{Events: events}
}
