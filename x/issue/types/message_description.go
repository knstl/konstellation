package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgDescription{}

func NewMsgDescription(owner sdk.AccAddress, denom string, description string) *MsgDescription {
	return &MsgDescription{
		Owner:       owner.String(),
		Denom:       denom,
		Description: description,
	}
}

func (m *MsgDescription) Route() string {
	return RouterKey
}

func (m *MsgDescription) Type() string {
	return TypeMsgDescription
}

func (m *MsgDescription) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(m.Owner)}
}

func (m *MsgDescription) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(m)
	return sdk.MustSortJSON(bz)
}

func (m *MsgDescription) ValidateBasic() error {
	if sdk.AccAddress(m.Owner).Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "missing owner address")
	}
	if m.Denom == "" {
		return ErrInvalidDenom(m.Denom)
	}

	if len(m.Description) > CoinDescriptionMaxLength {
		return ErrCoinDescriptionMaxLengthNotValid
	}
	return nil
}
