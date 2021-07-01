package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	TypeMsgIssueCreate       = "issue_create"
	TypeMsgDescription       = "description"
	TypeMsgDisableFeature    = "disable_feature"
	TypeMsgEnableFeature     = "enable_feature"
	TypeMsgFeatures          = "features"
	TypeMsgTransfer          = "transfer"
	TypeMsgTransferFrom      = "transfer_from"
	TypeMsgApprove           = "approve"
	TypeMsgIncreaseAllowance = "increase_allowance"
	TypeMsgDecreaseAllowance = "decrease_allowance"
	TypeMsgMint              = "mint"
	TypeMsgBurn              = "burn"
	TypeMsgBurnFrom          = "burn_from"
	TypeMsgTransferOwnership = "transfer_ownership"
	TypeMsgFreeze            = "freeze"
	TypeMsgUnfreeze          = "unfreeze"
)

var _ sdk.Msg = &MsgIssueCreate{}
var _ sdk.Msg = &MsgDescription{}
var _ sdk.Msg = &MsgDisableFeature{}
var _ sdk.Msg = &MsgEnableFeature{}
var _ sdk.Msg = &MsgFeatures{}
var _ sdk.Msg = &MsgTransfer{}
var _ sdk.Msg = &MsgApprove{}
var _ sdk.Msg = &MsgIncreaseAllowance{}
var _ sdk.Msg = &MsgDecreaseAllowance{}
var _ sdk.Msg = &MsgTransferFrom{}
var _ sdk.Msg = &MsgMint{}
var _ sdk.Msg = &MsgBurn{}
var _ sdk.Msg = &MsgBurnFrom{}
var _ sdk.Msg = &MsgTransferOwnership{}
var _ sdk.Msg = &MsgFreeze{}
var _ sdk.Msg = &MsgUnfreeze{}

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
	if msg.TotalSupply.IsZero() || !msg.TotalSupply.IsPositive() {
		return sdk.ErrInvalidCoins("Cannot issue 0 or negative coin amounts")
	}
	//if utils.QuoDecimals(msg.TotalSupply, msg.Decimals).GT(CoinMaxTotalSupply) {
	//	return ErrCoinTotalSupplyMaxValueNotValid()
	//}
	if len(msg.Symbol) < CoinSymbolMinLength || len(msg.Symbol) > CoinSymbolMaxLength {
		return ErrCoinSymbolNotValid()
	}
	if msg.Decimals > CoinDecimalsMaxValue {
		return ErrCoinDecimalsMaxValueNotValid()
	}
	if msg.Decimals%CoinDecimalsMultiple != 0 {
		return ErrCoinDecimalsMultipleNotValid()
	}
	if len(msg.Description) > CoinDescriptionMaxLength {
		return ErrCoinDescriptionMaxLengthNotValid()
	}
	return nil
}

type MsgDescription struct {
	Owner       sdk.AccAddress `json:"owner" yaml:"owner"`
	Denom       string         `json:"denom" yaml:"denom"`
	Description string         `json:"description" yaml:"description"`
}

func NewMsgDescription(owner sdk.AccAddress, denom, description string) MsgDescription {
	return MsgDescription{Owner: owner, Denom: denom, Description: description}
}

// Route Implements Msg.
func (msg MsgDescription) Route() string { return RouterKey }

// Type Implements Msg.
func (msg MsgDescription) Type() string { return TypeMsgDescription }

// ValidateBasic Implements Msg.
func (msg MsgDescription) ValidateBasic() sdk.Error {
	if msg.Owner.Empty() {
		return sdk.ErrInvalidAddress("missing owner address")
	}
	if msg.Denom == "" {
		return ErrInvalidDenom(msg.Denom)
	}

	if len(msg.Description) > CoinDescriptionMaxLength {
		return ErrCoinDescriptionMaxLengthNotValid()
	}
	return nil
}

// GetSignBytes Implements Msg.
func (msg MsgDescription) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners Implements Msg.
func (msg MsgDescription) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Owner}
}

