package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/konstellation/konstellation/x/issue/types"
)

func (k msgServer) IncreaseAllowance(goCtx context.Context, msg *types.MsgIncreaseAllowance) (*types.MsgIncreaseAllowanceResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	spender, err := sdk.AccAddressFromBech32(msg.Spender)
	if err != nil {
		return nil, err
	}

	owner, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		return nil, err
	}

	if err := k.keeper.IncreaseAllowance(ctx, owner, spender, msg.Amount); err != nil {
		return nil, err
	}

	return &types.MsgIncreaseAllowanceResponse{}, nil
}
