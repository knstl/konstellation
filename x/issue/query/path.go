package query

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/konstellation/kn-sdk/x/issue/types"
)

func PathQueryIssueAllowance(owner sdk.AccAddress, spender sdk.AccAddress, denom string) string {
	return fmt.Sprintf("%s/%s/%s/%s/%s/%s", types.Custom, types.QuerierRoute, types.QueryAllowance, denom, owner.String(), spender.String())
}

func PathQueryIssueAllowances(owner sdk.AccAddress, denom string) string {
	return fmt.Sprintf("%s/%s/%s/%s/%s", types.Custom, types.QuerierRoute, types.QueryAllowances, denom, owner.String())
}

func PathQueryIssue(denom string) string {
	return fmt.Sprintf("%s/%s/%s/%s", types.Custom, types.QuerierRoute, types.QueryIssue, denom)
}

func PathQueryIssues() string {
	return fmt.Sprintf("%s/%s/%s", types.Custom, types.QuerierRoute, types.QueryIssues)
}

func PathQueryIssuesAll() string {
	return fmt.Sprintf("%s/%s/%s", types.Custom, types.QuerierRoute, types.QueryIssuesAll)
}

func PathParams() string {
	return fmt.Sprintf("%s/%s/%s", types.Custom, types.QuerierRoute, types.QueryParams)
}

func PathQueryIssueFreeze(denom string, holder sdk.AccAddress) string {
	return fmt.Sprintf("%s/%s/%s/%s/%s", types.Custom, types.QuerierRoute, types.QueryFreeze, denom, holder.String())
}

func PathQueryIssueFreezes(denom string) string {
	return fmt.Sprintf("%s/%s/%s/%s", types.Custom, types.QuerierRoute, types.QueryFreezes, denom)
}
