package handler

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/konstellation/konstellation/x/issue/keeper"
	"github.com/konstellation/konstellation/x/issue/types"
)

func HandleMsgApprove(ctx sdk.Context, k keeper.Keeper, msg types.MsgApprove) sdk.Result {
	// TODO set fee

	if err := k.Approve(ctx, msg.Owner, msg.Spender, msg.Amount); err != nil {
		return err.Result()
	}

	return sdk.Result{Events: ctx.EventManager().Events()}
}
