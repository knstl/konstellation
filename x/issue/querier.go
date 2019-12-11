package issue

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/konstellation/konstellation/x/issue/keeper"
	"github.com/konstellation/konstellation/x/issue/query"
	"github.com/konstellation/konstellation/x/issue/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

// NewQuerier creates a querier for auth REST endpoints
func NewQuerier(k keeper.Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) ([]byte, sdk.Error) {
		switch path[0] {
		case types.QueryIssue:
			return query.Issue(ctx, k, path[1])
		case types.QueryIssues:
			return query.Issues(ctx, k, req.Data)
		case types.QueryIssuesAll:
			return query.IssuesAll(ctx, k)
		case types.QueryAllowance:
			return query.Allowance(ctx, k, path[1], path[2], path[3])
		default:
			return nil, sdk.ErrUnknownRequest("unknown auth query endpoint")
		}
	}
}
