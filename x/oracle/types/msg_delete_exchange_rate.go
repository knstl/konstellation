package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgDeleteExchangeRate{}

func NewMsgDeleteExchangeRate(denom string, creator string) MsgDeleteExchangeRate {
	return MsgDeleteExchangeRate{
		Denom:   denom,
		Creator: creator,
	}
}

func (msg *MsgDeleteExchangeRate) Route() string {
	return RouterKey
}

func (msg *MsgDeleteExchangeRate) Type() string {
	return "delete_exchange_rate"
}

func (msg *MsgDeleteExchangeRate) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

func (msg *MsgDeleteExchangeRate) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteExchangeRate) ValidateBasic() error {
	if msg.Creator == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
	}
	return nil
}
