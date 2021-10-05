package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgDisableFeature{}

func NewMsgDisableFeature(owner sdk.AccAddress, denom string, feature string) *MsgDisableFeature {
	return &MsgDisableFeature{
		Owner:   owner.String(),
		Denom:   denom,
		Feature: feature,
	}
}

func (m *MsgDisableFeature) Route() string {
	return RouterKey
}

func (m *MsgDisableFeature) Type() string {
	return TypeMsgDisableFeature
}

func (m *MsgDisableFeature) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(m.Owner)}
}

func (m *MsgDisableFeature) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(m)
	return sdk.MustSortJSON(bz)
}

func (m *MsgDisableFeature) ValidateBasic() error {
	if sdk.AccAddress(m.Owner).Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "missing owner address")
	}

	if m.Denom == "" {
		return ErrInvalidDenom(m.Denom)
	}
	if m.Feature == "" {
		return ErrInvalidFeature(m.Feature)
	}
	return nil
}
