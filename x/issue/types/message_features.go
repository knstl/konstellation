package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgFeatures{}

func NewMsgFeatures(owner sdk.AccAddress, denom string, features *IssueFeatures) *MsgFeatures {
	return &MsgFeatures{
		owner.String(),
		denom,
		features,
	}
}

func (msg *MsgFeatures) Route() string {
	return RouterKey
}

func (msg *MsgFeatures) Type() string {
	return "Features"
}

func (msg *MsgFeatures) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.Owner)}
}

func (msg *MsgFeatures) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgFeatures) ValidateBasic() error {
	if sdk.AccAddress(msg.Owner).Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "Owner address cannot be empty")
	}
	if msg.Denom == "" {
		return ErrInvalidDenom(msg.Denom)
	}
	return nil
}
