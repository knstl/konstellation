package handler

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	abcitypes "github.com/tendermint/tendermint/abci/types"

	"github.com/konstellation/konstellation/x/issue/keeper"
	"github.com/konstellation/konstellation/x/issue/types"
)

func HandleMsgBurnFrom(ctx sdk.Context, k keeper.Keeper, msg *types.MsgBurnFrom) *sdk.Result {
	// Sub fee from sender
	fee := k.GetParams(ctx).BurnFromFee
	if err := k.ChargeFee(ctx, msg.Burner, fee); err != nil {
		return &sdk.Result{Log: err.Error()}
	}

	if err := k.BurnFrom(ctx, msg.Burner, msg.FromAddress, msg.Amount); err != nil {
		return &sdk.Result{Log: err.Error()}
	}

	events := []abcitypes.Event{}
	for _, event := range ctx.EventManager().Events() {
		evt := abcitypes.Event(event)
		events = append(events, evt)
	}
	return &sdk.Result{Events: events}
}
