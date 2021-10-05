package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgBurnFrom{}

func NewMsgBurnFrom(burner, fromAddr sdk.AccAddress, amount sdk.Coins) *MsgBurnFrom {
	return &MsgBurnFrom{Burner: burner.String(), FromAddress: fromAddr.String(), Amount: amount}
}

func (msg *MsgBurnFrom) Route() string {
	return RouterKey
}

func (msg *MsgBurnFrom) Type() string {
	return TypeMsgBurnFrom
}

func (msg *MsgBurnFrom) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.Burner)}
}

func (msg *MsgBurnFrom) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgBurnFrom) ValidateBasic() error {
	if sdk.AccAddress(msg.Burner).Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "missing burner address")
	}
	if sdk.AccAddress(msg.FromAddress).Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "missing recipient address")
	}
	if !msg.Amount.IsValid() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, "send amount is invalid: "+msg.Amount.String())
	}
	if !msg.Amount.IsAllPositive() {
		return sdkerrors.Wrap(sdkerrors.ErrInsufficientFunds, "send amount must be positive")
	}
	return nil
}
