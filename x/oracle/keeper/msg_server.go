package keeper

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/konstellation/konstellation/x/oracle/types"
)

var _ types.MsgServer = msgServer{}

const (
	AttributeKeyPair      = "pair"
	AttributeKeyRate      = "rate"
	AttributeKeyDenoms    = "denoms"
	AttributeKeyHeight    = "height"
	AttributeKeyTimestamp = "timestamp"

	AttributeKeyFunction        = "function"
	AttributeKeyFunctionValue1  = "SetExchangeRate"
	AttributeKeyFunctionValue2  = "SetExchangeRates"
	AttributeKeyFunctionValue3  = "DeleteExchangeRate"
	AttributeKeyFunctionValue4  = "DeteteExchangeRates"
	AttributeKeyFunctionValue5  = "SetAdminAddr"
	AttributeKeyAdminAddrAdd    = "add"
	AttributeKeyAdminAddrDelete = "delete"
)

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
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(AttributeKeyFunction, AttributeKeyFunctionValue1),
			sdk.NewAttribute(sdk.AttributeKeySender, msgSetExchangeRate.Sender),
			sdk.NewAttribute(AttributeKeyPair, msgSetExchangeRate.ExchangeRate.Pair),
			sdk.NewAttribute(AttributeKeyRate, msgSetExchangeRate.ExchangeRate.Rate.String()),
			sdk.NewAttribute(AttributeKeyDenoms, strings.Join(msgSetExchangeRate.ExchangeRate.Denoms, ",")),
			sdk.NewAttribute(AttributeKeyHeight, strconv.FormatInt(msgSetExchangeRate.ExchangeRate.Height, 10)),
			sdk.NewAttribute(AttributeKeyTimestamp, msgSetExchangeRate.ExchangeRate.Timestamp.String()),
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

	events := []sdk.Event{}
	for _, exchangeRate := range msgSetExchangeRates.ExchangeRates {
		events = append(events,
			sdk.NewEvent(
				sdk.EventTypeMessage,
				sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
				sdk.NewAttribute(AttributeKeyFunction, AttributeKeyFunctionValue2),
				sdk.NewAttribute(sdk.AttributeKeySender, msgSetExchangeRates.Sender),
				sdk.NewAttribute(AttributeKeyPair, exchangeRate.Pair),
				sdk.NewAttribute(AttributeKeyRate, exchangeRate.Rate.String()),
				sdk.NewAttribute(AttributeKeyDenoms, strings.Join(exchangeRate.Denoms, ",")),
				sdk.NewAttribute(AttributeKeyHeight, strconv.FormatInt(exchangeRate.Height, 10)),
				sdk.NewAttribute(AttributeKeyTimestamp, exchangeRate.Timestamp.String()),
			))
	}
	ctx.EventManager().EmitEvents(sdk.Events(events))

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
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(AttributeKeyFunction, AttributeKeyFunctionValue3),
			sdk.NewAttribute(sdk.AttributeKeySender, msgDeleteExchangeRate.Sender),
			sdk.NewAttribute(AttributeKeyPair, msgDeleteExchangeRate.Pair),
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

	events := []sdk.Event{}
	for _, pair := range msgDeleteExchangeRates.Pairs {
		events = append(events,
			sdk.NewEvent(
				sdk.EventTypeMessage,
				sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
				sdk.NewAttribute(AttributeKeyFunction, AttributeKeyFunctionValue4),
				sdk.NewAttribute(sdk.AttributeKeySender, msgDeleteExchangeRates.Sender),
				sdk.NewAttribute(AttributeKeyPair, pair),
			))
	}
	ctx.EventManager().EmitEvents(sdk.Events(events))

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

	event := sdk.NewEvent(
		sdk.EventTypeMessage,
		sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
		sdk.NewAttribute(AttributeKeyFunction, AttributeKeyFunctionValue5),
		sdk.NewAttribute(sdk.AttributeKeySender, msgSetAdminAddr.Sender),
	)
	for i, addrAdd := range msgSetAdminAddr.Add {
		event.AppendAttributes(sdk.NewAttribute(fmt.Sprintf("%s%d", AttributeKeyAdminAddrAdd, i), addrAdd.String()))
	}
	for i, addrDelete := range msgSetAdminAddr.Delete {
		event.AppendAttributes(sdk.NewAttribute(fmt.Sprintf("%s%d", AttributeKeyAdminAddrAdd, i), addrDelete.String()))
	}
	ctx.EventManager().EmitEvents(sdk.Events{event})

	return &types.MsgSetAdminAddrResponse{}, nil
}
