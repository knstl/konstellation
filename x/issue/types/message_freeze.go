package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgFreeze{}

func NewMsgFreeze(freezer, holder sdk.AccAddress, denom, op string) MsgFreeze {
	return MsgFreeze{Freezer: freezer.String(), Holder: holder.String(), Denom: denom, Op: op}
}

// Route Implements Msg.
func (msg MsgFreeze) Route() string { return RouterKey }

// Type Implements Msg.
func (msg MsgFreeze) Type() string { return TypeMsgFreeze }

// ValidateBasic Implements Msg.
func (msg MsgFreeze) ValidateBasic() error {
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

// GetSignBytes Implements Msg.
func (msg MsgFreeze) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

// GetSigners Implements Msg.
func (msg MsgFreeze) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.Freezer)}
}
