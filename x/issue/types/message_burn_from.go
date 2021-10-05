package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgBurnFrom{}

func NewMsgBurnFrom(burner, fromAddr sdk.AccAddress, amount sdk.Coins) *MsgBurnFrom {
	return &MsgBurnFrom{Burner: burner.String(), FromAddress: fromAddr.String(), Amount: amount}
}

func (m *MsgBurnFrom) Route() string {
	return RouterKey
}

func (m *MsgBurnFrom) Type() string {
	return TypeMsgBurnFrom
}

func (m *MsgBurnFrom) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(m.Burner)}
}

func (m *MsgBurnFrom) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(m)
	return sdk.MustSortJSON(bz)
}

func (m *MsgBurnFrom) ValidateBasic() error {
	if sdk.AccAddress(m.Burner).Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "missing burner address")
	}
	if sdk.AccAddress(m.FromAddress).Empty() {
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
