package handler

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	abcitypes "github.com/tendermint/tendermint/abci/types"

	"github.com/konstellation/konstellation/x/issue/keeper"
	"github.com/konstellation/konstellation/x/issue/types"
)

func HandleMsgFeatures(ctx sdk.Context, k keeper.Keeper, msg *types.MsgFeatures) *sdk.Result {
	if err := k.ChangeFeatures(ctx, msg.Owner, msg.Denom, msg.IssueFeatures); err != nil {
		return &sdk.Result{Log: err.Error()}
	}

	events := []abcitypes.Event{}
	for _, event := range ctx.EventManager().Events() {
		evt := abcitypes.Event(event)
		events = append(events, evt)
	}
	return &sdk.Result{Events: events}
}
