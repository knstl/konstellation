package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/konstellation/konstellation/x/issue/types"
)

func (k msgServer) Burn(goCtx context.Context, msg *types.MsgBurn) (*types.MsgBurnResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	burner, err := sdk.AccAddressFromBech32(msg.Burner)
	if err != nil {
		return nil, err
	}

	// Sub fee from sender
	fee := k.keeper.GetParams(ctx).BurnFee
	if err := k.keeper.ChargeFee(ctx, burner, fee); err != nil {
		return nil, err
	}

	if err := k.keeper.Burn(ctx, burner, msg.Amount); err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
		),
	)

	return &types.MsgBurnResponse{}, nil
}
