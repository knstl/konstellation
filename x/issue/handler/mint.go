package handler

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/konstellation/kn-sdk/x/issue/keeper"
	"github.com/konstellation/kn-sdk/x/issue/types"
)

func HandleMsgMint(ctx sdk.Context, k keeper.Keeper, msg types.MsgMint) sdk.Result {
	// Sub fee from sender
	fee := k.GetParams(ctx).MintFee
	if err := k.ChargeFee(ctx, msg.Minter, fee); err != nil {
		return err.Result()
	}

	if err := k.Mint(ctx, msg.Minter, msg.ToAddress, msg.Amount); err != nil {
		return err.Result()
	}

	events := []types.Event{}
	for _, event := range ctx.EventManager().Events() {
		events = append(events, event)
	}
	return sdk.Result{Events: events}
}
