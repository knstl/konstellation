package keeper

import (
	"context"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/konstellation/konstellation/x/oracle/types"
)

func (m msgServer) SetExchangeRates(goCtx context.Context, msgSetExchangeRates *types.MsgSetExchangeRates) (*types.MsgSetExchangeRatesResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	senderAddr, err := sdk.AccAddressFromBech32(msgSetExchangeRates.Sender)
	if err != nil {
		return nil, err
	}

	err = m.keeper.SetExchangeRates(ctx, senderAddr, msgSetExchangeRates.ExchangeRates)
	if err != nil {
		return nil, err
	}

	allEvents := sdk.Events{
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msgSetExchangeRates.Sender),
		),
	}
	for _, exchangeRate := range msgSetExchangeRates.ExchangeRates {
		events := sdk.Events{
			sdk.NewEvent(
				types.EventTypeSetExchangeRates,
				sdk.NewAttribute(types.AttributeKeyPair, exchangeRate.Pair),
				sdk.NewAttribute(types.AttributeKeyRate, exchangeRate.Rate.String()),
				sdk.NewAttribute(types.AttributeKeyDenoms, strings.Join(exchangeRate.Denoms, ",")),
				sdk.NewAttribute(types.AttributeKeyTimestamp, exchangeRate.Timestamp.String()),
			),
		}
		allEvents.AppendEvents(events)
	}
	ctx.EventManager().EmitEvents(allEvents)

	return &types.MsgSetExchangeRatesResponse{}, nil
}
