package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/konstellation/konstellation/x/issue/types"
)

func (k msgServer) IssueCreate(goCtx context.Context, msg *types.MsgIssueCreate) (*types.MsgIssueCreateResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgIssueCreateResponse{}, nil
}
