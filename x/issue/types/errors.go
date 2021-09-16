package types

// DONTCOVER

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	DefaultCodespace                = ModuleName
	CodeInvalidGenesis              = 102
	CodeUnknownIssue                = 1
	CodeIssuerMismatch              = 2
	CodeInvalidDenom                = 3
	CodeAmountLowerAllowance        = 4
	CodeIssueExists                 = 5
	CodeNotEnoughFee                = 6
	CodeInvalidFeature              = 7
	CodeCanNotMint                  = 8
	CodeCanNotBurnOwner             = 9
	CodeCanNotBurnHolder            = 10
	CodeCanNotBurnFrom              = 11
	CodeCanNotFreeze                = 12
	CodeAmountNotValid              = 13
	CodeInvalidCoinDecimals         = 14
	CodeInvalidTotalSupply          = 15
	CodeInvalidDescription          = 16
	CodeInvalidSymbol               = 17
	CodeInvalidFreezeOp             = 18
	CodeNotTransferOut              = 19
	CodeNotTransferIn               = 20
	CodeInvalidCoinDecimalsMultiple = 21
	CodeInvalidDescriptionMaxLength = 22
	CodeInvalidInput                = 400
	CodeInvalidIssueFee             = 401
	CodeInvalidMintFee              = 402
	CodeInvalidBurnFee              = 402
	CodeInvalidBurnFromFee          = 403
	CodeInvalidFreezeFee            = 404
	CodeInvalidUnFreezeFee          = 405
	CodeInvalidTransferOwnerFee     = 406
)

// x/issue module sentinel errors
var (
	ErrSample = sdkerrors.Register(ModuleName, 1100, "sample error")
	// this line is used by starport scaffolding # ibc/errors
	ErrInvalidIssueParams               = sdkerrors.Register(DefaultCodespace, CodeInvalidInput, "Invalid issue params")
	ErrIssueAlreadyExists               = sdkerrors.Register(DefaultCodespace, CodeIssueExists, "Invalid already exists")
	ErrNotEnoughFee                     = sdkerrors.Register(DefaultCodespace, CodeNotEnoughFee, "Not enough fee")
	ErrCoinDecimalsMaxValueNotValid     = sdkerrors.Register(DefaultCodespace, CodeInvalidCoinDecimals, fmt.Sprintf("Decimals max value is %d", CoinDecimalsMaxValue))
	ErrCoinDecimalsMultipleNotValid     = sdkerrors.Register(DefaultCodespace, CodeInvalidCoinDecimalsMultiple, fmt.Sprintf("Decimals must be a multiple of %d", CoinDecimalsMultiple))
	ErrCoinTotalSupplyMaxValueNotValid  = sdkerrors.Register(DefaultCodespace, CodeInvalidTotalSupply, fmt.Sprintf("Total supply max value is %s", CoinMaxTotalSupply.String()))
	ErrCoinDescriptionNotValid          = sdkerrors.Register(DefaultCodespace, CodeInvalidDescription, "Description is not valid json")
	ErrCoinDescriptionMaxLengthNotValid = sdkerrors.Register(DefaultCodespace, CodeInvalidDescriptionMaxLength, fmt.Sprintf("Description max length is %d", CoinDescriptionMaxLength))
	ErrCoinSymbolNotValid               = sdkerrors.Register(DefaultCodespace, CodeInvalidSymbol, "Invalid symbol")
)

//convert *sdkerrors.Error to error
func Errorf(err *sdkerrors.Error) error {
	return fmt.Errorf(err.Error())
}

func ErrUnknownIssue(denom string) *sdkerrors.Error {
	return sdkerrors.New(DefaultCodespace, CodeUnknownIssue, fmt.Sprintf("Unknown issue %s", denom))
}

func ErrOwnerMismatch(issueID string) *sdkerrors.Error {
	return sdkerrors.New(DefaultCodespace, CodeIssuerMismatch, fmt.Sprintf("Owner mismatch with token %s", issueID))
}

