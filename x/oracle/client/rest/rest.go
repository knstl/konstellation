package rest

import (
	"github.com/gorilla/mux"

	"github.com/cosmos/cosmos-sdk/client"
	//"github.com/cosmos/cosmos-sdk/client/rest"
)

// RegisterRoutes registers minting module REST handlers on the provided router.
func RegisterHandlers(clientCtx client.Context, rtr *mux.Router) {
	//r := rest.WithHTTPDeprecationHeaders(rtr)
	//registerQueryRoutes(clientCtx, r)
}
