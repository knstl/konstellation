package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgIncreaseAllowance{}

func NewMsgIncreaseAllowance(owner, spender sdk.AccAddress, amount sdk.Coins) *MsgIncreaseAllowance {
	return &MsgIncreaseAllowance{owner.String(), spender.String(), amount}
}

func (msg *MsgIncreaseAllowance) Route() string {
	return RouterKey
}

func (msg *MsgIncreaseAllowance) Type() string {
	return TypeMsgIncreaseAllowance
}

func (msg *MsgIncreaseAllowance) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.Owner)}
}

func (msg *MsgIncreaseAllowance) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgIncreaseAllowance) ValidateBasic() error {
	if sdk.AccAddress(msg.Owner).Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "missing owner address")
	}
	if sdk.AccAddress(msg.Spender).Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "missing spender address")
	}
	if !msg.Amount.IsValid() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, "send amount is invalid: "+msg.Amount.String())
	}
	if !msg.Amount.IsAllPositive() {
		return sdkerrors.Wrap(sdkerrors.ErrInsufficientFunds, "send amount must be positive")
	}
	return nil
}
