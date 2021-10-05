package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/konstellation/konstellation/x/issue/types"
)

func (k msgServer) Issue(goCtx context.Context, msg *types.MsgIssue) (*types.MsgIssueResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	issuerAddr, err := sdk.AccAddressFromBech32(msg.Issuer)
	if err != nil {
		return nil, err
	}

	ownerAddr, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		return nil, err
	}

	// Sub fee from issuer
	fee := k.keeper.GetParams(ctx).IssueFee
	if err := k.keeper.ChargeFee(ctx, issuerAddr, fee); err != nil {
		return nil, err
	}

	params, err := types.NewIssueParams(msg.IssueParams)
	if err != nil {
		return nil, types.ErrInvalidIssueParams
	}

	ci := k.keeper.CreateIssue(ctx, ownerAddr, issuerAddr, params)
	if err := k.keeper.Issue(ctx, ci); err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
		),
	)

	return &types.MsgIssueResponse{}, nil
}
