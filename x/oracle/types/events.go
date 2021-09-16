package types

// oracle module event types
const (
	EventTypeSetExchangeRate     = "set_exchange_rate"
	EventTypeSetExchangeRates    = "set_exchange_rates"
	EventTypeDeleteExchangeRate  = "delete_exchange_rate"
	EventTypeDeleteExchangeRates = "delete_exchange_rates"
	EventTypeSetAdminAddr        = "set_admin_addr"

	AttributeValueCategory = ModuleName
	AttributeKeyPair       = "pair"
	AttributeKeyRate       = "rate"
	AttributeKeyDenoms     = "denoms"
	AttributeKeyTimestamp  = "timestamp"
	AttributeKeyAdd        = "add"
	AttributeKeyDelete     = "delete"
)
