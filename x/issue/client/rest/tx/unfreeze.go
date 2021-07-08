package tx

import (
	"net/http"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"

	"github.com/konstellation/konstellation/x/issue/types"
)

type (
	// unfreezeRequest defines the properties of transfer issues body.
	unfreezeRequest struct {
		BaseReq rest.BaseReq `json:"base_req" yaml:"base_req"`
		Holder  string       `json:"holder" yaml:"holder"`
		Denom   string       `json:"denom" yaml:"denom"`
		Op      string       `json:"op" yaml:"op"`
	}
)

func unfreezeHandlerFn(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req unfreezeRequest
		if !rest.ReadRESTReq(w, r, clientCtx.LegacyAmino, &req) {
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

		msg := types.NewMsgUnfreeze(freezerAddr, holderAddr, req.Denom, req.Op)
		if err := msg.ValidateBasic(); err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		tx.WriteGeneratedTxResponse(clientCtx, w, req.BaseReq, &msg)
	}
}
