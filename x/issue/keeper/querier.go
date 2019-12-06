package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/konstellation/konstellation/x/issue/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

// NewQuerier creates a querier for auth REST endpoints
func NewQuerier(keeper Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) ([]byte, sdk.Error) {
		switch path[0] {
		//case types.QueryParams:
		//	return QueryParams(ctx, keeper)
		case types.QueryIssue:
			return QueryIssue(ctx, path[1], keeper)
		//case types.QueryAccount:
		//	return queryAccount(ctx, req, keeper)
		default:
			return nil, sdk.ErrUnknownRequest("unknown auth query endpoint")
		}
	}
}
