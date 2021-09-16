package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/konstellation/konstellation/x/oracle/types"
)

func (m msgServer) DeleteExchangeRate(goCtx context.Context, msgDeleteExchangeRate *types.MsgDeleteExchangeRate) (*types.MsgDeleteExchangeRateResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	senderAddr, err := sdk.AccAddressFromBech32(msgDeleteExchangeRate.Sender)
	if err != nil {
		return nil, err
	}

	err = m.keeper.DeleteExchangeRate(ctx, senderAddr, msgDeleteExchangeRate.Pair)
	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeDeleteExchangeRate,
			sdk.NewAttribute(types.AttributeKeyPair, msgDeleteExchangeRate.Pair),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msgDeleteExchangeRate.Sender),
		),
	})

	return &types.MsgDeleteExchangeRateResponse{}, nil
}
