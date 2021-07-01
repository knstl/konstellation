package handler

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/konstellation/kn-sdk/x/issue/keeper"
	"github.com/konstellation/kn-sdk/x/issue/types"
)

func HandleMsgIncreaseAllowance(ctx sdk.Context, k keeper.Keeper, msg types.MsgIncreaseAllowance) sdk.Result {
	if err := k.IncreaseAllowance(ctx, msg.Owner, msg.Spender, msg.Amount); err != nil {
		return err.Result()
	}

	return sdk.Result{Events: ctx.EventManager().Events()}
}
