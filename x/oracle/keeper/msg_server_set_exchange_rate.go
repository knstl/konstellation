package keeper

import (
	"context"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/konstellation/konstellation/x/oracle/types"
)

func (m msgServer) SetExchangeRate(goCtx context.Context, msgSetExchangeRate *types.MsgSetExchangeRate) (*types.MsgSetExchangeRateResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	senderAddr, err := sdk.AccAddressFromBech32(msgSetExchangeRate.Sender)
	if err != nil {
		return nil, err
	}

	err = m.keeper.SetExchangeRate(ctx, senderAddr, msgSetExchangeRate.ExchangeRate)
	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeSetExchangeRate,
			sdk.NewAttribute(types.AttributeKeyPair, msgSetExchangeRate.ExchangeRate.Pair),
			sdk.NewAttribute(types.AttributeKeyRate, msgSetExchangeRate.ExchangeRate.Rate.String()),
			sdk.NewAttribute(types.AttributeKeyDenoms, strings.Join(msgSetExchangeRate.ExchangeRate.Denoms, ",")),
			sdk.NewAttribute(types.AttributeKeyTimestamp, msgSetExchangeRate.ExchangeRate.Timestamp.String()),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msgSetExchangeRate.Sender),
		),
	})

	return &types.MsgSetExchangeRateResponse{}, nil
}
