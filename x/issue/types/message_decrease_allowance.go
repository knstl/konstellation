package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgDecreaseAllowance{}

func NewMsgDecreaseAllowance(owner, spender sdk.AccAddress, amount sdk.Coins) *MsgDecreaseAllowance {
	return &MsgDecreaseAllowance{owner.String(), spender.String(), amount}
}

func (m *MsgDecreaseAllowance) Route() string {
	return RouterKey
}

func (m *MsgDecreaseAllowance) Type() string {
	return TypeMsgDecreaseAllowance
}

func (m *MsgDecreaseAllowance) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(m.Owner)}
}

func (m *MsgDecreaseAllowance) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(m)
	return sdk.MustSortJSON(bz)
}

func (m *MsgDecreaseAllowance) ValidateBasic() error {
	if sdk.AccAddress(m.Owner).Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "missing owner address")
	}
	if sdk.AccAddress(m.Spender).Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "missing spender address")
	}
	if !m.Amount.IsValid() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, "send amount is invalid: "+m.Amount.String())
	}
	if !m.Amount.IsAllPositive() {
		return sdkerrors.Wrap(sdkerrors.ErrInsufficientFunds, "send amount must be positive")
	}
	return nil
}
