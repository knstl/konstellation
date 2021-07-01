package handler

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/konstellation/kn-sdk/x/issue/keeper"
	"github.com/konstellation/kn-sdk/x/issue/types"
)

func HandleMsgBurn(ctx sdk.Context, k keeper.Keeper, msg types.MsgBurn) sdk.Result {
	// Sub fee from sender
	fee := k.GetParams(ctx).BurnFee
	if err := k.ChargeFee(ctx, msg.Burner, fee); err != nil {
		return err.Result()
	}

	if err := k.Burn(ctx, msg.Burner, msg.Amount); err != nil {
		return err.Result()
	}

	return sdk.Result{Events: ctx.EventManager().Events()}
}
