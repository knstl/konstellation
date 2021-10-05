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
func (m MsgFreeze) Route() string { return RouterKey }

// Type Implements Msg.
func (m MsgFreeze) Type() string { return TypeMsgFreeze }

// ValidateBasic Implements Msg.
func (m MsgFreeze) ValidateBasic() error {
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

// GetSignBytes Implements Msg.
func (m MsgFreeze) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&m))
}

// GetSigners Implements Msg.
func (m MsgFreeze) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(m.Freezer)}
}
