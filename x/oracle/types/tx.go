package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgSetExchangeRate{}
var _ sdk.Msg = &MsgDeleteExchangeRate{}

// NewMsgSetExchangeRate is the constructor function for MsgSetExchangeRate
func NewMsgSetExchangeRate(exchangeRate *sdk.Coin, sender string) MsgSetExchangeRate {
	return MsgSetExchangeRate{
		ExchangeRate: *exchangeRate,
		Sender:       sender,
	}
}

// Route should return the name of the module
func (msg MsgSetExchangeRate) Route() string { return RouterKey }

// Type should return the action
func (msg MsgSetExchangeRate) Type() string { return "set_exchange_rate" }

// GetSignBytes encodes the message for signing
func (msg MsgSetExchangeRate) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)

}

// GetSigners defines whose signature is required
func (msg MsgSetExchangeRate) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.Sender)}
}

// ValidateBasic runs stateless checks on the message
func (msg MsgSetExchangeRate) ValidateBasic() error {
	if msg.Sender == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Sender)
	}
	if !msg.ExchangeRate.IsPositive() {
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
		Add:    add,
		Delete: del,
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
