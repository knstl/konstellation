package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgTransfer{}

func NewMsgTransfer(fromAddr, toAddr sdk.AccAddress, amount sdk.Coins) *MsgTransfer {
	return &MsgTransfer{FromAddress: fromAddr.String(), ToAddress: toAddr.String(), Amount: amount}
}

func (m *MsgTransfer) Route() string {
	return RouterKey
}

func (m *MsgTransfer) Type() string {
	return TypeMsgTransfer
}

func (m *MsgTransfer) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(m.FromAddress)}
}

func (m *MsgTransfer) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(m)
	return sdk.MustSortJSON(bz)
}

func (m *MsgTransfer) ValidateBasic() error {
	if sdk.AccAddress(m.FromAddress).Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "missing sender address")
	}
	if sdk.AccAddress(m.ToAddress).Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "missing recipient address")
	}
	if !m.Amount.IsValid() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, "send amount is invalid: "+m.Amount.String())
	}
	if !m.Amount.IsAllPositive() {
		return sdkerrors.Wrap(sdkerrors.ErrInsufficientFunds, "send amount must be positive")
	}
	return nil
}
