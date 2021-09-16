package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/konstellation/konstellation/x/issue/types"
)

func (k msgServer) Description(goCtx context.Context, msg *types.MsgDescription) (*types.MsgDescriptionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgDescriptionResponse{}, nil
}
