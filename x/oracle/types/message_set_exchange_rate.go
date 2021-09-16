package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgSetExchangeRate{}

func NewMsgSetExchangeRate(sender string, exchangeRate *ExchangeRate) *MsgSetExchangeRate {
	return &MsgSetExchangeRate{
		Sender:       sender,
		ExchangeRate: exchangeRate,
	}
}

func (m *MsgSetExchangeRate) Route() string {
	return RouterKey
}

func (m *MsgSetExchangeRate) Type() string {
	return TypeMsgSetExchangeRate
}

func (m *MsgSetExchangeRate) GetSigners() []sdk.AccAddress {
	sender, err := sdk.AccAddressFromBech32(m.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{sender}
}

func (m *MsgSetExchangeRate) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(m)
	return sdk.MustSortJSON(bz)
}

func (m *MsgSetExchangeRate) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(m.Sender)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid sender address (%s)", err)
	}
	return nil
}
