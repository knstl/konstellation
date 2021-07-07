package rest

import (
	"github.com/gorilla/mux"

	"github.com/cosmos/cosmos-sdk/client/context"

	"github.com/konstellation/konstellation/x/issue/client/rest/query"
	"github.com/konstellation/konstellation/x/issue/client/rest/tx"
)

// RegisterRoutes registers the auth module REST routes.
func RegisterRoutes(cliCtx context.CLIContext, r *mux.Router) {
	query.RegisterQueryRoutes(cliCtx, r)
	tx.RegisterTxRoutes(cliCtx, r)
}
