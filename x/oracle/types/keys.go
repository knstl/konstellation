package types

var (
	AllowedAddressKey = []byte{0x00}
	ExchangeRateKey   = []byte{0x01}
)

const (
	// ModuleName is the name of the module
	ModuleName = "oracle"

	RouterKey = ModuleName

	// StoreKey is the default store key for mint
	StoreKey     = ModuleName
	QuerierRoute = StoreKey

	QueryExchangeRate = "exchange_rate"
)
