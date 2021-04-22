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
	err := k.SetExchangeRate(ctx, msg.Sender, msg.ExchangeRate)
	if err != nil {
		return nil, err
	}
	return &sdk.Result{}, nil
}

// Handle a message to delete exchange rate
func handleMsgDeleteExchangeRate(ctx sdk.Context, k keeper.Keeper, msg *types.MsgDeleteExchangeRate) (*sdk.Result, error) {
	err := k.DeleteExchangeRate(ctx, msg.Sender)
	if err != nil {
		return nil, err
	}
	return &sdk.Result{}, nil
}

// Handle a message to set admin address
func handleMsgSetAdminAddr(ctx sdk.Context, k keeper.Keeper, msg *types.MsgSetAdminAddr) (*sdk.Result, error) {
	err := k.SetAdminAddr(ctx, msg.Sender, msg.Add, msg.Delete)
	if err != nil {
		return nil, err
	}
	return &sdk.Result{}, nil
}
