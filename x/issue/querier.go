package issue

import (
	abci "github.com/tendermint/tendermint/abci/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/konstellation/konstellation/x/issue/keeper"
	"github.com/konstellation/konstellation/x/issue/query"
	"github.com/konstellation/konstellation/x/issue/types"
)

// NewQuerier creates a querier for auth REST endpoints
func NewQuerier(k keeper.Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) ([]byte, *sdkerrors.Error) {
		switch path[0] {
		case types.QueryIssue:
			return query.Issue(ctx, k, path[1])
		case types.QueryIssues:
			return query.Issues(ctx, k, req.Data)
		case types.QueryIssuesAll:
			return query.IssuesAll(ctx, k)
		case types.QueryAllowance:
			return query.Allowance(ctx, k, path[1], path[2], path[3])
		case types.QueryAllowances:
			return query.Allowances(ctx, k, path[1], path[2])
		case types.QueryFreeze:
			return query.Freeze(ctx, k, path[1], path[2])
		case types.QueryFreezes:
			return query.Freezes(ctx, k, path[1])
		case types.QueryParams:
			return query.Params(ctx, k)
		default:
			return nil, sdkerrors.ErrUnknownRequest("unknown issue query endpoint")
		}
	}
}
