package types

type QueryAllExchangeRatesParams struct {
	Page, Limit int
}

func NewQueryAllExchangeRatesParams(page, limit int) QueryAllExchangeRatesParams {
	return QueryAllExchangeRatesParams{
		Page:  page,
		Limit: limit,
	}
}

type QueryExchangeRateParams struct {
	Pair string
}

func NewQueryExchangeRateParams(pair string) QueryExchangeRateParams {
	return QueryExchangeRateParams{
		Pair: pair,
	}
}
