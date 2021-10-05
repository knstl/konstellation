package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/konstellation/konstellation/x/issue/types"
)

func (k msgServer) Mint(goCtx context.Context, msg *types.MsgMint) (*types.MsgMintResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	minter, err := sdk.AccAddressFromBech32(msg.Minter)
	if err != nil {
		return nil, err
	}

	toAddr, err := sdk.AccAddressFromBech32(msg.ToAddress)
	if err != nil {
		return nil, err
	}

	fee := k.keeper.GetParams(ctx).MintFee
	if err := k.keeper.ChargeFee(ctx, minter, fee); err != nil {
		return nil, err
	}

	if err := k.keeper.Mint(ctx, minter, toAddr, msg.Amount); err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
		),
	)

	return &types.MsgMintResponse{}, nil
}
