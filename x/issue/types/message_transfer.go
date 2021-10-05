package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgTransfer{}

func NewMsgTransfer(fromAddr, toAddr sdk.AccAddress, amount sdk.Coins) *MsgTransfer {
	return &MsgTransfer{FromAddress: fromAddr.String(), ToAddress: toAddr.String(), Amount: amount}
}

func (msg *MsgTransfer) Route() string {
	return RouterKey
}

func (msg *MsgTransfer) Type() string {
	return TypeMsgTransfer
}

func (msg *MsgTransfer) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.FromAddress)}
}

func (msg *MsgTransfer) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgTransfer) ValidateBasic() error {
	if sdk.AccAddress(msg.FromAddress).Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "missing sender address")
	}
	if sdk.AccAddress(msg.ToAddress).Empty() {
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
