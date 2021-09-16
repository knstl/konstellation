package keeper

import (
	"context"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/konstellation/konstellation/x/oracle/types"
)

func (m msgServer) SetAdminAddr(goCtx context.Context, msgSetAdminAddr *types.MsgSetAdminAddr) (*types.MsgSetAdminAddrResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	senderAddr, err := sdk.AccAddressFromBech32(msgSetAdminAddr.Sender)
	if err != nil {
		return nil, err
	}

	var add []types.AdminAddr
	for _, a := range msgSetAdminAddr.Add {
		_, err := sdk.AccAddressFromBech32(a.GetAddress())
		if err != nil {
			return nil, err
		}

		add = append(add, *a)
	}
	var del []types.AdminAddr
	for _, a := range msgSetAdminAddr.Delete {
		_, err := sdk.AccAddressFromBech32(a.GetAddress())
		if err != nil {
			return nil, err
		}

		del = append(del, *a)
	}

	if err := m.keeper.SetAdminAddr(ctx, senderAddr, add, del); err != nil {
		return nil, err
	}

	addAddresses := []string{}
	for _, add := range msgSetAdminAddr.Add {
		addAddresses = append(addAddresses, add.Address)
	}
	deleteAddresses := []string{}
	for _, del := range msgSetAdminAddr.Delete {
		addAddresses = append(addAddresses, del.Address)
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeSetAdminAddr,
			sdk.NewAttribute(types.AttributeKeyAdd, strings.Join(addAddresses, ",")),
			sdk.NewAttribute(types.AttributeKeyDelete, strings.Join(deleteAddresses, ",")),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msgSetAdminAddr.Sender),
		),
	})

	return &types.MsgSetAdminAddrResponse{}, nil
}
