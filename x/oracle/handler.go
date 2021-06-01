package oracle

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/konstellation/konstellation/x/oracle/keeper"
	"github.com/konstellation/konstellation/x/oracle/types"
)

// RouterKey
const RouterKey = types.ModuleName

// NewHandler returns a handler for "oracle" type messages.
func NewHandler(k keeper.Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		switch msg := msg.(type) {
		case *types.MsgSetExchangeRate:
			return handleMsgSetExchangeRate(ctx, k, msg)
		case *types.MsgDeleteExchangeRate:
			return handleMsgDeleteExchangeRate(ctx, k, msg)
		case *types.MsgSetAdminAddr:
			return handleMsgSetAdminAddr(ctx, k, msg)
		default:
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, fmt.Sprintf("Unrecognized oracle Msg type: %v", msg.Type()))
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
