package query

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/client/context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/gorilla/mux"
	"github.com/konstellation/konstellation/x/issue/types"
	"net/http"
)

const (
	flagSpender = "spender"
	flagDenom   = "denom"
)

func pathQueryIssueAllowance(owner sdk.AccAddress, spender sdk.AccAddress, denom string) string {
	return fmt.Sprintf("%s/%s/%s/%s/%s/%s", types.Custom, types.QuerierRoute, types.QueryAllowance, denom, owner.String(), spender.String())
}

// HTTP request handler to query specified issues
func allowanceHandlerFn(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		owner := vars[flagOwner]
		spender := vars[flagSpender]
		denom := vars[flagDenom]

		//denom := r.URL.Query().Get(flagDenom)
		//if denom == "" {
		//	rest.WriteErrorResponse(w, http.StatusBadRequest, "Empty denom")
		//	return
		//}
		//owner := r.URL.Query().Get(flagOwner)
		//if owner == "" {
		//	rest.WriteErrorResponse(w, http.StatusB																					adRequest, "Empty owner")
		//	return
		//}
		//spender := r.URL.Query().Get(flagSpender)
		//if spender == "" {
		//	rest.WriteErrorResponse(w, http.StatusBadRequest, "Empty spender")
		//	return
		//}

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

		res, height, err := cliCtx.QueryWithData(pathQueryIssueAllowance(ownerAddr, spenderAddr, denom), nil)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		rest.PostProcessResponse(w, cliCtx.WithHeight(height), res)
	}
}
