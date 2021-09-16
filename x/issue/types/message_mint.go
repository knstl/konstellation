package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgMint{}

func NewMsgMint(minter, toAddr sdk.AccAddress, amount sdk.Coins) *MsgMint {
	coins := Coins{Coins: []sdk.Coin{}}
	for _, coin := range amount {
		coins.Coins = append(coins.Coins, coin)
	}
	return &MsgMint{Minter: minter.String(), ToAddress: toAddr.String(), Amount: &coins}
}

func (msg *MsgMint) Route() string {
	return RouterKey
}

func (msg *MsgMint) Type() string {
	//return "Mint"
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
	if !sdk.Coins(msg.Amount.Coins).IsValid() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, "send amount is invalid: "+sdk.Coins(msg.Amount.Coins).String())
	}
	if !sdk.Coins(msg.Amount.Coins).IsAllPositive() {
		return sdkerrors.Wrap(sdkerrors.ErrInsufficientFunds, "send amount must be positive")
	}
	return nil
}
