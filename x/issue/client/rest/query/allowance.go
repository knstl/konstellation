package query

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/cosmos/cosmos-sdk/client"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
)

const (
	flagSpender = "spender"
	flagDenom   = "denom"
)

// HTTP request handler to query specified issues
func allowanceHandlerFn(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		owner := vars[flagOwner]
		spender := vars[flagSpender]
		denom := vars[flagDenom]

		ownerAddr, err := sdk.AccAddressFromBech32(owner)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		spenderAddr, err := sdk.AccAddressFromBech32(spender)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		res, height, err := clientCtx.QueryWithData(pathQueryIssueAllowance(ownerAddr, spenderAddr, denom), nil)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		rest.PostProcessResponse(w, clientCtx.WithHeight(height), res)
	}
}
