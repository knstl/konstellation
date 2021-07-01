package handler

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/konstellation/kn-sdk/x/issue/keeper"
	"github.com/konstellation/kn-sdk/x/issue/types"
)

func HandleMsgTransferFrom(ctx sdk.Context, k keeper.Keeper, msg types.MsgTransferFrom) sdk.Result {
	if err := k.TransferFrom(ctx, msg.Sender, msg.FromAddress, msg.ToAddress, msg.Amount); err != nil {
		return err.Result()
	}

	return sdk.Result{Events: ctx.EventManager().Events()}
}
