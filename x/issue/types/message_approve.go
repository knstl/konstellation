package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgApprove{}

func NewMsgApprove(owner, spender sdk.AccAddress, amount sdk.Coins) *MsgApprove {
	coins := Coins{Coins: []sdk.Coin{}}
	for _, coin := range amount {
		coins.Coins = append(coins.Coins, coin)
	}
	return &MsgApprove{owner.String(), spender.String(), &coins}
}

func (msg *MsgApprove) Route() string {
	return RouterKey
}

func (msg *MsgApprove) Type() string {
	return TypeMsgApprove
}

func (msg *MsgApprove) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.Owner)}
}

func (msg *MsgApprove) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgApprove) ValidateBasic() error {
	if sdk.AccAddress(msg.Owner).Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "missing owner address")
	}
	if sdk.AccAddress(msg.Spender).Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "missing spender address")
	}
	if !sdk.Coins(msg.Amount.Coins).IsValid() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, "send amount is invalid: "+sdk.Coins(msg.Amount.Coins).String())
	}
	if !sdk.Coins(msg.Amount.Coins).IsAllPositive() {
		return sdkerrors.Wrap(sdkerrors.ErrInsufficientFunds, "send amount must be positive")
	}
	return nil
}
