package types

type NetConfig struct {
	Type         string           `json:"type"` // localnet,testnet
	GlobalConfig *Config          `json:"config"`
	GenAccounts  []*GenAccount    `json:"gen_accounts"`
	Validators   []*ValidatorInfo `json:"validators"`
}
