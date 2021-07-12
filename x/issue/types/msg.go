package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
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

func NewMsgIssueCreate(owner, issuer sdk.AccAddress, params *IssueParams) MsgIssueCreate {
	return MsgIssueCreate{
		owner.String(),
		issuer.String(),
		params,
	}
}

func (msg MsgIssueCreate) Route() string { return ModuleName }
func (msg MsgIssueCreate) Type() string  { return TypeMsgIssueCreate }

// Return address that must sign over msg.GetSignBytes()
func (msg MsgIssueCreate) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.Owner)}
}

// get the bytes for the message signer to sign on
func (msg MsgIssueCreate) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// quick validity check
func (msg MsgIssueCreate) ValidateBasic() error {
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

func NewMsgDescription(owner sdk.AccAddress, denom, description string) MsgDescription {
	return MsgDescription{Owner: owner.String(), Denom: denom, Description: description}
}

// Route Implements Msg.
func (msg MsgDescription) Route() string { return RouterKey }

// Type Implements Msg.
func (msg MsgDescription) Type() string { return TypeMsgDescription }

// ValidateBasic Implements Msg.
func (msg MsgDescription) ValidateBasic() error {
	if sdk.AccAddress(msg.Owner).Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "missing owner address")
	}
	if msg.Denom == "" {
		return ErrInvalidDenom(msg.Denom)
	}

	if len(msg.Description) > CoinDescriptionMaxLength {
		return ErrCoinDescriptionMaxLengthNotValid
	}
	return nil
}

// GetSignBytes Implements Msg.
func (msg MsgDescription) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners Implements Msg.
func (msg MsgDescription) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.Owner)}
}

func NewMsgDisableFeature(owner sdk.AccAddress, denom, feature string) MsgDisableFeature {
	return MsgDisableFeature{Owner: owner.String(), Denom: denom, Feature: feature}
}

// Route Implements Msg.
func (msg MsgDisableFeature) Route() string { return RouterKey }

// Type Implements Msg.
func (msg MsgDisableFeature) Type() string { return TypeMsgDisableFeature }

// ValidateBasic Implements Msg.
func (msg MsgDisableFeature) ValidateBasic() error {
	if sdk.AccAddress(msg.Owner).Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "missing owner address")
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
	return []sdk.AccAddress{sdk.AccAddress(msg.Owner)}
}

func NewMsgEnableFeature(owner sdk.AccAddress, denom, feature string) MsgEnableFeature {
	return MsgEnableFeature{Owner: owner.String(), Denom: denom, Feature: feature}
}

// Route Implements Msg.
func (msg MsgEnableFeature) Route() string { return RouterKey }

// Type Implements Msg.
func (msg MsgEnableFeature) Type() string { return TypeMsgEnableFeature }

// ValidateBasic Implements Msg.
func (msg MsgEnableFeature) ValidateBasic() error {
	if sdk.AccAddress(msg.Owner).Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "missing owner address")
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
	return []sdk.AccAddress{sdk.AccAddress(msg.Owner)}
}

func NewMsgFeatures(owner sdk.AccAddress, denom string, features *IssueFeatures) MsgFeatures {
	return MsgFeatures{
		owner.String(),
		denom,
		features,
	}
}

func (msg MsgFeatures) Route() string { return ModuleName }
func (msg MsgFeatures) Type() string  { return TypeMsgFeatures }

// Return address that must sign over msg.GetSignBytes()
func (msg MsgFeatures) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.Owner)}
}

// get the bytes for the message signer to sign on
func (msg MsgFeatures) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// quick validity check
func (msg MsgFeatures) ValidateBasic() error {
	if sdk.AccAddress(msg.Owner).Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "Owner address cannot be empty")
	}
	if msg.Denom == "" {
		return ErrInvalidDenom(msg.Denom)
	}
	return nil
}

// NewMsgTransfer - construct arbitrary multi-in, multi-out send msg.
func NewMsgTransfer(fromAddr, toAddr sdk.AccAddress, amount sdk.Coins) MsgTransfer {
	coins := Coins{Coins: []sdk.Coin{}}
	for _, coin := range amount {
		coins.Coins = append(coins.Coins, coin)
	}
	return MsgTransfer{FromAddress: fromAddr.String(), ToAddress: toAddr.String(), Amount: &coins}
}

// Route Implements Msg.
func (msg MsgTransfer) Route() string { return RouterKey }

// Type Implements Msg.
func (msg MsgTransfer) Type() string { return TypeMsgTransfer }

