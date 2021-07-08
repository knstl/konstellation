package handler

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	abcitypes "github.com/tendermint/tendermint/abci/types"

	"github.com/konstellation/konstellation/x/issue/keeper"
	"github.com/konstellation/konstellation/x/issue/types"
)

func HandleMsgIncreaseAllowance(ctx sdk.Context, k keeper.Keeper, msg *types.MsgIncreaseAllowance) *sdk.Result {
	if err := k.IncreaseAllowance(ctx, sdk.AccAddress(msg.Owner), sdk.AccAddress(msg.Spender), msg.Amount.Coins); err != nil {
		return &sdk.Result{Log: err.Error()}
	}

	events := []abcitypes.Event{}
	for _, event := range ctx.EventManager().Events() {
		evt := abcitypes.Event(event)
		events = append(events, evt)
	}
	return &sdk.Result{Events: events}
}
