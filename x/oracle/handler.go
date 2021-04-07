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
		default:
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, fmt.Sprintf("Unrecognized oracle Msg type: %v", msg.Type()))
		}
	}
}

// Handle a message to set exchange rate
func handleMsgSetExchangeRate(ctx sdk.Context, k keeper.Keeper, msg *types.MsgSetExchangeRate) (*sdk.Result, error) {
	// Checks if the the bid price is greater than the price paid by the current owner
	if k.GetAllowedAddress(ctx) != msg.Setter {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Incorrect allowed address") // If not, throw an error
	}
	k.SetExchangeRate(ctx, *msg.ExchangeRate)
	return &sdk.Result{}, nil
}

// Handle a message to set exchange rate
func handleMsgDeleteExchangeRate(ctx sdk.Context, k keeper.Keeper, msg *types.MsgDeleteExchangeRate) (*sdk.Result, error) {
	// Checks if the the bid price is greater than the price paid by the current owner
	if k.GetAllowedAddress(ctx) != msg.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Incorrect allowed address") // If not, throw an error
	}
	k.DeleteExchangeRate(ctx)
	return &sdk.Result{}, nil
}
