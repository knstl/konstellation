package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgSetAdminAddr{}

func NewMsgSetAdminAddr(sender string, add []*AdminAddr, del []*AdminAddr) *MsgSetAdminAddr {
	return &MsgSetAdminAddr{
		Sender: sender,
		Add:    add,
		Delete: del,
	}
}

func (m *MsgSetAdminAddr) Route() string {
	return RouterKey
}

func (m *MsgSetAdminAddr) Type() string {
	return TypeMsgSetAdminAddr
}

func (m *MsgSetAdminAddr) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(m.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (m *MsgSetAdminAddr) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(m)
	return sdk.MustSortJSON(bz)
}

func (m *MsgSetAdminAddr) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(m.Sender)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid sender address (%s)", err)
	}
	return nil
}
