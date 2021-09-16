package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/konstellation/konstellation/x/oracle/types"
)

func (m msgServer) DeleteExchangeRates(goCtx context.Context, msgDeleteExchangeRates *types.MsgDeleteExchangeRates) (*types.MsgDeleteExchangeRatesResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	senderAddr, err := sdk.AccAddressFromBech32(msgDeleteExchangeRates.Sender)
	if err != nil {
		return nil, err
	}

	err = m.keeper.DeleteExchangeRates(ctx, senderAddr, msgDeleteExchangeRates.Pairs)
	if err != nil {
		return nil, err
	}

	allEvents := sdk.Events{
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msgDeleteExchangeRates.Sender),
		),
	}
	for _, pair := range msgDeleteExchangeRates.Pairs {
		events := sdk.Events{
			sdk.NewEvent(
				types.EventTypeDeleteExchangeRates,
				sdk.NewAttribute(types.AttributeKeyPair, pair),
			),
		}
		allEvents.AppendEvents(events)
	}
	ctx.EventManager().EmitEvents(allEvents)

	return &types.MsgDeleteExchangeRatesResponse{}, nil
}
