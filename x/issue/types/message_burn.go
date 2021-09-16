package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgBurn{}

func NewMsgBurn(burner sdk.AccAddress, amount sdk.Coins) *MsgBurn {
	coins := Coins{Coins: []sdk.Coin{}}
	for _, coin := range amount {
		coins.Coins = append(coins.Coins, coin)
	}
	return &MsgBurn{Burner: burner.String(), Amount: &coins}
}

func (msg *MsgBurn) Route() string {
	return RouterKey
}

func (msg *MsgBurn) Type() string {
	//return "Burn"
	return TypeMsgBurn
}

func (msg *MsgBurn) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.Burner)}
}

func (msg *MsgBurn) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgBurn) ValidateBasic() error {
	if sdk.AccAddress(msg.Burner).Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "missing burner address")
	}
	if !sdk.Coins(msg.Amount.Coins).IsValid() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, "send amount is invalid: "+sdk.Coins(msg.Amount.Coins).String())
	}
	if !sdk.Coins(msg.Amount.Coins).IsAllPositive() {
		return sdkerrors.Wrap(sdkerrors.ErrInsufficientFunds, "send amount must be positive")
	}
	return nil
}