func ErrAmountGreaterThanAllowance(amt sdk.Coin, allowance sdk.Coin) *sdkerrors.Error {
	return sdkerrors.New(DefaultCodespace, CodeAmountLowerAllowance, fmt.Sprintf("Amount greater than allowance %s > %s", amt.String(), allowance.String()))
}

func ErrInvalidDenom(denom string) *sdkerrors.Error {
	return sdkerrors.New(DefaultCodespace, CodeInvalidDenom, fmt.Sprintf("Denom invalid %s", denom))
}

func ErrInvalidFreezeOp(op string) *sdkerrors.Error {
	return sdkerrors.New(DefaultCodespace, CodeInvalidFreezeOp, fmt.Sprintf("Invalid freeze type %s", op))
}

func ErrInvalidFeature(feature string) *sdkerrors.Error {
	return sdkerrors.New(DefaultCodespace, CodeInvalidFeature, fmt.Sprintf("Feature invalid %s", feature))
}

func ErrCanNotMint(denom string) *sdkerrors.Error {
	return sdkerrors.New(DefaultCodespace, CodeCanNotMint, fmt.Sprintf("Can not mint the token %s", denom))
}

func ErrCanNotBurnOwner(denom string) *sdkerrors.Error {
	return sdkerrors.New(DefaultCodespace, CodeCanNotBurnOwner, fmt.Sprintf("Can not burn the token %s", denom))
}

func ErrCanNotBurnHolder(denom string) *sdkerrors.Error {
	return sdkerrors.New(DefaultCodespace, CodeCanNotBurnHolder, fmt.Sprintf("Can not burn the token %s", denom))
}

func ErrCanNotBurnFrom(denom string) *sdkerrors.Error {
	return sdkerrors.New(DefaultCodespace, CodeCanNotBurnFrom, fmt.Sprintf("Can not burn the token %s", denom))
}

func ErrCanNotFreeze(denom string) *sdkerrors.Error {
	return sdkerrors.New(DefaultCodespace, CodeCanNotFreeze, fmt.Sprintf("Can not freeze the token %s", denom))
}

func ErrCanNotTransferIn(denom string, accAddress string) *sdkerrors.Error {
	return sdkerrors.New(DefaultCodespace, CodeNotTransferIn, fmt.Sprintf("Can not transfer in %s to %s", denom, accAddress))
}

func ErrCanNotTransferOut(denom string, accAddress string) *sdkerrors.Error {
	return sdkerrors.New(DefaultCodespace, CodeNotTransferOut, fmt.Sprintf("Can not transfer out %s from %s", denom, accAddress))
}

func ErrInvalidIssueFee(fee string) *sdkerrors.Error {
	return sdkerrors.New(DefaultCodespace, CodeInvalidIssueFee, fmt.Sprintf("invalid issue fee: %s", fee))
}

func ErrInvalidMintFee(fee string) *sdkerrors.Error {
	return sdkerrors.New(DefaultCodespace, CodeInvalidMintFee, fmt.Sprintf("invalid mint fee: %s", fee))
}

func ErrInvalidBurnFee(fee string) *sdkerrors.Error {
	return sdkerrors.New(DefaultCodespace, CodeInvalidBurnFee, fmt.Sprintf("invalid burn fee: %s", fee))
}

func ErrInvalidBurnFromFee(fee string) *sdkerrors.Error {
	return sdkerrors.New(DefaultCodespace, CodeInvalidBurnFromFee, fmt.Sprintf("invalid burn from fee: %s", fee))
}

func ErrInvalidFreezeFee(fee string) *sdkerrors.Error {
	return sdkerrors.New(DefaultCodespace, CodeInvalidFreezeFee, fmt.Sprintf("invalid freeze fee: %s", fee))
}

func ErrInvalidUnfreezeFee(fee string) *sdkerrors.Error {
	return sdkerrors.New(DefaultCodespace, CodeInvalidUnFreezeFee, fmt.Sprintf("invalid unfreeze fee: %s", fee))
}

func ErrInvalidTransferOwnerFee(fee string) *sdkerrors.Error {
	return sdkerrors.New(DefaultCodespace, CodeInvalidTransferOwnerFee, fmt.Sprintf("invalid transfer owner fee: %s", fee))
}