// ValidateBasic Implements Msg.
func (msg MsgTransfer) ValidateBasic() error {
	if sdk.AccAddress(msg.FromAddress).Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "missing sender address")
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

// GetSignBytes Implements Msg.
func (msg MsgTransfer) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners Implements Msg.
func (msg MsgTransfer) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.FromAddress)}
}

func NewMsgTransferFrom(sender, fromAddr, toAddr sdk.AccAddress, amount sdk.Coins) MsgTransferFrom {
	coins := Coins{Coins: []sdk.Coin{}}
	for _, coin := range amount {
		coins.Coins = append(coins.Coins, coin)
	}
	return MsgTransferFrom{Sender: sender.String(), FromAddress: fromAddr.String(), ToAddress: toAddr.String(), Amount: &coins}
}

// Route Implements Msg.
func (msg MsgTransferFrom) Route() string { return RouterKey }

// Type Implements Msg.
func (msg MsgTransferFrom) Type() string { return TypeMsgTransferFrom }

// ValidateBasic Implements Msg.
func (msg MsgTransferFrom) ValidateBasic() error {
	if sdk.AccAddress(msg.Sender).Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "missing sender address")
	}
	if sdk.AccAddress(msg.FromAddress).Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "missing from address")
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

// GetSignBytes Implements Msg.
func (msg MsgTransferFrom) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners Implements Msg.
func (msg MsgTransferFrom) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.Sender)}
}

func NewMsgTransferOwnership(owner, toAddr sdk.AccAddress, denom string) MsgTransferOwnership {
	return MsgTransferOwnership{Owner: owner.String(), ToAddress: toAddr.String(), Denom: denom}
}

// Route Implements Msg.
func (msg MsgTransferOwnership) Route() string { return RouterKey }

// Type Implements Msg.
func (msg MsgTransferOwnership) Type() string { return TypeMsgTransferOwnership }

// ValidateBasic Implements Msg.
func (msg MsgTransferOwnership) ValidateBasic() error {
	if sdk.AccAddress(msg.Owner).Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "missing owner address")
	}
	if sdk.AccAddress(msg.ToAddress).Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "missing recipient address")
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
	return []sdk.AccAddress{sdk.AccAddress(msg.Owner)}
}

// NewMsgApprove - construct arbitrary multi-in, multi-out send msg.
func NewMsgApprove(owner, spender sdk.AccAddress, amount sdk.Coins) MsgApprove {
	coins := Coins{Coins: []sdk.Coin{}}
	for _, coin := range amount {
		coins.Coins = append(coins.Coins, coin)
	}
	return MsgApprove{owner.String(), spender.String(), &coins}
}

// Route Implements Msg.
func (msg MsgApprove) Route() string { return RouterKey }

// Type Implements Msg.
func (msg MsgApprove) Type() string { return TypeMsgApprove }

// ValidateBasic Implements Msg.
func (msg MsgApprove) ValidateBasic() error {
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

// GetSignBytes Implements Msg.
func (msg MsgApprove) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners Implements Msg.
func (msg MsgApprove) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.Owner)}
}

func NewMsgIncreaseAllowance(owner, spender sdk.AccAddress, amount sdk.Coins) MsgIncreaseAllowance {
	coins := Coins{Coins: []sdk.Coin{}}
	for _, coin := range amount {
		coins.Coins = append(coins.Coins, coin)
	}
	return MsgIncreaseAllowance{owner.String(), spender.String(), &coins}
}

// Route Implements Msg.
func (msg MsgIncreaseAllowance) Route() string { return RouterKey }

// Type Implements Msg.
func (msg MsgIncreaseAllowance) Type() string { return TypeMsgIncreaseAllowance }

// ValidateBasic Implements Msg.
func (msg MsgIncreaseAllowance) ValidateBasic() error {
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

// GetSignBytes Implements Msg.
func (msg MsgIncreaseAllowance) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners Implements Msg.
func (msg MsgIncreaseAllowance) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.Owner)}
}

// NewMsgApprove - construct arbitrary multi-in, multi-out send msg.
func NewMsgDecreaseAllowance(owner, spender sdk.AccAddress, amount sdk.Coins) MsgDecreaseAllowance {
	coins := Coins{Coins: []sdk.Coin{}}
	for _, coin := range amount {
		coins.Coins = append(coins.Coins, coin)
	}
	return MsgDecreaseAllowance{owner.String(), spender.String(), &coins}
}

// Route Implements Msg.
func (msg MsgDecreaseAllowance) Route() string { return RouterKey }

