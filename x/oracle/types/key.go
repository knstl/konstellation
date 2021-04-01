package types

var (
	SetExchangeRateKey    = []byte{0x00}
	DeleteExchangeRateKey = []byte{0x01}
)

const (
	// ModuleName is the name of the module
	ModuleName = "oracle"

	RouterKey = ModuleName

	// StoreKey is the default store key for mint
	StoreKey = ModuleName

	QueryExchangeRate = "exchange-rate"
)
