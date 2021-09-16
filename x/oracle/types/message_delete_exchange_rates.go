package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgDeleteExchangeRates{}

func NewMsgDeleteExchangeRates(sender string, pairs []string) *MsgDeleteExchangeRates {
	return &MsgDeleteExchangeRates{
		Sender: sender,
		Pairs:  pairs,
	}
}

func (m *MsgDeleteExchangeRates) Route() string {
	return RouterKey
}

func (m *MsgDeleteExchangeRates) Type() string {
	return TypeMsgDeleteExchangeRates
}

func (m *MsgDeleteExchangeRates) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(m.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (m *MsgDeleteExchangeRates) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(m)
	return sdk.MustSortJSON(bz)
}

func (m *MsgDeleteExchangeRates) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(m.Sender)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid sender address (%s)", err)
	}
	return nil
}
