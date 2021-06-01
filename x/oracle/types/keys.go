package types

import sdk "github.com/cosmos/cosmos-sdk/types"

const (
	// ModuleName is the name of the module
	ModuleName = "oracle"

	RouterKey = ModuleName

	// StoreKey is the default store key for mint
	StoreKey     = ModuleName
	QuerierRoute = StoreKey

	QueryExchangeRate = "exchange_rate"
)

var (
	AllowedAddressKey = []byte{0x21} // prefix for each key to a allowed address

	ExchangeRateKey = []byte{0x31} // prefix for  each key to exchange rate
)

func GetExchangeRateKey(pair string) []byte {
	return append(ExchangeRateKey, []byte(pair)...)
}

func GetAllowedAddressKey(addr sdk.AccAddress) []byte {
	return append(AllowedAddressKey, addr.Bytes()...)
}
