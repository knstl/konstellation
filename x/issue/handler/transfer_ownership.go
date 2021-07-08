package handler

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	abcitypes "github.com/tendermint/tendermint/abci/types"

	"github.com/konstellation/konstellation/x/issue/keeper"
	"github.com/konstellation/konstellation/x/issue/types"
)

func HandleMsgTransferOwnership(ctx sdk.Context, k keeper.Keeper, msg *types.MsgTransferOwnership) *sdk.Result {
	// Sub fee from sender
	fee := k.GetParams(ctx).TransferOwnerFee
	if err := k.ChargeFee(ctx, sdk.AccAddress(msg.Owner), fee); err != nil {
		return &sdk.Result{Log: err.Error()}
	}

	if err := k.TransferOwnership(ctx, sdk.AccAddress(msg.Owner), sdk.AccAddress(msg.ToAddress), msg.Denom); err != nil {
		return &sdk.Result{Log: err.Error()}
	}

	events := []abcitypes.Event{}
	for _, event := range ctx.EventManager().Events() {
		evt := abcitypes.Event(event)
		events = append(events, evt)
	}
	return &sdk.Result{Events: events}
}
