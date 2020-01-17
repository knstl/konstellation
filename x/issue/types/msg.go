package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	TypeMsgIssueCreate       = "issue_create"
	TypeMsgTransfer          = "transfer"
	TypeMsgTransferFrom      = "transfer_from"
	TypeMsgApprove           = "approve"
	TypeMsgIncreaseAllowance = "increase_allowance"
	TypeMsgDecreaseAllowance = "decrease_allowance"
	TypeMsgMint              = "mint"
	TypeMsgMintTo            = "mint_to"
	TypeMsgBurn              = "burn"
	TypeMsgBurnFrom          = "burn_from"
)

var _ sdk.Msg = &MsgIssueCreate{}
var _ sdk.Msg = &MsgTransfer{}
var _ sdk.Msg = &MsgApprove{}
var _ sdk.Msg = &MsgIncreaseAllowance{}
var _ sdk.Msg = &MsgDecreaseAllowance{}
var _ sdk.Msg = &MsgTransferFrom{}
var _ sdk.Msg = &MsgMint{}
var _ sdk.Msg = &MsgMintTo{}
var _ sdk.Msg = &MsgBurn{}
var _ sdk.Msg = &MsgBurnFrom{}

type MsgIssueCreate struct {
	Owner        sdk.AccAddress `json:"owner" yaml:"owner"`
	Issuer       sdk.AccAddress `json:"issuer" yaml:"issuer"`
	*IssueParams `json:"params"`
}

func NewMsgIssueCreate(owner, issuer sdk.AccAddress, params *IssueParams) MsgIssueCreate {
	return MsgIssueCreate{
		owner,
		issuer,
		params,
	}
}

func (msg MsgIssueCreate) Route() string { return ModuleName }
func (msg MsgIssueCreate) Type() string  { return TypeMsgIssueCreate }

// Return address that must sign over msg.GetSignBytes()
func (msg MsgIssueCreate) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Owner}
}

