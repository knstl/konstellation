package handler

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	abcitypes "github.com/tendermint/tendermint/abci/types"

	"github.com/konstellation/konstellation/x/issue/keeper"
	"github.com/konstellation/konstellation/x/issue/types"
)

func HandleMsgMint(ctx sdk.Context, k keeper.Keeper, msg *types.MsgMint) *sdk.Result {
	// Sub fee from sender
	fee := k.GetParams(ctx).MintFee
	if err := k.ChargeFee(ctx, msg.Minter, fee); err != nil {
		return &sdk.Result{Log: err.Error()}
	}

	if err := k.Mint(ctx, msg.Minter, msg.ToAddress, msg.Amount); err != nil {
		return &sdk.Result{Log: err.Error()}
	}

	events := []abcitypes.Event{}
	for _, event := range ctx.EventManager().Events() {
		evt := abcitypes.Event(event)
		events = append(events, evt)
	}
	return &sdk.Result{Events: events}
}
