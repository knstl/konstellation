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

var (
	// Keys for store prefixes
	// Last* values are constant during a block.
	LastValidatorPowerKey = []byte{0x11} // prefix for each key to a validator index, for bonded validators
	LastTotalPowerKey     = []byte{0x12} // prefix for the total power

	DenomsKey                 = []byte{0x21} // prefix for each key to a denom
	OwnershipsKey             = []byte{0x22} // prefix for each key to a address corresponding to denom
	ValidatorsByPowerIndexKey = []byte{0x23} // prefix for each key to a validator index, sorted by power

	DelegationKey                    = []byte{0x31} // key for a delegation
	UnbondingDelegationKey           = []byte{0x32} // key for an unbonding-delegation
	UnbondingDelegationByValIndexKey = []byte{0x33} // prefix for each key for an unbonding-delegation, by validator operator
	RedelegationKey                  = []byte{0x34} // key for a redelegation
	RedelegationByValSrcIndexKey     = []byte{0x35} // prefix for each key for an redelegation, by source validator operator
	RedelegationByValDstIndexKey     = []byte{0x36} // prefix for each key for an redelegation, by destination validator operator

	UnbondingQueueKey    = []byte{0x41} // prefix for the timestamps in unbonding queue
	RedelegationQueueKey = []byte{0x42} // prefix for the timestamps in redelegations queue
	ValidatorQueueKey    = []byte{0x43} // prefix for the timestamps in validator queue

	HistoricalInfoKey = []byte{0x50} // prefix for the historical info
)

//func BytesString(b []byte) string {
//	return *(*string)(unsafe.Pointer(&b))
//}

func GetDenomKey(denom string) []byte {
	return append(DenomsKey, []byte(denom)...)
}

func GetOwnershipsKey(addr string) []byte {
	return append(OwnershipsKey, []byte(addr)...)
}

// Key for getting a specific address from the store
func GetOwnershipKey(addr, denom string) []byte {
	return append(GetOwnershipsKey(addr), []byte(denom)...)
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
