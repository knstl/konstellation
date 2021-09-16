package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/konstellation/konstellation/x/issue/types"
)

func (k msgServer) IncreaseAllowance(goCtx context.Context, msg *types.MsgIncreaseAllowance) (*types.MsgIncreaseAllowanceResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgIncreaseAllowanceResponse{}, nil
}
