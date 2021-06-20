package types

// oracle module event types
const (
	EventTypeSetExchangeRate     = "set_exchage_rate"
	EventTypeSetExchangeRates    = "set_exchage_rates"
	EventTypeDeleteExchangeRate  = "delete_exchage_rate"
	EventTypeDeleteExchangeRates = "delete_exchage_rates"
	EventTypeSetAdminAddr        = "set_admin_addr"

	AttributeValueCategory = ModuleName
	AttributeKeyPair       = "pair"
	AttributeKeyRate       = "rate"
	AttributeKeyDenoms     = "denoms"
	AttributeKeyTimestamp  = "timestamp"
	AttributeKeyAdd        = "add"
	AttributeKeyDelete     = "delete"
)