type MsgDisableFeature struct {
	Owner   sdk.AccAddress `json:"owner" yaml:"owner"`
	Denom   string         `json:"denom" yaml:"denom"`
	Feature string         `json:"feature" yaml:"feature"`
}

func NewMsgDisableFeature(owner sdk.AccAddress, denom, feature string) MsgDisableFeature {
	return MsgDisableFeature{Owner: owner, Denom: denom, Feature: feature}
}

// Route Implements Msg.
func (msg MsgDisableFeature) Route() string { return RouterKey }

// Type Implements Msg.
func (msg MsgDisableFeature) Type() string { return TypeMsgDisableFeature }

// ValidateBasic Implements Msg.
func (msg MsgDisableFeature) ValidateBasic() sdk.Error {
	if msg.Owner.Empty() {
		return sdk.ErrInvalidAddress("missing owner address")
	}

	if msg.Denom == "" {
		return ErrInvalidDenom(msg.Denom)
	}
	if msg.Feature == "" {
		return ErrInvalidFeature(msg.Feature)
	}
	return nil
}

// GetSignBytes Implements Msg.
func (msg MsgDisableFeature) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners Implements Msg.
func (msg MsgDisableFeature) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Owner}
}

type MsgEnableFeature struct {
	Owner   sdk.AccAddress `json:"owner" yaml:"owner"`
	Denom   string         `json:"denom" yaml:"denom"`
	Feature string         `json:"feature" yaml:"feature"`
}

func NewMsgEnableFeature(owner sdk.AccAddress, denom, feature string) MsgEnableFeature {
	return MsgEnableFeature{Owner: owner, Denom: denom, Feature: feature}
}

// Route Implements Msg.
func (msg MsgEnableFeature) Route() string { return RouterKey }

// Type Implements Msg.
func (msg MsgEnableFeature) Type() string { return TypeMsgEnableFeature }

// ValidateBasic Implements Msg.
func (msg MsgEnableFeature) ValidateBasic() sdk.Error {
	if msg.Owner.Empty() {
		return sdk.ErrInvalidAddress("missing owner address")
	}
	if msg.Denom == "" {
		return ErrInvalidDenom(msg.Denom)
	}
	if msg.Feature == "" {
		return ErrInvalidFeature(msg.Feature)
	}
	return nil
}

// GetSignBytes Implements Msg.
func (msg MsgEnableFeature) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners Implements Msg.
func (msg MsgEnableFeature) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Owner}
}

type MsgFeatures struct {
	Owner          sdk.AccAddress `json:"owner" yaml:"owner"`
	Denom          string         `json:"denom"`
	*IssueFeatures `json:"features"`
}

func NewMsgFeatures(owner sdk.AccAddress, denom string, features *IssueFeatures) MsgFeatures {
	return MsgFeatures{
		owner,
		denom,
		features,
	}
}

func (msg MsgFeatures) Route() string { return ModuleName }
func (msg MsgFeatures) Type() string  { return TypeMsgFeatures }

// Return address that must sign over msg.GetSignBytes()
func (msg MsgFeatures) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Owner}
}

// get the bytes for the message signer to sign on
func (msg MsgFeatures) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// quick validity check
func (msg MsgFeatures) ValidateBasic() sdk.Error {
	if msg.Owner.Empty() {
		return sdk.ErrInvalidAddress("Owner address cannot be empty")
	}
	if msg.Denom == "" {
		return ErrInvalidDenom(msg.Denom)
	}
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

type MsgTransferOwnership struct {
	Owner     sdk.AccAddress `json:"owner" yaml:"owner"`
	ToAddress sdk.AccAddress `json:"to_address" yaml:"to_address"`
	Denom     string         `json:"denom" yaml:"denom"`
}

func NewMsgTransferOwnership(owner, toAddr sdk.AccAddress, denom string) MsgTransferOwnership {
	return MsgTransferOwnership{Owner: owner, ToAddress: toAddr, Denom: denom}
}

// Route Implements Msg.
func (msg MsgTransferOwnership) Route() string { return RouterKey }

// Type Implements Msg.
func (msg MsgTransferOwnership) Type() string { return TypeMsgTransferOwnership }

// ValidateBasic Implements Msg.
func (msg MsgTransferOwnership) ValidateBasic() sdk.Error {
	if msg.Owner.Empty() {
		return sdk.ErrInvalidAddress("missing owner address")
	}
	if msg.ToAddress.Empty() {
		return sdk.ErrInvalidAddress("missing recipient address")
	}
	if msg.Denom == "" {
		return ErrInvalidDenom(msg.Denom)
	}
	return nil
}

// GetSignBytes Implements Msg.
func (msg MsgTransferOwnership) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners Implements Msg.
func (msg MsgTransferOwnership) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Owner}
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
	Minter    sdk.AccAddress `json:"minter" yaml:"minter"`
	ToAddress sdk.AccAddress `json:"to_address" yaml:"to_address"`
	Amount    sdk.Coins      `json:"amount" yaml:"amount"`
}

