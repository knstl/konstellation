package query

import (
	"net/http"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"

	"github.com/konstellation/konstellation/x/issue/types"
)

const (
	flagAddress = "address"
	flagOwner   = "owner"
	flagLimit   = "limit"
	flagSymbol  = "symbol"
)

// HTTP request handler to query specified issues
func issuesHandlerFn(clientCtx client.Context) http.HandlerFunc {
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
		params.AddLimit(int32(limit))

		bz, err := clientCtx.LegacyAmino.MarshalJSON(params)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		res, height, err := clientCtx.QueryWithData(pathQueryIssues(), bz)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		rest.PostProcessResponse(w, clientCtx.WithHeight(height), res)
	}
}
