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

func (msg *MsgUnfreeze) Route() string {
	return RouterKey
}

func (msg *MsgUnfreeze) Type() string {
	return TypeMsgUnfreeze
}

func (msg *MsgUnfreeze) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.Freezer)}
}

func (msg *MsgUnfreeze) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUnfreeze) ValidateBasic() error {
	if sdk.AccAddress(msg.Freezer).Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "missing freezer address")
	}
	if sdk.AccAddress(msg.Holder).Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "missing holder address")
	}
	if msg.Denom == "" {
		return ErrInvalidDenom(msg.Denom)
	}
	if msg.Op == "" {
		return ErrInvalidFreezeOp(msg.Op)
	}
	return nil
}
