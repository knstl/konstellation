package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgTransferFrom{}

func NewMsgTransferFrom(sender, fromAddr, toAddr sdk.AccAddress, amount sdk.Coins) *MsgTransferFrom {
	return &MsgTransferFrom{Sender: sender.String(), FromAddress: fromAddr.String(), ToAddress: toAddr.String(), Amount: amount}
}

func (msg *MsgTransferFrom) Route() string {
	return RouterKey
}

func (msg *MsgTransferFrom) Type() string {
	return TypeMsgTransferFrom
}

func (msg *MsgTransferFrom) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.Sender)}
}

func (msg *MsgTransferFrom) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgTransferFrom) ValidateBasic() error {
	if sdk.AccAddress(msg.Sender).Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "missing sender address")
	}
	if sdk.AccAddress(msg.FromAddress).Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "missing from address")
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
