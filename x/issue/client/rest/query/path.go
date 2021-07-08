package query

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/konstellation/konstellation/x/issue/types"
)

func pathQueryIssueAllowance(owner sdk.AccAddress, spender sdk.AccAddress, denom string) string {
	return fmt.Sprintf("%s/%s/%s/%s/%s/%s", types.Custom, types.QuerierRoute, types.QueryAllowance, denom, owner.String(), spender.String())
}

func pathQueryIssueAllowances(owner sdk.AccAddress, denom string) string {
	return fmt.Sprintf("%s/%s/%s/%s/%s", types.Custom, types.QuerierRoute, types.QueryAllowances, denom, owner.String())
}

func pathQueryIssue(denom string) string {
	return fmt.Sprintf("%s/%s/%s/%s", types.Custom, types.QuerierRoute, types.QueryIssue, denom)
}

func pathQueryIssues() string {
	return fmt.Sprintf("%s/%s/%s", types.Custom, types.QuerierRoute, types.QueryIssues)
}

func pathQueryIssuesAll() string {
	return fmt.Sprintf("%s/%s/%s", types.Custom, types.QuerierRoute, types.QueryIssuesAll)
}

func pathParams() string {
	return fmt.Sprintf("%s/%s/%s", types.Custom, types.QuerierRoute, types.QueryParams)
}

func pathQueryIssueFreeze(denom string, holder sdk.AccAddress) string {
	return fmt.Sprintf("%s/%s/%s/%s/%s", types.Custom, types.QuerierRoute, types.QueryFreeze, denom, holder.String())
}

func pathQueryIssueFreezes(denom string) string {
	return fmt.Sprintf("%s/%s/%s/%s", types.Custom, types.QuerierRoute, types.QueryFreezes, denom)
}
