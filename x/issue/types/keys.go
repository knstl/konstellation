package types

const (
	// ModuleName defines the module name
	ModuleName = "issue"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for slashing
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_issue"

	// this line is used by starport scaffolding # ibc/keys/name
	CoinDecimalsMaxValue     = uint(18)
	CoinDecimalsMultiple     = uint(3)
	CoinSymbolMinLength      = 2
	CoinSymbolMaxLength      = 8
	CoinDescriptionMaxLength = 1024
)

// this line is used by starport scaffolding # ibc/keys/port

func KeyPrefix(p string) []byte {
	return []byte(p)
}

const (
	AddressKey      = "Address-value-"
	AddressCountKey = "Address-count-"
)

const (
	AllowanceKey      = "Allowance-value-"
	AllowanceCountKey = "Allowance-count-"
)

const (
	AllowanceListKey      = "AllowanceList-value-"
	AllowanceListCountKey = "AllowanceList-count-"
)

const (
	CoinsKey      = "Coins-value-"
	CoinsCountKey = "Coins-count-"
)

const (
	FreezeKey      = "Freeze-value-"
	FreezeCountKey = "Freeze-count-"
)

const (
	AddressFreezeKey      = "AddressFreeze-value-"
	AddressFreezeCountKey = "AddressFreeze-count-"
)

const (
	AddressFreezeListKey      = "AddressFreezeList-value-"
	AddressFreezeListCountKey = "AddressFreezeList-count-"
)

const (
	CoinIssueKey      = "CoinIssue-value-"
	CoinIssueCountKey = "CoinIssue-count-"
)

const (
	CoinIssueDenomKey      = "CoinIssueDenom-value-"
	CoinIssueDenomCountKey = "CoinIssueDenom-count-"
)

const (
	IssuesKey      = "Issues-value-"
	IssuesCountKey = "Issues-count-"
)

const (
	ParamsKey      = "Params-value-"
	ParamsCountKey = "Params-count-"
)

const (
	IssueFeaturesKey      = "IssueFeatures-value-"
	IssueFeaturesCountKey = "IssueFeatures-count-"
)

const (
	IssueParamsKey      = "IssueParams-value-"
	IssueParamsCountKey = "IssueParams-count-"
)

const (
	IssuesParamsKey      = "IssuesParams-value-"
	IssuesParamsCountKey = "IssuesParams-count-"
)

const (
	CoinIssueListKey      = "CoinIssueList-value-"
	CoinIssueListCountKey = "CoinIssueList-count-"
)

const (
	CoinIssueDenomsKey      = "CoinIssueDenoms-value-"
	CoinIssueDenomsCountKey = "CoinIssueDenoms-count-"
)
