package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgEnableFeature{}

func NewMsgEnableFeature(owner string, denom string, feature string) *MsgEnableFeature {
	return &MsgEnableFeature{
		Owner:   owner,
		Denom:   denom,
		Feature: feature,
	}
}

func (m *MsgEnableFeature) Route() string {
	return RouterKey
}

func (m *MsgEnableFeature) Type() string {
	return TypeMsgEnableFeature
}

func (m *MsgEnableFeature) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(m.Owner)}
}

func (m *MsgEnableFeature) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(m)
	return sdk.MustSortJSON(bz)
}

func (m *MsgEnableFeature) ValidateBasic() error {
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
