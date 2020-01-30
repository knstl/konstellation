package tx

import (
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/gorilla/mux"
)

// RegisterTxRoutes registers all transaction routes on the provided router.
func RegisterTxRoutes(cliCtx context.CLIContext, r *mux.Router) {
	r.HandleFunc(
		"/issue/issue",
		issueCreateHandlerFn(cliCtx),
	).Methods("POST")
	r.HandleFunc(
		"/issue/transfer",
		transferHandlerFn(cliCtx),
	).Methods("POST")
	r.HandleFunc(
		"/issue/transfer_from",
		transferFromHandlerFn(cliCtx),
	).Methods("POST")
	r.HandleFunc(
		"/issue/approve",
		approveHandlerFn(cliCtx),
	).Methods("POST")
	r.HandleFunc(
		"/issue/increase_allowance",
		increaseAllowanceHandlerFn(cliCtx),
	).Methods("POST")
	r.HandleFunc(
		"/issue/decrease_allowance",
		decreaseAllowanceHandlerFn(cliCtx),
	).Methods("POST")
	r.HandleFunc(
		"/issue/mint",
		mintHandlerFn(cliCtx),
	).Methods("POST")
	r.HandleFunc(
		"/issue/mint_to",
		mintToHandlerFn(cliCtx),
	).Methods("POST")
	r.HandleFunc(
		"/issue/burn",
		burnHandlerFn(cliCtx),
	).Methods("POST")
	r.HandleFunc(
		"/issue/burn_from",
		burnFromHandlerFn(cliCtx),
	).Methods("POST")
}
