package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/konstellation/konstellation/x/issue/types"
)

func (k msgServer) DisableFeature(goCtx context.Context, msg *types.MsgDisableFeature) (*types.MsgDisableFeatureResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgDisableFeatureResponse{}, nil
}
