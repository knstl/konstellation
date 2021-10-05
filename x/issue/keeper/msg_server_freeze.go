package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/konstellation/konstellation/x/issue/types"
)

func (k msgServer) Freeze(goCtx context.Context, msg *types.MsgFreeze) (*types.MsgFreezeResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	freezer, err := sdk.AccAddressFromBech32(msg.Freezer)
	if err != nil {
		return nil, err
	}
	holder, err := sdk.AccAddressFromBech32(msg.Holder)
	if err != nil {
		return nil, err
	}

	fee := k.keeper.GetParams(ctx).FreezeFee
	if err := k.keeper.ChargeFee(ctx, freezer, fee); err != nil {
		return nil, err
	}

	if err := k.keeper.Freeze(ctx, freezer, holder, msg.Denom, msg.Op); err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
		),
	)

	return &types.MsgFreezeResponse{}, nil
}