// Type Implements Msg.
func (msg MsgDecreaseAllowance) Type() string { return TypeMsgDecreaseAllowance }

// ValidateBasic Implements Msg.
func (msg MsgDecreaseAllowance) ValidateBasic() error {
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

// GetSignBytes Implements Msg.
func (msg MsgDecreaseAllowance) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners Implements Msg.
func (msg MsgDecreaseAllowance) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.Owner)}
}

func NewMsgMint(minter, toAddr sdk.AccAddress, amount sdk.Coins) MsgMint {
	coins := Coins{Coins: []sdk.Coin{}}
	for _, coin := range amount {
		coins.Coins = append(coins.Coins, coin)
	}
	return MsgMint{Minter: minter.String(), ToAddress: toAddr.String(), Amount: &coins}
}

// Route Implements Msg.
func (msg MsgMint) Route() string { return RouterKey }

// Type Implements Msg.
func (msg MsgMint) Type() string { return TypeMsgMint }

// ValidateBasic Implements Msg.
func (msg MsgMint) ValidateBasic() error {
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

// GetSignBytes Implements Msg.
func (msg MsgMint) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners Implements Msg.
func (msg MsgMint) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.Minter)}
}

func NewMsgBurn(burner sdk.AccAddress, amount sdk.Coins) MsgBurn {
	coins := Coins{Coins: []sdk.Coin{}}
	for _, coin := range amount {
		coins.Coins = append(coins.Coins, coin)
	}
	return MsgBurn{Burner: burner.String(), Amount: &coins}
}

// Route Implements Msg.
func (msg MsgBurn) Route() string { return RouterKey }

// Type Implements Msg.
func (msg MsgBurn) Type() string { return TypeMsgBurn }

// ValidateBasic Implements Msg.
func (msg MsgBurn) ValidateBasic() error {
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

// GetSignBytes Implements Msg.
func (msg MsgBurn) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners Implements Msg.
func (msg MsgBurn) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.Burner)}
}

func NewMsgBurnFrom(burner, fromAddr sdk.AccAddress, amount sdk.Coins) MsgBurnFrom {
	coins := Coins{Coins: []sdk.Coin{}}
	for _, coin := range amount {
		coins.Coins = append(coins.Coins, coin)
	}
	return MsgBurnFrom{Burner: burner.String(), FromAddress: fromAddr.String(), Amount: &coins}
}

// Route Implements Msg.
func (msg MsgBurnFrom) Route() string { return RouterKey }

// Type Implements Msg.
func (msg MsgBurnFrom) Type() string { return TypeMsgBurnFrom }

// ValidateBasic Implements Msg.
func (msg MsgBurnFrom) ValidateBasic() error {
	if sdk.AccAddress(msg.Burner).Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "missing burner address")
	}
	if sdk.AccAddress(msg.FromAddress).Empty() {
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

// GetSignBytes Implements Msg.
func (msg MsgBurnFrom) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners Implements Msg.
func (msg MsgBurnFrom) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.Burner)}
}

func NewMsgFreeze(freezer, holder sdk.AccAddress, denom, op string) MsgFreeze {
	return MsgFreeze{Freezer: freezer.String(), Holder: holder.String(), Denom: denom, Op: op}
}

// Route Implements Msg.
func (msg MsgFreeze) Route() string { return RouterKey }

// Type Implements Msg.
func (msg MsgFreeze) Type() string { return TypeMsgFreeze }

// ValidateBasic Implements Msg.
func (msg MsgFreeze) ValidateBasic() error {
	if sdk.AccAddress(msg.Freezer).Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "missing freezer address")
	}
	if sdk.AccAddress(msg.Holder).Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "missing holder address")
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
	return []sdk.AccAddress{sdk.AccAddress(msg.Freezer)}
}

func NewMsgUnfreeze(freezer, holder sdk.AccAddress, denom, op string) MsgUnfreeze {
	return MsgUnfreeze{Freezer: freezer.String(), Holder: holder.String(), Denom: denom, Op: op}
}

// Route Implements Msg.
func (msg MsgUnfreeze) Route() string { return RouterKey }

// Type Implements Msg.
func (msg MsgUnfreeze) Type() string { return TypeMsgUnfreeze }

// ValidateBasic Implements Msg.
func (msg MsgUnfreeze) ValidateBasic() error {
	if sdk.AccAddress(msg.Freezer).Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "missing freezer address")
	}
	if sdk.AccAddress(msg.Holder).Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "missing holder address")
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
	return []sdk.AccAddress{sdk.AccAddress(msg.Freezer)}
}
