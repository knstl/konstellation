package types

import sdk "github.com/cosmos/cosmos-sdk/types"

const (
	// module name
	ModuleName = "issue"

	// StoreKey is string representation of the store key for issue
	StoreKey = "issue"

	RouterKey = "issue"

	// QuerierRoute is the querier route for acc
	QuerierRoute = StoreKey

	CoinDecimalsMaxValue                  = uint(18)
	CoinDecimalsMultiple                  = uint(3)
	CodeInvalidGenesis       sdk.CodeType = 102
	CoinNameMinLength                     = 3
	CoinNameMaxLength                     = 32
	CoinSymbolMinLength                   = 2
	CoinSymbolMaxLength                   = 8
	CoinDescriptionMaxLength              = 1024
)
