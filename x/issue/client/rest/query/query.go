package query

import (
	"github.com/gorilla/mux"

	"github.com/cosmos/cosmos-sdk/client/context"
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
		"/issue/issue/{denom}",
		issueHandlerFn(cliCtx),
	).Methods("GET")
	r.HandleFunc(
		"/issue/allowance/{owner}/{spender}/{denom}",
		allowanceHandlerFn(cliCtx),
	).Methods("GET")
	r.HandleFunc(
		"/issue/allowances/{owner}/{denom}",
		allowancesHandlerFn(cliCtx),
	).Methods("GET")
	r.HandleFunc(
		"/issue/freeze/{holder}/{denom}",
		freezeHandlerFn(cliCtx),
	).Methods("GET")
	r.HandleFunc(
		"/issue/freezes/{denom}",
		freezesHandlerFn(cliCtx),
	).Methods("GET")
	r.HandleFunc(
		"/issue/params",
		paramsHandlerFn(cliCtx),
	).Methods("GET")
}
