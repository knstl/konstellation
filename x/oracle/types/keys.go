package types

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
	// Keys for store prefixes
	// Last* values are constant during a block.
	//LastValidatorPowerKey = []byte{0x11} // prefix for each key to a validator index, for bonded validators
	//LastTotalPowerKey     = []byte{0x12} // prefix for the total power

	AllowedAddressKey = []byte{0x21} // prefix for each key to a validator

	ExchangeRateKey = []byte{0x31} // key for a delegation
)

// gets the key for the validator with address
// VALUE: oracle/ExchangeRate
func GetExchangeRateKey(pair string) []byte {
	return append(ExchangeRateKey, []byte(pair)...)
}
