package keeper

import (
	"context"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/konstellation/konstellation/x/oracle/types"
)

var _ types.MsgServer = msgServer{}

type msgServer struct {
	keeper Keeper
}

func NewMsgServerImpl(k Keeper) types.MsgServer {
	return &msgServer{keeper: k}
}

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

	allEvents := sdk.Events{}
	for _, exchangeRate := range msgSetExchangeRates.ExchangeRates {
		events := sdk.Events{
			sdk.NewEvent(
				types.EventTypeSetExchangeRates,
				sdk.NewAttribute(types.AttributeKeyPair, exchangeRate.Pair),
				sdk.NewAttribute(types.AttributeKeyRate, exchangeRate.Rate.String()),
				sdk.NewAttribute(types.AttributeKeyDenoms, strings.Join(exchangeRate.Denoms, ",")),
				sdk.NewAttribute(types.AttributeKeyTimestamp, exchangeRate.Timestamp.String()),
			),
			sdk.NewEvent(
				sdk.EventTypeMessage,
				sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
				sdk.NewAttribute(sdk.AttributeKeySender, msgSetExchangeRates.Sender),
			)}
		allEvents.AppendEvents(events)
	}
	ctx.EventManager().EmitEvents(sdk.Events(allEvents))

	return &types.MsgSetExchangeRatesResponse{}, nil
}

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

	allEvents := sdk.Events{}
	for _, pair := range msgDeleteExchangeRates.Pairs {
		events := sdk.Events{
			sdk.NewEvent(
				types.EventTypeDeleteExchangeRates,
				sdk.NewAttribute(types.AttributeKeyPair, pair),
			),
			sdk.NewEvent(
				sdk.EventTypeMessage,
				sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
				sdk.NewAttribute(sdk.AttributeKeySender, msgDeleteExchangeRates.Sender),
			)}
		allEvents.AppendEvents(events)
	}
	ctx.EventManager().EmitEvents(sdk.Events(allEvents))

	return &types.MsgDeleteExchangeRatesResponse{}, nil
}

func (m msgServer) SetAdminAddr(goCtx context.Context, msgSetAdminAddr *types.MsgSetAdminAddr) (*types.MsgSetAdminAddrResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	senderAddr, err := sdk.AccAddressFromBech32(msgSetAdminAddr.Sender)
	if err != nil {
		return nil, err
	}

	var add []types.AdminAddr
	for _, a := range msgSetAdminAddr.Add {
		_, err := sdk.AccAddressFromBech32(a.GetAddress())
		if err != nil {
			return nil, err
		}

		add = append(add, *a)
	}
	var del []types.AdminAddr
	for _, a := range msgSetAdminAddr.Delete {
		_, err := sdk.AccAddressFromBech32(a.GetAddress())
		if err != nil {
			return nil, err
		}

		del = append(del, *a)
	}

	if err := m.keeper.SetAdminAddr(ctx, senderAddr, add, del); err != nil {
		return nil, err
	}

	addAddresses := []string{}
	for _, add := range msgSetAdminAddr.Add {
		addAddresses = append(addAddresses, add.Address)
	}
	deleteAddresses := []string{}
	for _, del := range msgSetAdminAddr.Delete {
		addAddresses = append(addAddresses, del.Address)
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeSetAdminAddr,
			sdk.NewAttribute(types.AttributeKeyAdd, strings.Join(addAddresses, ",")),
			sdk.NewAttribute(types.AttributeKeyDelete, strings.Join(deleteAddresses, ",")),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msgSetAdminAddr.Sender),
		),
	})

	return &types.MsgSetAdminAddrResponse{}, nil
}
