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

func (m *MsgFeatures) Route() string {
	return RouterKey
}

func (m *MsgFeatures) Type() string {
	return "Features"
}

func (m *MsgFeatures) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(m.Owner)}
}

func (m *MsgFeatures) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(m)
	return sdk.MustSortJSON(bz)
}

func (m *MsgFeatures) ValidateBasic() error {
	if sdk.AccAddress(m.Owner).Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "Owner address cannot be empty")
	}
	if m.Denom == "" {
		return ErrInvalidDenom(m.Denom)
	}
	return nil
}
