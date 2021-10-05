package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/konstellation/konstellation/x/issue/types"
)

func (k msgServer) BurnFrom(goCtx context.Context, msg *types.MsgBurnFrom) (*types.MsgBurnFromResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	burner, err := sdk.AccAddressFromBech32(msg.Burner)
	if err != nil {
		return nil, err
	}

	fromAddr, err := sdk.AccAddressFromBech32(msg.FromAddress)
	if err != nil {
		return nil, err
	}

	// Sub fee from sender
	fee := k.keeper.GetParams(ctx).BurnFromFee
	if err := k.keeper.ChargeFee(ctx, burner, fee); err != nil {
		return nil, err
	}

	if err := k.keeper.BurnFrom(ctx, burner, fromAddr, msg.Amount); err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
		),
	)

	return &types.MsgBurnFromResponse{}, nil
}
