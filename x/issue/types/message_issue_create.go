package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgIssueCreate{}

func NewMsgIssueCreate(owner, issuer sdk.AccAddress, params *IssueParams) *MsgIssueCreate {
	return &MsgIssueCreate{
		owner.String(),
		issuer.String(),
		params,
	}
}

func (msg *MsgIssueCreate) Route() string {
	return RouterKey
}

func (msg *MsgIssueCreate) Type() string {
	//return "IssueCreate"
	return TypeMsgIssueCreate
}

func (msg *MsgIssueCreate) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.Owner)}
}

func (msg *MsgIssueCreate) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgIssueCreate) ValidateBasic() error {
	if sdk.AccAddress(msg.Owner).Empty() {
		//return ErrNilOwner(DefaultCodespace)
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "Owner address cannot be empty")
	}
	// Cannot issue zero or negative coins
	if msg.TotalSupply.IsZero() || !msg.TotalSupply.IsPositive() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, "Cannot issue 0 or negative coin amounts")
	}
	//if utils.QuoDecimals(msg.TotalSupply, msg.Decimals).GT(CoinMaxTotalSupply) {
	//	return ErrCoinTotalSupplyMaxValueNotValid
	//}
	if len(msg.Symbol) < CoinSymbolMinLength || len(msg.Symbol) > CoinSymbolMaxLength {
		return ErrCoinSymbolNotValid
	}
	if uint(msg.Decimals) > CoinDecimalsMaxValue {
		return ErrCoinDecimalsMaxValueNotValid
	}
	if uint(msg.Decimals)%CoinDecimalsMultiple != 0 {
		return ErrCoinDecimalsMultipleNotValid
	}
	if len(msg.Description) > CoinDescriptionMaxLength {
		return ErrCoinDescriptionMaxLengthNotValid
	}
	return nil
}
