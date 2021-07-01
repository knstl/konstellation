package handler

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/konstellation/kn-sdk/x/issue/keeper"
	"github.com/konstellation/kn-sdk/x/issue/types"
)

func HandleMsgDescription(ctx sdk.Context, k keeper.Keeper, msg types.MsgDescription) sdk.Result {
	if err := k.ChangeDescription(ctx, msg.Owner, msg.Denom, msg.Description); err != nil {
		return err.Result()
	}

	return sdk.Result{Events: ctx.EventManager().Events()}
}
