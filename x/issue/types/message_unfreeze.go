package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgUnfreeze{}

func NewMsgUnfreeze(freezer sdk.AccAddress, holder sdk.AccAddress, denom string, op string) *MsgUnfreeze {
	return &MsgUnfreeze{
		Freezer: freezer.String(),
		Holder:  holder.String(),
		Denom:   denom,
		Op:      op,
	}
}

func (m *MsgUnfreeze) Route() string {
	return RouterKey
}

func (m *MsgUnfreeze) Type() string {
	return TypeMsgUnfreeze
}

func (m *MsgUnfreeze) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(m.Freezer)}
}

func (m *MsgUnfreeze) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(m)
	return sdk.MustSortJSON(bz)
}

func (m *MsgUnfreeze) ValidateBasic() error {
	if sdk.AccAddress(m.Freezer).Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "missing freezer address")
	}
	if sdk.AccAddress(m.Holder).Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "missing holder address")
	}
	if m.Denom == "" {
		return ErrInvalidDenom(m.Denom)
	}
	if m.Op == "" {
		return ErrInvalidFreezeOp(m.Op)
	}
	return nil
}
