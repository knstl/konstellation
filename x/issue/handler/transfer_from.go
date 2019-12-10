package handler

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/konstellation/konstellation/x/issue/keeper"
	"github.com/konstellation/konstellation/x/issue/types"
)

func HandleMsgTransferFrom(ctx sdk.Context, k keeper.Keeper, msg types.MsgTransferFrom) sdk.Result {
	// Sub fee from sender
	//fee := keeper.GetParams(ctx).IssueFee
	//if err := keeper.Fee(ctx, msg.Sender, fee); err != nil {
	//	return err.Result()
	//} sdk.AccAddress, amt sdk.Coin{}

	if err := k.TransferFrom(ctx, msg.Sender, msg.FromAddress, msg.ToAddress, msg.Amount); err != nil {
		return err.Result()
	}

	return sdk.Result{Events: ctx.EventManager().Events()}
}
