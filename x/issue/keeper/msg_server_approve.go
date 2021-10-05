package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/konstellation/konstellation/x/issue/types"
)

func (k msgServer) Approve(goCtx context.Context, msg *types.MsgApprove) (*types.MsgApproveResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	spender, err := sdk.AccAddressFromBech32(msg.Spender)
	if err != nil {
		return nil, err
	}

	owner, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		return nil, err
	}

	if err := k.keeper.Approve(ctx, owner, spender, msg.Amount); err != nil {
		return nil, err
	}

	return &types.MsgApproveResponse{}, nil
}
