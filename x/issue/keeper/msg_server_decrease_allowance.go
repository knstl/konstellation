package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/konstellation/konstellation/x/issue/types"
)

func (k msgServer) DecreaseAllowance(goCtx context.Context, msg *types.MsgDecreaseAllowance) (*types.MsgDecreaseAllowanceResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgDecreaseAllowanceResponse{}, nil
}
