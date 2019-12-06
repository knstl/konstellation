package types

import sdk "github.com/cosmos/cosmos-sdk/types"

const (
	TypeMsgIssue = "issue"
)

var _ sdk.Msg = &MsgIssue{}

type MsgIssue struct {
	Owner        sdk.AccAddress `json:"owner" yaml:"owner"`
	Issuer       sdk.AccAddress `json:"issuer" yaml:"issuer"`
	*IssueParams `json:"params"`
}

func NewMsgIssue(owner, issuer sdk.AccAddress, params *IssueParams) MsgIssue {
	return MsgIssue{
		owner,
		issuer,
		params,
	}
}

func (msg MsgIssue) Route() string { return ModuleName }
func (msg MsgIssue) Type() string  { return TypeMsgIssue }

// Return address that must sign over msg.GetSignBytes()
func (msg MsgIssue) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Owner}
}

// get the bytes for the message signer to sign on
func (msg MsgIssue) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// quick validity check
func (msg MsgIssue) ValidateBasic() sdk.Error {
	if msg.Owner.Empty() {
		//return ErrNilOwner(DefaultCodespace)
		return sdk.ErrInvalidAddress("Owner address cannot be empty")
	}
	// Cannot issue zero or negative coins
	//if msg.TotalSupply.IsZero() || !msg.TotalSupply.IsPositive() {
	//	return sdk.ErrInvalidCoins("Cannot issue 0 or negative coin amounts")
	//}
	//if utils.QuoDecimals(msg.TotalSupply, msg.Decimals).GT(CoinMaxTotalSupply) {
	//	return errors.ErrCoinTotalSupplyMaxValueNotValid()
	//}
	//if len(msg.Name) < CoinNameMinLength || len(msg.Name) > CoinNameMaxLength {
	//	return errors.ErrCoinNamelNotValid()
	//}
	//if len(msg.Symbol) < CoinSymbolMinLength || len(msg.Symbol) > CoinSymbolMaxLength {
	//	return errors.ErrCoinSymbolNotValid()
	//}
	//if msg.Decimals > CoinDecimalsMaxValue {
	//	return errors.ErrCoinDecimalsMaxValueNotValid()
	//}
	//if msg.Decimals%CoinDecimalsMultiple != 0 {
	//	return errors.ErrCoinDecimalsMultipleNotValid()
	//}
	//if len(msg.Description) > CoinDescriptionMaxLength {
	//	return errors.ErrCoinDescriptionMaxLengthNotValid()
	//}
	return nil
}
