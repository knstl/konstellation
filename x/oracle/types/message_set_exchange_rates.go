package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgSetExchangeRates{}

func NewMsgSetExchangeRates(sender string, exchangeRates []*ExchangeRate) *MsgSetExchangeRates {
	return &MsgSetExchangeRates{
		Sender:        sender,
		ExchangeRates: exchangeRates,
	}
}

func (m *MsgSetExchangeRates) Route() string {
	return RouterKey
}

func (m *MsgSetExchangeRates) Type() string {
	return TypeMsgSetExchangeRates
}

func (m *MsgSetExchangeRates) GetSigners() []sdk.AccAddress {
	sender, err := sdk.AccAddressFromBech32(m.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{sender}
}

func (m *MsgSetExchangeRates) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(m)
	return sdk.MustSortJSON(bz)
}

func (m *MsgSetExchangeRates) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(m.Sender)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid sender address (%s)", err)
	}
	return nil
}
