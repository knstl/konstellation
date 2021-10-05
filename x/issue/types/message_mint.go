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

func (m *MsgMint) Route() string {
	return RouterKey
}

func (m *MsgMint) Type() string {
	return TypeMsgMint
}

func (m *MsgMint) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(m.Minter)}
}

func (m *MsgMint) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(m)
	return sdk.MustSortJSON(bz)
}

func (m *MsgMint) ValidateBasic() error {
	if sdk.AccAddress(m.Minter).Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "missing minter address")
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
