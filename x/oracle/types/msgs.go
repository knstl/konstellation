package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgSetExchangeRate     = "set_exchange_rate"
	TypeMsgSetExchangeRates    = "set_exchange_rates"
	TypeMsgDeleteExchangeRate  = "delete_exchange_rate"
	TypeMsgDeleteExchangeRates = "delete_exchange_rates"
	TypeMsgSetAdminAddr        = "set_admin_addr"
)

var _ sdk.Msg = &MsgSetExchangeRate{}
var _ sdk.Msg = &MsgDeleteExchangeRate{}
var _ sdk.Msg = &MsgDeleteExchangeRates{}

// NewMsgSetExchangeRate is the constructor function for MsgSetExchangeRate
func NewMsgSetExchangeRate(senderAddr sdk.AccAddress, exchangeRate *ExchangeRate) *MsgSetExchangeRate {
	return &MsgSetExchangeRate{
		ExchangeRate: exchangeRate,
		Sender:       senderAddr.String(),
	}
}

// Route should return the name of the module
func (m MsgSetExchangeRate) Route() string { return RouterKey }

// Type should return the action
func (m MsgSetExchangeRate) Type() string { return TypeMsgSetExchangeRate }

// GetSignBytes encodes the message for signing
func (m MsgSetExchangeRate) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&m))
}

// GetSigners defines whose signature is required
func (m MsgSetExchangeRate) GetSigners() []sdk.AccAddress {
	sender, err := sdk.AccAddressFromBech32(m.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{sender}
}

// ValidateBasic runs stateless checks on the message
func (m MsgSetExchangeRate) ValidateBasic() error {
	if m.Sender == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, m.Sender)
	}
	if m.ExchangeRate.Rate.LT(sdk.ZeroDec()) {
		return sdkerrors.Wrap(ErrInvalidRate, m.ExchangeRate.Pair)
	}
	return nil
}

func NewMsgDeleteExchangeRate(senderAddr sdk.AccAddress, pair string) *MsgDeleteExchangeRate {
	return &MsgDeleteExchangeRate{
		Sender: senderAddr.String(),
		Pair:   pair,
	}
}

func (m MsgDeleteExchangeRate) Route() string {
	return RouterKey
}

func (m MsgDeleteExchangeRate) Type() string {
	return TypeMsgDeleteExchangeRate
}

func (m MsgDeleteExchangeRate) GetSigners() []sdk.AccAddress {
	sender, err := sdk.AccAddressFromBech32(m.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{sender}
}

func (m MsgDeleteExchangeRate) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&m)
	return sdk.MustSortJSON(bz)
}

func (m MsgDeleteExchangeRate) ValidateBasic() error {
	if m.Sender == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "sender can't be empty")
	}
	return nil
}

func NewMsgSetAdminAddr(senderAddr sdk.AccAddress, add []*AdminAddr, del []*AdminAddr) *MsgSetAdminAddr {
	return &MsgSetAdminAddr{
		Sender: senderAddr.String(),
		Add:    add,
		Delete: del,
	}
}

func (m MsgSetAdminAddr) Route() string {
	return RouterKey
}

func (m MsgSetAdminAddr) Type() string {
	return TypeMsgSetAdminAddr
}

func (m MsgSetAdminAddr) GetSigners() []sdk.AccAddress {
	sender, err := sdk.AccAddressFromBech32(m.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{sender}
}

func (m MsgSetAdminAddr) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&m)
	return sdk.MustSortJSON(bz)
}

func (m MsgSetAdminAddr) ValidateBasic() error {
	if m.Sender == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "sender can't be empty")
	}
	return nil
}

// NewMsgSetExchangeRate is the constructor function for MsgSetExchangeRate
func NewMsgSetExchangeRates(senderAddr sdk.AccAddress, exchangeRates []*ExchangeRate) *MsgSetExchangeRates {
	return &MsgSetExchangeRates{
		ExchangeRates: exchangeRates,
		Sender:        senderAddr.String(),
	}
}

// Route should return the name of the module
func (m MsgSetExchangeRates) Route() string {
	return RouterKey
}

// Type should return the action
func (m MsgSetExchangeRates) Type() string {
	return TypeMsgSetExchangeRates
}

// GetSignBytes encodes the message for signing
func (m MsgSetExchangeRates) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&m))
}

// GetSigners defines whose signature is required
func (m MsgSetExchangeRates) GetSigners() []sdk.AccAddress {
	sender, err := sdk.AccAddressFromBech32(m.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{sender}
}

// ValidateBasic runs stateless checks on the message
func (m MsgSetExchangeRates) ValidateBasic() error {
	if m.Sender == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, m.Sender)
	}
	return nil
}

func NewMsgDeleteExchangeRates(senderAddr sdk.AccAddress, pairs []string) *MsgDeleteExchangeRates {
	return &MsgDeleteExchangeRates{
		Sender: senderAddr.String(),
		Pairs:  pairs,
	}
}

func (m MsgDeleteExchangeRates) Route() string {
	return RouterKey
}

func (m MsgDeleteExchangeRates) Type() string {
	return TypeMsgDeleteExchangeRate
}

func (m MsgDeleteExchangeRates) GetSigners() []sdk.AccAddress {
	sender, err := sdk.AccAddressFromBech32(m.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{sender}
}

func (m MsgDeleteExchangeRates) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&m)
	return sdk.MustSortJSON(bz)
}

func (m MsgDeleteExchangeRates) ValidateBasic() error {
	if m.Sender == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "sender can't be empty")
	}
	return nil
}
