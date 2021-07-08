package rest

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/gorilla/mux"

	"github.com/konstellation/konstellation/x/issue/client/rest/query"
	"github.com/konstellation/konstellation/x/issue/client/rest/tx"
)

// RegisterRoutes registers the auth module REST routes.
func RegisterRoutes(clientCtx client.Context, r *mux.Router) {
	query.RegisterQueryRoutes(clientCtx, r)
	tx.RegisterTxRoutes(clientCtx, r)
}
