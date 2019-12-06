package query

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/konstellation/konstellation/x/issue/types"
)

func GetQueryIssuePath(issueID string) string {
	return fmt.Sprintf("%s/%s/%s/%s", "custom", types.QuerierRoute, types.QueryIssue, issueID)
}

func GetQueryIssueSearchPath(symbol string) string {
	return fmt.Sprintf("%s/%s/%s/%s", "custom", types.QuerierRoute, types.QuerySearch, symbol)
}

func GetQueryParamsPath() string {
	return fmt.Sprintf("%s/%s/%s", "custom", types.QuerierRoute, types.QueryParams)
}

func QueryIssueBySymbol(symbol string, cliCtx context.CLIContext) ([]byte, int64, error) {
	return cliCtx.QueryWithData(GetQueryIssueSearchPath(symbol), nil)
}

func QueryIssueByID(issueId string, cliCtx context.CLIContext) ([]byte, int64, error) {
	return cliCtx.QueryWithData(GetQueryIssuePath(issueId), nil)
}

func QueryParams(cliCtx context.CLIContext) ([]byte, int64, error) {
	return cliCtx.QueryWithData(GetQueryParamsPath(), nil)
}
