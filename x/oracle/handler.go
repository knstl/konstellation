package oracle

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/konstellation/konstellation/x/oracle/keeper"
	"github.com/konstellation/konstellation/x/oracle/types"
)

// RouterKey
const RouterKey = types.ModuleName

// NewHandler returns a handler for "oracle" type messages.
func NewHandler(k keeper.Keeper) sdk.Handler {
	msgServer := keeper.NewMsgServerImpl(k)

	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		ctx = ctx.WithEventManager(sdk.NewEventManager())

		switch msg := msg.(type) {
		case *types.MsgSetExchangeRate:
			res, err := msgServer.SetExchangeRate(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)
		case *types.MsgDeleteExchangeRate:
			res, err := msgServer.DeleteExchangeRate(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)
		case *types.MsgSetAdminAddr:
			res, err := msgServer.SetAdminAddr(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)
		default:
			return nil, sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "unrecognized oracle message type: %T", msg)
		}
	}
}

// Handle a message to set exchange rate
func handleMsgSetExchangeRate(ctx sdk.Context, k keeper.Keeper, msg *types.MsgSetExchangeRate) (*sdk.Result, error) {
	senderAddr, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return nil, err
	}

	if err := k.SetExchangeRate(ctx, senderAddr, msg.ExchangeRate); err != nil {
		return nil, err
	}

	return &sdk.Result{}, nil
}

// Handle a message to delete exchange rate
func handleMsgDeleteExchangeRate(ctx sdk.Context, k keeper.Keeper, msg *types.MsgDeleteExchangeRate) (*sdk.Result, error) {
	senderAddr, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return nil, err
	}

	if err := k.DeleteExchangeRate(ctx, senderAddr, msg.Pair); err != nil {
		return nil, err
	}

	return &sdk.Result{}, nil
}

// Handle a message to set admin address
func handleMsgSetAdminAddr(ctx sdk.Context, k keeper.Keeper, msg *types.MsgSetAdminAddr) (*sdk.Result, error) {
	senderAddr, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return nil, err
	}

	var add []types.AdminAddr
	for _, a := range msg.Add {
		_, err := sdk.AccAddressFromBech32(a.GetAddress())
		if err != nil {
			return nil, err
		}

		add = append(add, *a)
	}
	var del []types.AdminAddr
	for _, a := range msg.Delete {
		_, err := sdk.AccAddressFromBech32(a.GetAddress())
		if err != nil {
			return nil, err
		}

		del = append(del, *a)
	}

	if err := k.SetAdminAddr(ctx, senderAddr, add, del); err != nil {
		return nil, err
	}
	return &sdk.Result{}, nil
}
