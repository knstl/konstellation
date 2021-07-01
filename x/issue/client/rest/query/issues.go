package query

import (
	"github.com/konstellation/kn-sdk/x/issue/query"
	"net/http"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client/context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"

	"github.com/konstellation/kn-sdk/x/issue/types"
)

const (
	flagAddress = "address"
	flagOwner   = "owner"
	flagLimit   = "limit"
	flagSymbol  = "symbol"
)

// HTTP request handler to query specified issues
func issuesHandlerFn(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
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

		res, height, err := cliCtx.QueryWithData(query.PathQueryIssues(), bz)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		rest.PostProcessResponse(w, cliCtx.WithHeight(height), res)
	}
}
