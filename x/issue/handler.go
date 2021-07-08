package issue

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/konstellation/konstellation/x/issue/handler"
	"github.com/konstellation/konstellation/x/issue/types"
)

func NewHandler(k Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		ctx = ctx.WithEventManager(sdk.NewEventManager())

		ctx.EventManager().EmitEvent(
			sdk.NewEvent(
				sdk.EventTypeMessage,
				sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			),
		)

		switch msg := msg.(type) {

		case *types.MsgIssueCreate:
			return handler.HandleMsgIssueCreate(ctx, k, msg), nil
		case *types.MsgFeatures:
			return handler.HandleMsgFeatures(ctx, k, msg), nil
		case *types.MsgDescription:
			return handler.HandleMsgDescription(ctx, k, msg), nil
		case *types.MsgTransfer:
			return handler.HandleMsgTransfer(ctx, k, msg), nil
		case *types.MsgTransferFrom:
			return handler.HandleMsgTransferFrom(ctx, k, msg), nil
		case *types.MsgApprove:
			return handler.HandleMsgApprove(ctx, k, msg), nil
		case *types.MsgIncreaseAllowance:
			return handler.HandleMsgIncreaseAllowance(ctx, k, msg), nil
		case *types.MsgDecreaseAllowance:
			return handler.HandleMsgDecreaseAllowance(ctx, k, msg), nil
		case *types.MsgMint:
			return handler.HandleMsgMint(ctx, k, msg), nil
		case *types.MsgBurn:
			return handler.HandleMsgBurn(ctx, k, msg), nil
		case *types.MsgBurnFrom:
			return handler.HandleMsgBurnFrom(ctx, k, msg), nil
		case *types.MsgTransferOwnership:
			return handler.HandleMsgTransferOwnership(ctx, k, msg), nil
		case *types.MsgFreeze:
			return handler.HandleMsgFreeze(ctx, k, msg), nil
		case *types.MsgUnfreeze:
			return handler.HandleMsgUnfreeze(ctx, k, msg), nil

		default:
			errMsg := fmt.Sprintf("unrecognized issue message type: %T", msg)
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, errMsg)
		}
	}
}
