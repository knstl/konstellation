package tx

import (
	"net/http"

	"github.com/cosmos/cosmos-sdk/client/context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"

	"github.com/konstellation/kn-sdk/x/issue/types"
)

type (
	// freezeRequest defines the properties of transfer issues body.
	freezeRequest struct {
		BaseReq rest.BaseReq `json:"base_req" yaml:"base_req"`
		Holder  string       `json:"holder" yaml:"holder"`
		Denom   string       `json:"denom" yaml:"denom"`
		Op      string       `json:"op" yaml:"op"`
	}
)

func freezeHandlerFn(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req freezeRequest
		if !rest.ReadRESTReq(w, r, cliCtx.Codec, &req) {
			return
		}

		req.BaseReq = req.BaseReq.Sanitize()
		if !req.BaseReq.ValidateBasic(w) {
			return
		}

		freezerAddr, err := sdk.AccAddressFromBech32(req.BaseReq.From)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		holderAddr, err := sdk.AccAddressFromBech32(req.Holder)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		msg := types.NewMsgFreeze(freezerAddr, holderAddr, req.Denom, req.Op)
		if err := msg.ValidateBasic(); err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		utils.WriteGenerateStdTxResponse(w, cliCtx, req.BaseReq, []sdk.Msg{msg})
	}
}
