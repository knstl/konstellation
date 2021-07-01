package tx

import (
	"github.com/mitchellh/mapstructure"
	"net/http"

	"github.com/cosmos/cosmos-sdk/client/context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"

	"github.com/konstellation/kn-sdk/x/issue/types"
)

type (
	featuresRequest struct {
		BaseReq            rest.BaseReq `json:"base_req" yaml:"base_req"`
		Denom              string       `json:"denom" yaml:"denom"`
		BurnOwnerDisabled  bool         `json:"burn_owner_disabled" yaml:"burn_owner_disabled"`
		BurnHolderDisabled bool         `json:"burn_holder_disabled" yaml:"burn_holder_disabled"`
		BurnFromDisabled   bool         `json:"burn_from_disabled" yaml:"burn_from_disabled"`
		MintDisabled       bool         `json:"mint_disabled" yaml:"mint_disabled"`
		FreezeDisabled     bool         `json:"freeze_disabled" yaml:"freeze_disabled"`
	}
)

func featuresHandlerFn(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req featuresRequest
		if !rest.ReadRESTReq(w, r, cliCtx.Codec, &req) {
			return
		}

		req.BaseReq = req.BaseReq.Sanitize()
		if !req.BaseReq.ValidateBasic(w) {
			return
		}

		ownerAddr, err := sdk.AccAddressFromBech32(req.BaseReq.From)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		var issueFeatures types.IssueFeatures
		if err := mapstructure.Decode(req, &issueFeatures); err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		msg := types.NewMsgFeatures(ownerAddr, req.Denom, &issueFeatures)
		if err := msg.ValidateBasic(); err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		utils.WriteGenerateStdTxResponse(w, cliCtx, req.BaseReq, []sdk.Msg{msg})
	}
}
