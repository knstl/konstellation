package query

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/gorilla/mux"
	"github.com/konstellation/konstellation/x/issue/types"
	"net/http"
)

func pathQueryIssue() string {
	return fmt.Sprintf("%s/%s/%s", types.Custom, types.QuerierRoute, types.QueryIssue)
}

// HTTP request handler to query specified issues
func issueHandlerFn(cliCtx context.CLIContext, r *mux.Router) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]
		fmt.Println(id)
		//
		//var params types.IssuesParams
		//
		//ownerAddr := r.URL.Query().Get("owner")
		//if len(ownerAddr) != 0 {
		//	_, err := sdk.AccAddressFromBech32(ownerAddr)
		//	if err != nil {
		//		rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
		//		return
		//	}
		//	params.AddOwner(ownerAddr)
		//}
		//
		//limitStr := r.URL.Query().Get("limit")
		//if len(limitStr) != 0 {
		//	limit, err := strconv.Atoi(limitStr)
		//	if err != nil {
		//		rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
		//		return
		//	}
		//	params.AddLimit(limit)
		//}
		//
		//bz, err := cliCtx.Codec.MarshalJSON(params)
		//if err != nil {
		//	rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
		//	return
		//}
		//
		//res, height, err := cliCtx.QueryWithData(pathQueryIssues(), bz)
		//if err != nil {
		//	rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
		//	return
		//}
		//
		//cliCtx = cliCtx.WithHeight(height)
		//rest.PostProcessResponse(w, cliCtx, res)
	}
}
