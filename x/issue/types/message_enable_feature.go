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

func (msg *MsgEnableFeature) Route() string {
	return RouterKey
}

func (msg *MsgEnableFeature) Type() string {
	//return "EnableFeature"
	return TypeMsgEnableFeature
}

func (msg *MsgEnableFeature) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.Owner)}
}

func (msg *MsgEnableFeature) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgEnableFeature) ValidateBasic() error {
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
