package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgSetExchangeRate{}
var _ sdk.Msg = &MsgDeleteExchangeRate{}

const (
	TypeMsgSetExchangeRate = "set_exchange_rate"
)

// NewMsgSetExchangeRate is the constructor function for MsgSetExchangeRate
func NewMsgSetExchangeRate(sender sdk.AccAddress, exchangeRate *ExchangeRate) *MsgSetExchangeRate {
	return &MsgSetExchangeRate{
		ExchangeRate: exchangeRate,
		Sender:       sender.String(),
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
	if m.ExchangeRate.Rate <= 0 {
		return sdkerrors.ErrInsufficientFunds
	}
	return nil
}

func NewMsgDeleteExchangeRate(sender string) MsgDeleteExchangeRate {
	return MsgDeleteExchangeRate{
		Sender: sender,
	}
}

func (msg MsgDeleteExchangeRate) Route() string {
	return RouterKey
}

func (msg MsgDeleteExchangeRate) Type() string {
	return "delete_exchange_rate"
}

func (msg MsgDeleteExchangeRate) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.Sender)}
}

func (msg MsgDeleteExchangeRate) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}

func (msg MsgDeleteExchangeRate) ValidateBasic() error {
	if msg.Sender == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "sender can't be empty")
	}
	return nil
}

func NewMsgSetAdminAddr(sender string, add []string, del []string) MsgSetAdminAddr {
	return MsgSetAdminAddr{
		Sender: sender,
		//.Add:    add,
		//Delete: del,
	}
}

func (msg MsgSetAdminAddr) Route() string {
	return RouterKey
}

func (msg MsgSetAdminAddr) Type() string {
	return "delete_exchange_rate"
}

func (msg MsgSetAdminAddr) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.Sender)}
}

func (msg MsgSetAdminAddr) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}

func (msg MsgSetAdminAddr) ValidateBasic() error {
	if msg.Sender == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "sender can't be empty")
	}
	return nil
}
