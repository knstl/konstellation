package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgMint{}

func NewMsgMint(minter, toAddr sdk.AccAddress, amount sdk.Coins) *MsgMint {
	return &MsgMint{
		Minter:    minter.String(),
		ToAddress: toAddr.String(),
		Amount:    amount,
	}
}

func (msg *MsgMint) Route() string {
	return RouterKey
}

func (msg *MsgMint) Type() string {
	return TypeMsgMint
}

func (msg *MsgMint) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.Minter)}
}

func (msg *MsgMint) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgMint) ValidateBasic() error {
	if sdk.AccAddress(msg.Minter).Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "missing minter address")
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
