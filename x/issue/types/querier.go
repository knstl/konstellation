package types

import (
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

func NewIssueFeatures(data interface{}) (*IssueFeatures, error) {
	var features IssueFeatures
	err := mapstructure.Decode(data, &features)
	return &features, err
}

func NewIssueParams(data interface{}) (*IssueParams, error) {
	var issue IssueParams
	err := mapstructure.Decode(data, &issue)
	return &issue, err
}

func (ip *IssueParams) AddTotalSupply(totalSupply *sdk.Int) {
	ip.TotalSupply = sdk.NewIntWithDecimal(totalSupply.Int64(), cast.ToInt(ip.Decimals))
}

func NewIssuesParams(owner string, limit int32) IssuesParams {
	return IssuesParams{owner, limit}
}

func (ip *IssuesParams) AddOwner(owner string) {
	ip.Owner = owner
}

func (ip *IssuesParams) AddLimit(limit int32) {
	ip.Limit = limit
}
