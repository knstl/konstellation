package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	TypeMsgIssue = "issue"
)

var _, _, _ sdk.Msg = &MsgIssue{}, &MsgTransfer{}, &MsgApprove{}

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

// MsgTransfer - high level transaction of the coin module
type MsgTransfer struct {
	FromAddress sdk.AccAddress `json:"from_address" yaml:"from_address"`
	ToAddress   sdk.AccAddress `json:"to_address" yaml:"to_address"`
	Amount      sdk.Coins      `json:"amount" yaml:"amount"`
}

// NewMsgTransfer - construct arbitrary multi-in, multi-out send msg.
func NewMsgTransfer(fromAddr, toAddr sdk.AccAddress, amount sdk.Coins) MsgTransfer {
	return MsgTransfer{FromAddress: fromAddr, ToAddress: toAddr, Amount: amount}
}

// Route Implements Msg.
func (msg MsgTransfer) Route() string { return RouterKey }

// Type Implements Msg.
func (msg MsgTransfer) Type() string { return "send" }

// ValidateBasic Implements Msg.
func (msg MsgTransfer) ValidateBasic() sdk.Error {
	if msg.FromAddress.Empty() {
		return sdk.ErrInvalidAddress("missing sender address")
	}
	if msg.ToAddress.Empty() {
		return sdk.ErrInvalidAddress("missing recipient address")
	}
	if !msg.Amount.IsValid() {
		return sdk.ErrInvalidCoins("send amount is invalid: " + msg.Amount.String())
	}
	if !msg.Amount.IsAllPositive() {
		return sdk.ErrInsufficientCoins("send amount must be positive")
	}
	return nil
}

// GetSignBytes Implements Msg.
func (msg MsgTransfer) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners Implements Msg.
func (msg MsgTransfer) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.FromAddress}
}

// MsgApprove - high level transaction of the coin module
type MsgApprove struct {
	Owner   sdk.AccAddress `json:"owner" yaml:"owner"`
	Spender sdk.AccAddress `json:"spender" yaml:"spender"`
	Amount  sdk.Coin       `json:"amount" yaml:"amount"`
}

// NewMsgApprove - construct arbitrary multi-in, multi-out send msg.
func NewMsgApprove(owner, spender sdk.AccAddress, amount sdk.Coin) MsgApprove {
	return MsgApprove{owner, spender, amount}
}

// Route Implements Msg.
func (msg MsgApprove) Route() string { return RouterKey }

// Type Implements Msg.
func (msg MsgApprove) Type() string { return "send" }

// ValidateBasic Implements Msg.
func (msg MsgApprove) ValidateBasic() sdk.Error {
	if msg.Owner.Empty() {
		return sdk.ErrInvalidAddress("missing owner address")
	}
	if msg.Spender.Empty() {
		return sdk.ErrInvalidAddress("missing spender address")
	}
	if !msg.Amount.IsValid() {
		return sdk.ErrInvalidCoins("send amount is invalid: " + msg.Amount.String())
	}
	if !msg.Amount.IsPositive() {
		return sdk.ErrInsufficientCoins("send amount must be positive")
	}
	return nil
}

// GetSignBytes Implements Msg.
func (msg MsgApprove) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners Implements Msg.
func (msg MsgApprove) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Owner}
}
