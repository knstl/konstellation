package types

import (
	"encoding/json"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/cast"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// querier keys
const (
	QueryParams     = "params"
	QueryIssue      = "query"
	QueryIssues     = "list"
	QueryIssuesAll  = "list-all"
	QueryFreeze     = "freeze"
	QueryFreezes    = "freezes"
	QuerySearch     = "search"
	QueryAllowance  = "allowance"
	QueryAllowances = "allowances"
)

type IssueFeatures struct {
	BurnOwnerDisabled  bool `json:"burn_owner_disabled"`
	BurnHolderDisabled bool `json:"burn_holder_disabled"`
	BurnFromDisabled   bool `json:"burn_from_disabled"`
	MintDisabled       bool `json:"mint_disabled"`
	FreezeDisabled     bool `json:"freeze_disabled"`
}

func NewIssueFeatures(data interface{}) (*IssueFeatures, error) {
	var features IssueFeatures
	err := mapstructure.Decode(data, &features)
	return &features, err
}

func (fs *IssueFeatures) String() string {
	f, _ := json.Marshal(fs)
	return string(f)
}

type IssueParams struct {
	Denom         string  `json:"denom"`
	Symbol        string  `json:"symbol"`
	TotalSupply   sdk.Int `json:"total_supply"`
	Decimals      uint    `json:"decimals"`
	Description   string  `json:"description"`
	IssueFeatures `json:"features"`
}

func NewIssueParams(data interface{}) (*IssueParams, error) {
	var issue IssueParams
	err := mapstructure.Decode(data, &issue)
	return &issue, err
}

func (ip *IssueParams) AddTotalSupply(totalSupply *sdk.Int) {
	ip.TotalSupply = sdk.NewIntWithDecimal(totalSupply.Int64(), cast.ToInt(ip.Decimals))
}

type IssuesParams struct {
	Owner string `json:"owner"`
	Limit int    `json:"limit"`
}

func NewIssuesParams(owner string, limit int) IssuesParams {
	return IssuesParams{owner, limit}
}

func (ip *IssuesParams) AddOwner(owner string) {
	ip.Owner = owner
}

func (ip *IssuesParams) AddLimit(limit int) {
	ip.Limit = limit
}
