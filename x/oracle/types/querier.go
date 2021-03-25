package types

const QueryExchangeRate = "exchange-rate"

// QueryResResolve Queries Result Payload for a resolve query
type QueryResExchangeRate struct {
	Value string `json:"value"`
}

// implement fmt.Stringer
func (r QueryResExchangeRate) String() string {
	return r.Value
}