// get the bytes for the message signer to sign on
func (msg MsgIssueCreate) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// quick validity check
func (msg MsgIssueCreate) ValidateBasic() sdk.Error {
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
func (msg MsgTransfer) Type() string { return TypeMsgTransfer }

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

type MsgTransferFrom struct {
	Sender      sdk.AccAddress `json:"sender" yaml:"sender"`
	FromAddress sdk.AccAddress `json:"from_address" yaml:"from_address"`
	ToAddress   sdk.AccAddress `json:"to_address" yaml:"to_address"`
	Amount      sdk.Coins      `json:"amount" yaml:"amount"`
}

func NewMsgTransferFrom(sender, fromAddr, toAddr sdk.AccAddress, amount sdk.Coins) MsgTransferFrom {
	return MsgTransferFrom{Sender: sender, FromAddress: fromAddr, ToAddress: toAddr, Amount: amount}
}

// Route Implements Msg.
func (msg MsgTransferFrom) Route() string { return RouterKey }

// Type Implements Msg.
func (msg MsgTransferFrom) Type() string { return TypeMsgTransferFrom }

// ValidateBasic Implements Msg.
func (msg MsgTransferFrom) ValidateBasic() sdk.Error {
	if msg.Sender.Empty() {
		return sdk.ErrInvalidAddress("missing sender address")
	}
	if msg.FromAddress.Empty() {
		return sdk.ErrInvalidAddress("missing from address")
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
func (msg MsgTransferFrom) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners Implements Msg.
func (msg MsgTransferFrom) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Sender}
}

// MsgApprove - high level transaction of the coin module
type MsgApprove struct {
	Owner   sdk.AccAddress `json:"owner" yaml:"owner"`
	Spender sdk.AccAddress `json:"spender" yaml:"spender"`
	Amount  sdk.Coins      `json:"amount" yaml:"amount"`
}

// NewMsgApprove - construct arbitrary multi-in, multi-out send msg.
func NewMsgApprove(owner, spender sdk.AccAddress, amount sdk.Coins) MsgApprove {
	return MsgApprove{owner, spender, amount}
}

// Route Implements Msg.
func (msg MsgApprove) Route() string { return RouterKey }

// Type Implements Msg.
func (msg MsgApprove) Type() string { return TypeMsgApprove }

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
	if !msg.Amount.IsAllPositive() {
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

type MsgIncreaseAllowance struct {
	Owner   sdk.AccAddress `json:"owner" yaml:"owner"`
	Spender sdk.AccAddress `json:"spender" yaml:"spender"`
	Amount  sdk.Coins      `json:"amount" yaml:"amount"`
}

func NewMsgIncreaseAllowance(owner, spender sdk.AccAddress, amount sdk.Coins) MsgIncreaseAllowance {
	return MsgIncreaseAllowance{owner, spender, amount}
}

// Route Implements Msg.
func (msg MsgIncreaseAllowance) Route() string { return RouterKey }

// Type Implements Msg.
func (msg MsgIncreaseAllowance) Type() string { return TypeMsgIncreaseAllowance }

// ValidateBasic Implements Msg.
func (msg MsgIncreaseAllowance) ValidateBasic() sdk.Error {
	if msg.Owner.Empty() {
		return sdk.ErrInvalidAddress("missing owner address")
	}
	if msg.Spender.Empty() {
		return sdk.ErrInvalidAddress("missing spender address")
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
func (msg MsgIncreaseAllowance) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners Implements Msg.
func (msg MsgIncreaseAllowance) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Owner}
}

// MsgApprove - high level transaction of the coin module
type MsgDecreaseAllowance struct {
	Owner   sdk.AccAddress `json:"owner" yaml:"owner"`
	Spender sdk.AccAddress `json:"spender" yaml:"spender"`
	Amount  sdk.Coins      `json:"amount" yaml:"amount"`
}

// NewMsgApprove - construct arbitrary multi-in, multi-out send msg.
func NewMsgDecreaseAllowance(owner, spender sdk.AccAddress, amount sdk.Coins) MsgDecreaseAllowance {
	return MsgDecreaseAllowance{owner, spender, amount}
}

// Route Implements Msg.
func (msg MsgDecreaseAllowance) Route() string { return RouterKey }

// Type Implements Msg.
func (msg MsgDecreaseAllowance) Type() string { return TypeMsgDecreaseAllowance }

// ValidateBasic Implements Msg.
func (msg MsgDecreaseAllowance) ValidateBasic() sdk.Error {
	if msg.Owner.Empty() {
		return sdk.ErrInvalidAddress("missing owner address")
	}
	if msg.Spender.Empty() {
		return sdk.ErrInvalidAddress("missing spender address")
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
func (msg MsgDecreaseAllowance) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners Implements Msg.
func (msg MsgDecreaseAllowance) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Owner}
}

type MsgMint struct {
	Minter sdk.AccAddress `json:"minter" yaml:"minter"`
	Amount sdk.Coins      `json:"amount" yaml:"amount"`
}

func NewMsgMint(minter sdk.AccAddress, amount sdk.Coins) MsgMint {
	return MsgMint{Minter: minter, Amount: amount}
}

// Route Implements Msg.
func (msg MsgMint) Route() string { return RouterKey }

// Type Implements Msg.
func (msg MsgMint) Type() string { return TypeMsgMint }

// ValidateBasic Implements Msg.
func (msg MsgMint) ValidateBasic() sdk.Error {
	if msg.Minter.Empty() {
		return sdk.ErrInvalidAddress("missing minter address")
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
func (msg MsgMint) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners Implements Msg.
func (msg MsgMint) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Minter}
}

type MsgMintTo struct {
	Minter    sdk.AccAddress `json:"minter" yaml:"minter"`
	ToAddress sdk.AccAddress `json:"to_address" yaml:"to_address"`
	Amount    sdk.Coins      `json:"amount" yaml:"amount"`
}

func NewMsgMintTo(minter, toAddr sdk.AccAddress, amount sdk.Coins) MsgMintTo {
	return MsgMintTo{Minter: minter, ToAddress: toAddr, Amount: amount}
}

// Route Implements Msg.
func (msg MsgMintTo) Route() string { return RouterKey }

// Type Implements Msg.
func (msg MsgMintTo) Type() string { return TypeMsgMintTo }

// ValidateBasic Implements Msg.
func (msg MsgMintTo) ValidateBasic() sdk.Error {
	if msg.Minter.Empty() {
		return sdk.ErrInvalidAddress("missing minter address")
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
func (msg MsgMintTo) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners Implements Msg.
func (msg MsgMintTo) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Minter}
}

type MsgBurn struct {
	Burner sdk.AccAddress `json:"burner" yaml:"burner"`
	Amount sdk.Coins      `json:"amount" yaml:"amount"`
}

func NewMsgBurn(burner sdk.AccAddress, amount sdk.Coins) MsgBurn {
	return MsgBurn{Burner: burner, Amount: amount}
}

// Route Implements Msg.
func (msg MsgBurn) Route() string { return RouterKey }

// Type Implements Msg.
func (msg MsgBurn) Type() string { return TypeMsgBurn }

// ValidateBasic Implements Msg.
func (msg MsgBurn) ValidateBasic() sdk.Error {
	if msg.Burner.Empty() {
		return sdk.ErrInvalidAddress("missing minter address")
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
func (msg MsgBurn) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners Implements Msg.
func (msg MsgBurn) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Burner}
}

type MsgBurnFrom struct {
	Burner      sdk.AccAddress `json:"burner" yaml:"burner"`
	FromAddress sdk.AccAddress `json:"from_address" yaml:"from_address"`
	Amount      sdk.Coins      `json:"amount" yaml:"amount"`
}

func NewMsgBurnFrom(burner, fromAddr sdk.AccAddress, amount sdk.Coins) MsgBurnFrom {
	return MsgBurnFrom{Burner: burner, FromAddress: fromAddr, Amount: amount}
}

// Route Implements Msg.
func (msg MsgBurnFrom) Route() string { return RouterKey }

// Type Implements Msg.
func (msg MsgBurnFrom) Type() string { return TypeMsgBurnFrom }

// ValidateBasic Implements Msg.
func (msg MsgBurnFrom) ValidateBasic() sdk.Error {
	if msg.Burner.Empty() {
		return sdk.ErrInvalidAddress("missing burner address")
	}
	if msg.FromAddress.Empty() {
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
func (msg MsgBurnFrom) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners Implements Msg.
func (msg MsgBurnFrom) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Burner}
}
