package query

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/gorilla/mux"
)

// RegisterRoutes registers the module REST routes.
func RegisterQueryRoutes(clientCtx client.Context, r *mux.Router) {
	// Get all issues
	r.HandleFunc(
		"/issue/issues/all",
		issuesAllHandlerFn(clientCtx),
	).Methods("GET")
	r.HandleFunc(
		"/issue/issues",
		issuesHandlerFn(clientCtx),
	).Methods("GET")
	r.HandleFunc(
		"/issue/issue/{denom}",
		issueHandlerFn(clientCtx),
	).Methods("GET")
	r.HandleFunc(
		"/issue/allowance/{owner}/{spender}/{denom}",
		allowanceHandlerFn(clientCtx),
	).Methods("GET")
	r.HandleFunc(
		"/issue/allowances/{owner}/{denom}",
		allowancesHandlerFn(clientCtx),
	).Methods("GET")
	r.HandleFunc(
		"/issue/freeze/{holder}/{denom}",
		freezeHandlerFn(clientCtx),
	).Methods("GET")
	r.HandleFunc(
		"/issue/freezes/{denom}",
		freezesHandlerFn(clientCtx),
	).Methods("GET")
	r.HandleFunc(
		"/issue/params",
		paramsHandlerFn(clientCtx),
	).Methods("GET")
}
