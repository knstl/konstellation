package issue

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/konstellation/kn-sdk/x/issue/handler"
	"github.com/konstellation/kn-sdk/x/issue/types"
)

func NewHandler(k Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) sdk.Result {
		ctx = ctx.WithEventManager(sdk.NewEventManager())

		ctx.EventManager().EmitEvent(
			sdk.NewEvent(
				sdk.EventTypeMessage,
				sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			),
		)

		switch msg := msg.(type) {

		case types.MsgIssueCreate:
			return handler.HandleMsgIssueCreate(ctx, k, msg)
		case types.MsgFeatures:
			return handler.HandleMsgFeatures(ctx, k, msg)
		case types.MsgDescription:
			return handler.HandleMsgDescription(ctx, k, msg)
		case types.MsgTransfer:
			return handler.HandleMsgTransfer(ctx, k, msg)
		case types.MsgTransferFrom:
			return handler.HandleMsgTransferFrom(ctx, k, msg)
		case types.MsgApprove:
			return handler.HandleMsgApprove(ctx, k, msg)
		case types.MsgIncreaseAllowance:
			return handler.HandleMsgIncreaseAllowance(ctx, k, msg)
		case types.MsgDecreaseAllowance:
			return handler.HandleMsgDecreaseAllowance(ctx, k, msg)
		case types.MsgMint:
			return handler.HandleMsgMint(ctx, k, msg)
		case types.MsgBurn:
			return handler.HandleMsgBurn(ctx, k, msg)
		case types.MsgBurnFrom:
			return handler.HandleMsgBurnFrom(ctx, k, msg)
		case types.MsgTransferOwnership:
			return handler.HandleMsgTransferOwnership(ctx, k, msg)
		case types.MsgFreeze:
			return handler.HandleMsgFreeze(ctx, k, msg)
		case types.MsgUnfreeze:
			return handler.HandleMsgUnfreeze(ctx, k, msg)

		default:
			errMsg := fmt.Sprintf("unrecognized issue message type: %T", msg)
			return sdk.ErrUnknownRequest(errMsg).Result()
		}
	}
}
