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
