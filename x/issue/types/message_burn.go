package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgBurn{}

func NewMsgBurn(burner sdk.AccAddress, amount sdk.Coins) *MsgBurn {
	return &MsgBurn{Burner: burner.String(), Amount: amount}
}

func (m *MsgBurn) Route() string {
	return RouterKey
}

func (m *MsgBurn) Type() string {
	return TypeMsgBurn
}

func (m *MsgBurn) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(m.Burner)}
}

func (m *MsgBurn) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(m)
	return sdk.MustSortJSON(bz)
}

func (m *MsgBurn) ValidateBasic() error {
	if sdk.AccAddress(m.Burner).Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "missing burner address")
	}
	if !m.Amount.IsValid() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, "send amount is invalid: "+m.Amount.String())
	}
	if !m.Amount.IsAllPositive() {
		return sdkerrors.Wrap(sdkerrors.ErrInsufficientFunds, "send amount must be positive")
	}
	return nil
}
