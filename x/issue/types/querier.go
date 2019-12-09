package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/mitchellh/mapstructure"
)

// querier keys
const (
	QueryParams                      = "params"
	QueryIssue                       = "query"
	QueryIssues                      = "list"
	QuerySearch                      = "search"
	QueryAllowance                   = "allowance"
	QueryValidatorOutstandingRewards = "validator_outstanding_rewards"

	ParamCommunityTax        = "community_tax"
	ParamBaseProposerReward  = "base_proposer_reward"
	ParamBonusProposerReward = "bonus_proposer_reward"
	ParamWithdrawAddrEnabled = "withdraw_addr_enabled"
)

type IssueParams struct {
	Name               string  `json:"name"`
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

type IssuesParams struct {
	StartIssueId string `json:"start_issue_id"`
	Owner        string `json:"owner"`
	Limit        int    `json:"limit"`
}

func NewIssuesParams(startIssueId string, owner string, limit int) IssuesParams {
	return IssuesParams{startIssueId, owner, limit}
}
