package query

import (
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/gorilla/mux"
)

// RegisterRoutes registers the module REST routes.
func RegisterQueryRoutes(cliCtx context.CLIContext, r *mux.Router) {
	// Get all issues
	r.HandleFunc(
		"/issue/issues/all",
		issuesAllHandlerFn(cliCtx),
	).Methods("GET")
	r.HandleFunc(
		"/issue/issues",
		issuesHandlerFn(cliCtx),
	).Methods("GET")
	r.HandleFunc(
		"/issue/issue/{id}",
		issuesHandlerFn(cliCtx),
	).Methods("GET")
	r.HandleFunc(
		"/issue/allowance/{owner}/{spender}/{denom}",
		allowanceHandlerFn(cliCtx),
	).Methods("GET")
}
