package issue

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/konstellation/konstellation/x/issue/types"
)

func NewHandler(k Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) sdk.Result {
		ctx = ctx.WithEventManager(sdk.NewEventManager())

		switch msg := msg.(type) {
		case types.MsgIssue:
			return handleMsgIssue(ctx, msg, k)

		default:
			errMsg := fmt.Sprintf("unrecognized issue message type: %T", msg)
			return sdk.ErrUnknownRequest(errMsg).Result()
		}
	}
}

// These functions assume everything has been authenticated (ValidateBasic passed, and signatures checked)
func handleMsgIssue(ctx sdk.Context, msg types.MsgIssue, k Keeper) sdk.Result {
	// Sub fee from sender
	//fee := keeper.GetParams(ctx).IssueFee
	//if err := keeper.Fee(ctx, msg.Sender, fee); err != nil {
	//	return err.Result()
	//}

	params, err := types.NewIssueParams(msg.IssueParams)
	if err != nil {
		return types.ErrInvalidIssueParams().Result()
	}

	coinIssue := types.NewCoinIssue(msg.Owner, msg.Issuer, params)

	if err := k.CreateIssue(ctx, coinIssue); err != nil {
		return err.Result()
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Owner.String()),
		),
	)

	return sdk.Result{Events: ctx.EventManager().Events()}
}
