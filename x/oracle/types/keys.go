package types

import sdk "github.com/cosmos/cosmos-sdk/types"

const (
	// ModuleName defines the module name
	ModuleName = "oracle"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for slashing
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_oracle"

	// this line is used by starport scaffolding # ibc/keys/name
	QueryExchangeRate     = "exchange_rate"
	QueryAllExchangeRates = "all_exchange_rates"
)

// this line is used by starport scaffolding # ibc/keys/port

func KeyPrefix(p string) []byte {
	return []byte(p)
}

const (
	ExchangeRateKey      = "ExchangeRate-value-"
	ExchangeRateCountKey = "ExchangeRate-count-"
)

const (
	AdminAddrKey      = "AdminAddr-value-"
	AdminAddrCountKey = "AdminAddr-count-"
)

const (
	ParamsKey      = "Params-value-"
	ParamsCountKey = "Params-count-"
)

var (
	AllowedAddressKey = []byte{0x21} // prefix for each key to a allowed address

	ExchangeRateKeyValue = []byte{0x31} // prefix for  each key to exchange rate
)

func GetExchangeRateKey(pair string) []byte {
	return append(ExchangeRateKeyValue, []byte(pair)...)
}

func GetAllowedAddressKey(addr sdk.AccAddress) []byte {
	return append(AllowedAddressKey, addr.Bytes()...)
}
