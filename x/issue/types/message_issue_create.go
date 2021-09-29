package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgIssue{}

func NewMsgIssue(owner, issuer sdk.AccAddress, params *IssueParams) *MsgIssue {
	return &MsgIssue{
		owner.String(),
		issuer.String(),
		params,
	}
}

func (m *MsgIssue) Route() string {
	return RouterKey
}

func (m *MsgIssue) Type() string {
	return TypeMsgIssue
}

func (m *MsgIssue) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(m.Owner)}
}

func (m *MsgIssue) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(m)
	return sdk.MustSortJSON(bz)
}

func (m *MsgIssue) ValidateBasic() error {
	if sdk.AccAddress(m.Owner).Empty() {
		//return ErrNilOwner(DefaultCodespace)
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "Owner address cannot be empty")
	}
	// Cannot issue zero or negative coins
	if m.TotalSupply.IsZero() || !m.TotalSupply.IsPositive() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, "Cannot issue 0 or negative coin amounts")
	}
	//if utils.QuoDecimals(m.TotalSupply, m.Decimals).GT(CoinMaxTotalSupply) {
	//	return ErrCoinTotalSupplyMaxValueNotValid
	//}
	if len(m.Symbol) < CoinSymbolMinLength || len(m.Symbol) > CoinSymbolMaxLength {
		return ErrCoinSymbolNotValid
	}
	if uint(m.Decimals) > CoinDecimalsMaxValue {
		return ErrCoinDecimalsMaxValueNotValid
	}
	if uint(m.Decimals)%CoinDecimalsMultiple != 0 {
		return ErrCoinDecimalsMultipleNotValid
	}
	if len(m.Description) > CoinDescriptionMaxLength {
		return ErrCoinDescriptionMaxLengthNotValid
	}
	return nil
}
