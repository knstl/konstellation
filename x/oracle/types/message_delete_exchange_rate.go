package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgDeleteExchangeRate{}

func NewMsgDeleteExchangeRate(sender string, pair string) *MsgDeleteExchangeRate {
	return &MsgDeleteExchangeRate{
		Sender: sender,
		Pair:   pair,
	}
}

func (m *MsgDeleteExchangeRate) Route() string {
	return RouterKey
}

func (m *MsgDeleteExchangeRate) Type() string {
	return TypeMsgDeleteExchangeRate
}

func (m *MsgDeleteExchangeRate) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(m.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (m *MsgDeleteExchangeRate) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(m)
	return sdk.MustSortJSON(bz)
}

func (m *MsgDeleteExchangeRate) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(m.Sender)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid sender address (%s)", err)
	}
	return nil
}
