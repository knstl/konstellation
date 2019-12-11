// nolint
package types

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type CodeType = sdk.CodeType

const (
	DefaultCodespace         sdk.CodespaceType = "issue"
	CodeInvalidInput         CodeType          = 103
	CodeIssuerMismatch       sdk.CodeType      = 2
	CodeIssueIDNotValid      sdk.CodeType      = 3
	CodeAmountLowerAllowance sdk.CodeType      = 4
	CodeUnknownIssue         sdk.CodeType      = 10
)

//convert sdk.Error to error
func Errorf(err sdk.Error) error {
	return fmt.Errorf(err.Stacktrace().Error())
}

//func ErrNil(codespace sdk.CodespaceType) sdk.Error {
//	return sdk.NewError(codespace, CodeInvalidInput, "is nil")
//}
//
//func ErrNilOwner(codespace sdk.CodespaceType) sdk.Error {
//	return sdk.NewError(codespace, CodeInvalidInput, "Owner is nil")
//}

func ErrInvalidIssueParams() sdk.Error {
	return sdk.NewError(DefaultCodespace, CodeInvalidInput, "Invalid issue params")
}

func ErrUnknownIssue(issueID string) sdk.Error {
	return sdk.NewError(DefaultCodespace, CodeUnknownIssue, fmt.Sprintf("Unknown issue with id %s", issueID))
}

func ErrOwnerMismatch(issueID string) sdk.Error {
	return sdk.NewError(DefaultCodespace, CodeIssuerMismatch, fmt.Sprintf("Owner mismatch with token %s", issueID))
}

func ErrAmountGreaterThanAllowance(amt sdk.Coin, allowance sdk.Coin) sdk.Error {
	return sdk.NewError(DefaultCodespace, CodeAmountLowerAllowance, fmt.Sprintf("Amount greater than allowance %x > %x", amt.String(), allowance.String()))
}

// Error constructors

//func ErrNotEnoughFee() sdk.Error {
//	return sdk.NewError(DefaultCodespace, CodeNotEnoughFee, fmt.Sprintf("Not enough fee"))
//}
//func ErrAmountNotValid(key string) sdk.Error {
//	return sdk.NewError(DefaultCodespace, CodeAmountNotValid, "%s is not a valid amount", key)
//}
//func ErrCoinDecimalsMaxValueNotValid() sdk.Error {
//	return sdk.NewError(DefaultCodespace, CodeIssueCoinDecimalsNotValid, fmt.Sprintf("Decimals max value is %d", types.CoinDecimalsMaxValue))
//}
//func ErrCoinDecimalsMultipleNotValid() sdk.Error {
//	return sdk.NewError(DefaultCodespace, CodeIssueCoinDecimalsNotValid, fmt.Sprintf("Decimals must be a multiple of %d", types.CoinDecimalsMultiple))
//}
//func ErrCoinTotalSupplyMaxValueNotValid() sdk.Error {
//	return sdk.NewError(DefaultCodespace, CodeIssueTotalSupplyNotValid, fmt.Sprintf("Total supply max value is %s", types.CoinMaxTotalSupply.String()))
//}
//func ErrCoinSymbolNotValid() sdk.Error {
//	return sdk.NewError(DefaultCodespace, CodeIssueSymbolNotValid, fmt.Sprintf("Symbol length is %d-%d character", types.CoinSymbolMinLength, types.CoinSymbolMaxLength))
//}
//func ErrFreezeEndTimestampNotValid() sdk.Error {
//	return sdk.NewError(DefaultCodespace, CodeFreezeEndTimeNotValid, "end-time is not a valid timestamp")
//}
//func ErrCoinNamelNotValid() sdk.Error {
//	return sdk.NewError(DefaultCodespace, CodeIssueNameNotValid, fmt.Sprintf("The length of the name is between %d and %d", types.CoinNameMinLength, types.CoinNameMaxLength))
//}
//func ErrCoinDescriptionNotValid() sdk.Error {
//	return sdk.NewError(DefaultCodespace, CodeIssueDescriptionNotValid, "Description is not valid json")
//}
//func ErrCoinDescriptionMaxLengthNotValid() sdk.Error {
//	return sdk.NewError(DefaultCodespace, CodeIssueDescriptionNotValid, "Description max length is %d", types.CoinDescriptionMaxLength)
//}
//func ErrCanNotMint(issueID string) sdk.Error {
//	return sdk.NewError(DefaultCodespace, CanNotMint, fmt.Sprintf("Can not mint the token %s", issueID))
//}
//func ErrCanNotBurn(issueID string, burnType string) sdk.Error {
//	return sdk.NewError(DefaultCodespace, CanNotBurn, fmt.Sprintf("Can not burn the token %s by %s", issueID, burnType))
//}
//func ErrUnknownFeatures() sdk.Error {
//	return sdk.NewError(DefaultCodespace, CodeUnknownFeature, fmt.Sprintf("Unknown feature"))
//}
//func ErrCanNotFreeze(issueID string) sdk.Error {
//	return sdk.NewError(DefaultCodespace, CodeCanNotFreeze, fmt.Sprintf("Can not freeze the token %s", issueID))
//}
//func ErrUnknownFreezeType() sdk.Error {
//	return sdk.NewError(DefaultCodespace, CodeUnknownFreezeType, fmt.Sprintf("Unknown type"))
//}
//func ErrNotEnoughAmountToTransfer() sdk.Error {
//	return sdk.NewError(DefaultCodespace, CodeNotEnoughAmountToTransfer, fmt.Sprintf("Not enough amount allowed to transfer"))
//}
//func ErrCanNotTransferIn(issueID string, accAddress string) sdk.Error {
//	return sdk.NewError(DefaultCodespace, CodeNotTransferIn, fmt.Sprintf("Can not transfer in %s to %s", issueID, accAddress))
//}
//func ErrCanNotTransferOut(issueID string, accAddress string) sdk.Error {
//	return sdk.NewError(DefaultCodespace, CodeNotTransferOut, fmt.Sprintf("Can not transfer out %s from %s", issueID, accAddress))
//}
