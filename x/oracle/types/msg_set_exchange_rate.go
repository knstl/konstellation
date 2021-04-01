package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// NewMsgSetExchangeRate is the constructor function for MsgSetExchangeRate
func NewMsgSetExchangeRate(denom string, exchangeRate *sdk.Coin, setter string) MsgSetExchangeRate {
	return MsgSetExchangeRate{
		ExchangeRate: exchangeRate,
		Setter:       setter,
	}
}

// Route should return the name of the module
func (msg *MsgSetExchangeRate) Route() string { return RouterKey }

// Type should return the action
func (msg *MsgSetExchangeRate) Type() string { return "set_exchange_rate" }

// ValidateBasic runs stateless checks on the message
func (msg *MsgSetExchangeRate) ValidateBasic() error {
	if msg.Setter == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Setter)
	}
	if !msg.ExchangeRate.IsPositive() {
		return sdkerrors.ErrInsufficientFunds
	}
	return nil
}

// GetSignBytes encodes the message for signing
func (msg *MsgSetExchangeRate) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)

}

// GetSigners defines whose signature is required
func (msg *MsgSetExchangeRate) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.Setter)}
}
