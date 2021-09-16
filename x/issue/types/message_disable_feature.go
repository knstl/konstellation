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

func (msg *MsgDisableFeature) Route() string {
	return RouterKey
}

func (msg *MsgDisableFeature) Type() string {
	//return "DisableFeature"
	return TypeMsgDisableFeature
}

func (msg *MsgDisableFeature) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.Owner)}
}

func (msg *MsgDisableFeature) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDisableFeature) ValidateBasic() error {
	if sdk.AccAddress(msg.Owner).Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "missing owner address")
	}

	if msg.Denom == "" {
		return ErrInvalidDenom(msg.Denom)
	}
	if msg.Feature == "" {
		return ErrInvalidFeature(msg.Feature)
	}
	return nil
}
