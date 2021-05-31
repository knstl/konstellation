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
	err := m.keeper.SetExchangeRate(ctx, msgSetExchangeRate.Sender, msgSetExchangeRate.ExchangeRate)
	if err != nil {
		return nil, err
	}
	return &types.MsgSetExchangeRateResponse{}, nil
}

func (m msgServer) DeleteExchangeRate(goCtx context.Context, msgDeleteExchangeRate *types.MsgDeleteExchangeRate) (*types.MsgDeleteExchangeRateResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	err := m.keeper.DeleteExchangeRate(ctx, msgDeleteExchangeRate.Sender, msgDeleteExchangeRate.Pair)
	if err != nil {
		return nil, err
	}
	return &types.MsgDeleteExchangeRateResponse{}, nil
}

func (m msgServer) SetAdminAddr(goCtx context.Context, msgSetAdminAddr *types.MsgSetAdminAddr) (*types.MsgSetAdminAddrResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	err := m.keeper.SetAdminAddr(ctx, msgSetAdminAddr.Sender, msgSetAdminAddr.Add, msgSetAdminAddr.Delete)
	if err != nil {
		return nil, err
	}
	return &types.MsgSetAdminAddrResponse{}, nil
}
