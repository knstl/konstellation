package keeper

import (
	"context"

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
	return &types.MsgSetExchangeRateResponse{}, nil
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
	return &types.MsgDeleteExchangeRateResponse{}, nil
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

	//
	//ctx.EventManager().EmitEvents(sdk.Events{
	//	sdk.NewEvent(
	//		types.EventTypeCreateValidator,
	//		sdk.NewAttribute(types.AttributeKeyValidator, msg.ValidatorAddress),
	//		sdk.NewAttribute(sdk.AttributeKeyAmount, msg.Value.Amount.String()),
	//	),
	//	sdk.NewEvent(
	//		sdk.EventTypeMessage,
	//		sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
	//		sdk.NewAttribute(sdk.AttributeKeySender, msg.DelegatorAddress),
	//	),
	//})

	return &types.MsgSetAdminAddrResponse{}, nil
}
