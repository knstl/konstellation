package handler

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/konstellation/kn-sdk/x/issue/keeper"
	"github.com/konstellation/kn-sdk/x/issue/types"
)

func HandleMsgBurnFrom(ctx sdk.Context, k keeper.Keeper, msg types.MsgBurnFrom) sdk.Result {
	// Sub fee from sender
	fee := k.GetParams(ctx).BurnFromFee
	if err := k.ChargeFee(ctx, msg.Burner, fee); err != nil {
		return err.Result()
	}

	if err := k.BurnFrom(ctx, msg.Burner, msg.FromAddress, msg.Amount); err != nil {
		return err.Result()
	}

	return sdk.Result{Events: ctx.EventManager().Events()}
}
