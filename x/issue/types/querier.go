package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/cast"
)

// querier keys
const (
	QueryParams    = "params"
	QueryIssue     = "query"
	QueryIssues    = "list"
	QueryIssuesAll = "list-all"
	QuerySearch    = "search"
	QueryAllowance = "allowance"
)

type IssueParams struct {
	Denom              string  `json:"denom"`
	Symbol             string  `json:"symbol"`
	TotalSupply        sdk.Int `json:"total_supply"`
	Decimals           uint    `json:"decimals"`
	Description        string  `json:"description"`
	BurnOwnerDisabled  bool    `json:"burn_owner_disabled"`
	BurnHolderDisabled bool    `json:"burn_holder_disabled"`
	BurnFromDisabled   bool    `json:"burn_from_disabled"`
	MintingFinished    bool    `json:"minting_finished"`
	FreezeDisabled     bool    `json:"freeze_disabled"`
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
