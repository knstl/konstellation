package tx

import (
	"github.com/gorilla/mux"

	"github.com/cosmos/cosmos-sdk/client/context"
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
		"/issue/burn",
		burnHandlerFn(cliCtx),
	).Methods("POST")
	r.HandleFunc(
		"/issue/burn_from",
		burnFromHandlerFn(cliCtx),
	).Methods("POST")
	r.HandleFunc(
		"/issue/transfer_ownership",
		transferOwnershipHandlerFn(cliCtx),
	).Methods("POST")
	r.HandleFunc(
		"/issue/disable_feature",
		disableFeatureHandlerFn(cliCtx),
	).Methods("POST")
	r.HandleFunc(
		"/issue/enable_feature",
		enableFeatureHandlerFn(cliCtx),
	).Methods("POST")
	r.HandleFunc(
		"/issue/features",
		featuresHandlerFn(cliCtx),
	).Methods("POST")
	r.HandleFunc(
		"/issue/description",
		descriptionHandlerFn(cliCtx),
	).Methods("POST")
	r.HandleFunc(
		"/issue/freeze",
		freezeHandlerFn(cliCtx),
	).Methods("POST")
	r.HandleFunc(
		"/issue/unfreeze",
		unfreezeHandlerFn(cliCtx),
	).Methods("POST")
}
