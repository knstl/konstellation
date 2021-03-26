package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// MsgSetExchangeRate defines the SetExchangeRate message
type MsgSetExchangeRate struct {
	ExchangeRate sdk.Coin       `json:"exchange_rate"`
	Setter       sdk.AccAddress `json:"setter"`
}

// NewMsgSetExchangeRate is the constructor function for MsgSetExchangeRate
func NewMsgSetExchangeRate(denom string, exchangeRate sdk.Coin, setter sdk.AccAddress) MsgSetExchangeRate {
	return MsgSetExchangeRate{
		ExchangeRate: exchangeRate,
		Setter:       setter,
	}
}

// Route should return the name of the module
func (msg MsgSetExchangeRate) Route() string { return RouterKey }

// Type should return the action
func (msg MsgSetExchangeRate) Type() string { return "set_exchange_rate" }

// ValidateBasic runs stateless checks on the message
func (msg MsgSetExchangeRate) ValidateBasic() error {
	if msg.Setter.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Setter.String())
	}
	if !msg.ExchangeRate.IsPositive() {
		return sdkerrors.ErrInsufficientFunds
	}
	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgSetExchangeRate) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)

}

// GetSigners defines whose signature is required
func (msg MsgSetExchangeRate) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Setter}
}

func (msg MsgSetExchangeRate) ProtoMessage() {
}

func (msg MsgSetExchangeRate) Reset() {
}

func (msg MsgSetExchangeRate) String() string {
	return msg.Type()
}
