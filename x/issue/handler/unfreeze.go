package handler

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	abcitypes "github.com/tendermint/tendermint/abci/types"

	"github.com/konstellation/konstellation/x/issue/keeper"
	"github.com/konstellation/konstellation/x/issue/types"
)

func HandleMsgUnfreeze(ctx sdk.Context, k keeper.Keeper, msg *types.MsgUnfreeze) *sdk.Result {
	// Sub fee from sender
	fee := k.GetParams(ctx).UnfreezeFee
	if err := k.ChargeFee(ctx, sdk.AccAddress(msg.Freezer), fee); err != nil {
		return &sdk.Result{Log: err.Error()}
	}

	if err := k.Unfreeze(ctx, sdk.AccAddress(msg.Freezer), sdk.AccAddress(msg.Holder), msg.Denom, msg.Op); err != nil {
		return &sdk.Result{Log: err.Error()}
	}

	events := []abcitypes.Event{}
	for _, event := range ctx.EventManager().Events() {
		evt := abcitypes.Event(event)
		events = append(events, evt)
	}
	return &sdk.Result{Events: events}
}
