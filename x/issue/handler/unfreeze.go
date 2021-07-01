package handler

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/konstellation/kn-sdk/x/issue/keeper"
	"github.com/konstellation/kn-sdk/x/issue/types"
)

func HandleMsgUnfreeze(ctx sdk.Context, k keeper.Keeper, msg types.MsgUnfreeze) sdk.Result {
	// Sub fee from sender
	fee := k.GetParams(ctx).UnfreezeFee
	if err := k.ChargeFee(ctx, msg.Freezer, fee); err != nil {
		return err.Result()
	}

	if err := k.Unfreeze(ctx, msg.Freezer, msg.Holder, msg.Denom, msg.Op); err != nil {
		return err.Result()
	}

	return sdk.Result{Events: ctx.EventManager().Events()}
}
