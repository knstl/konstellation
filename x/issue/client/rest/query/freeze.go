package query

import (
	"net/http"

	"github.com/konstellation/konstellation/x/issue/query"

	"github.com/gorilla/mux"

	"github.com/cosmos/cosmos-sdk/client"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
)

const (
	flagHolder = "holder"
)

// HTTP request handler to query specified issues
func freezeHandlerFn(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		holder := vars[flagHolder]
		denom := vars[flagDenom]

		holderAddr, err := sdk.AccAddressFromBech32(holder)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		res, height, err := clientCtx.QueryWithData(query.PathQueryIssueFreeze(denom, holderAddr), nil)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		rest.PostProcessResponse(w, clientCtx.WithHeight(height), res)
	}
}
