package keeper

import (
	"fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Key for getting a the next available proposalID from the store
var (
	KeyFirstIssueDenom = []byte("firstIssueDenom")
	KeyLastIssueDenom  = []byte("lastIssueDenom")
	KeyLastIssueId     = []byte("lastIssueId")
	KeyDelimiter       = ":"
)

//func BytesString(b []byte) string {
//	return *(*string)(unsafe.Pointer(&b))
//}

// Key for getting a specific issuer from the store
func KeyIssuer(denom string) []byte {
	return []byte(fmt.Sprintf("issues:%s", denom))
}

// Key for getting a specific address from the store
func KeyAddressDenoms(addr string) []byte {
	return []byte(fmt.Sprintf("address:%s", addr))
}

// Key for getting a specific allowed from the store
func KeyAllowance(denom string, owner sdk.AccAddress, spender sdk.AccAddress) []byte {
	return []byte(fmt.Sprintf("allowed:%s:%s:%s", denom, owner.String(), spender.String()))
}

// Key for getting a specific allowed from the store
func KeyAllowances(denom string, owner sdk.AccAddress) []byte {
	return []byte(fmt.Sprintf("allowed:%s:%s", denom, owner.String()))
}

func KeyFreeze(denom string, accAddress sdk.AccAddress) []byte {
	return []byte(fmt.Sprintf("freeze:%s:%s", denom, accAddress.String()))
}

func PrefixFreeze(denom string) []byte {
	return []byte(fmt.Sprintf("freeze:%s", denom))
}

func KeySymbolDenom(symbol string) []byte {
	return []byte(fmt.Sprintf("symbol:%s", strings.ToUpper(symbol)))
}

func KeyIdDenom(id uint64) []byte {
	return []byte(fmt.Sprintf("id:%x", id))
}
