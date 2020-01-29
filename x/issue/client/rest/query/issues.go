package query

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/client/context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/gorilla/mux"
	"github.com/konstellation/konstellation/x/issue/types"
	"net/http"
	"strconv"
)

const (
	flagAddress = "address"
	flagOwner   = "owner"
	flagLimit   = "limit"
	flagSymbol  = "symbol"
)

func pathQueryIssues() string {
	return fmt.Sprintf("%s/%s/%s", types.Custom, types.QuerierRoute, types.QueryIssues)
}

// HTTP request handler to query specified issues
func issuesHandlerFn(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]
		fmt.Println(id)

		var params types.IssuesParams

		ownerAddr := r.URL.Query().Get(flagOwner)
		if len(ownerAddr) != 0 {
			_, err := sdk.AccAddressFromBech32(ownerAddr)
			if err != nil {
				rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
				return
			}
			params.AddOwner(ownerAddr)
		}

		limitStr := r.URL.Query().Get(flagLimit)
		if len(limitStr) == 0 {
			limitStr = "30"
		}
		limit, err := strconv.Atoi(limitStr)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}
		params.AddLimit(limit)

		bz, err := cliCtx.Codec.MarshalJSON(params)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		res, height, err := cliCtx.QueryWithData(pathQueryIssues(), bz)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		rest.PostProcessResponse(w, cliCtx.WithHeight(height), res)
	}
}
