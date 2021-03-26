package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgDeleteExchangeRate{}

type MsgDeleteExchangeRate struct {
	Denom   string         `json:"denom" yaml:"denom"`
	Creator sdk.AccAddress `json:"creator" yaml:"creator"`
}

func NewMsgDeleteExchangeRate(denom string, creator sdk.AccAddress) MsgDeleteExchangeRate {
	return MsgDeleteExchangeRate{
		Denom:   denom,
		Creator: creator,
	}
}

func (msg MsgDeleteExchangeRate) Route() string {
	return RouterKey
}

func (msg MsgDeleteExchangeRate) Type() string {
	return "DeleteName"
}

func (msg MsgDeleteExchangeRate) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

func (msg MsgDeleteExchangeRate) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}

func (msg MsgDeleteExchangeRate) ValidateBasic() error {
	if msg.Creator.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
	}
	return nil
}

func (msg MsgDeleteExchangeRate) ProtoMessage() {
}

func (msg MsgDeleteExchangeRate) Reset() {
}

func (msg MsgDeleteExchangeRate) String() string {
	return msg.Type()
}
