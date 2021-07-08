package handler

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/konstellation/konstellation/x/issue/keeper"
	"github.com/konstellation/konstellation/x/issue/types"
	abcitypes "github.com/tendermint/tendermint/abci/types"
)

func HandleMsgApprove(ctx sdk.Context, k keeper.Keeper, msg *types.MsgApprove) *sdk.Result {
	if err := k.Approve(ctx, sdk.AccAddress(msg.Owner), sdk.AccAddress(msg.Spender), msg.Amount.Coins); err != nil {
		return &sdk.Result{Log: err.Error()}
	}

	events := []abcitypes.Event{}
	for _, event := range ctx.EventManager().Events() {
		evt := abcitypes.Event(event)
		events = append(events, evt)
	}
	return &sdk.Result{Events: events}
}