func NewMsgMint(minter, toAddr sdk.AccAddress, amount sdk.Coins) MsgMint {
	return MsgMint{Minter: minter, ToAddress: toAddr, Amount: amount}
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
func (msg MsgMint) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners Implements Msg.
func (msg MsgMint) GetSigners() []sdk.AccAddress {
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
		return sdk.ErrInvalidAddress("missing burner address")
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

type MsgFreeze struct {
	Freezer sdk.AccAddress `json:"freezer" yaml:"freezer"`
	Holder  sdk.AccAddress `json:"holder" yaml:"holder"`
	Denom   string         `json:"denom" yaml:"denom"`
	Op      string         `json:"op" yaml:"op"`
}

func NewMsgFreeze(freezer, holder sdk.AccAddress, denom, op string) MsgFreeze {
	return MsgFreeze{Freezer: freezer, Holder: holder, Denom: denom, Op: op}
}

// Route Implements Msg.
func (msg MsgFreeze) Route() string { return RouterKey }

// Type Implements Msg.
func (msg MsgFreeze) Type() string { return TypeMsgFreeze }

// ValidateBasic Implements Msg.
func (msg MsgFreeze) ValidateBasic() sdk.Error {
	if msg.Freezer.Empty() {
		return sdk.ErrInvalidAddress("missing freezer address")
	}
	if msg.Holder.Empty() {
		return sdk.ErrInvalidAddress("missing holder address")
	}
	if msg.Denom == "" {
		return ErrInvalidDenom(msg.Denom)
	}
	if msg.Op == "" {
		return ErrInvalidFreezeOp(msg.Op)
	}
	return nil
}

// GetSignBytes Implements Msg.
func (msg MsgFreeze) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners Implements Msg.
func (msg MsgFreeze) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Freezer}
}

type MsgUnfreeze struct {
	Freezer sdk.AccAddress `json:"freezer" yaml:"freezer"`
	Holder  sdk.AccAddress `json:"holder" yaml:"holder"`
	Denom   string         `json:"denom" yaml:"denom"`
	Op      string         `json:"op" yaml:"op"`
}

func NewMsgUnfreeze(freezer, holder sdk.AccAddress, denom, op string) MsgUnfreeze {
	return MsgUnfreeze{Freezer: freezer, Holder: holder, Denom: denom, Op: op}
}

// Route Implements Msg.
func (msg MsgUnfreeze) Route() string { return RouterKey }

// Type Implements Msg.
func (msg MsgUnfreeze) Type() string { return TypeMsgUnfreeze }

// ValidateBasic Implements Msg.
func (msg MsgUnfreeze) ValidateBasic() sdk.Error {
	if msg.Freezer.Empty() {
		return sdk.ErrInvalidAddress("missing freezer address")
	}
	if msg.Holder.Empty() {
		return sdk.ErrInvalidAddress("missing holder address")
	}
	if msg.Denom == "" {
		return ErrInvalidDenom(msg.Denom)
	}
	if msg.Op == "" {
		return ErrInvalidFreezeOp(msg.Op)
	}
	return nil
}

// GetSignBytes Implements Msg.
func (msg MsgUnfreeze) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners Implements Msg.
func (msg MsgUnfreeze) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Freezer}
}
