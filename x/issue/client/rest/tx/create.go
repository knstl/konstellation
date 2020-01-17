package tx

import (
	"github.com/cosmos/cosmos-sdk/client/context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"
	"github.com/konstellation/konstellation/x/issue/types"
	"github.com/mitchellh/mapstructure"
	"net/http"
)

type (
	// issueCreateRequest defines the properties of issue create's body.
	issueCreateRequest struct {
		BaseReq            rest.BaseReq `json:"base_req" yaml:"base_req"`
		Denom              string       `json:"denom" yaml:"denom"`
		Symbol             string       `json:"symbol" yaml:"symbol"`
		TotalSupply        sdk.Int      `json:"total_supply" yaml:"total_supply"  mapstructure:"total_supply"`
		Decimals           uint         `json:"decimals" yaml:"decimals"`
		Description        string       `json:"description" yaml:"description"`
		BurnOwnerDisabled  bool         `json:"burn_owner_disabled" yaml:"burn_owner_disabled"`
		BurnHolderDisabled bool         `json:"burn_holder_disabled" yaml:"burn_holder_disabled"`
		BurnFromDisabled   bool         `json:"burn_from_disabled" yaml:"burn_from_disabled"`
		MintingFinished    bool         `json:"minting_finished" yaml:"minting_finished"`
		FreezeDisabled     bool         `json:"freeze_disabled" yaml:"freeze_disabled"`
	}
)

/**

POST localhost:1317/issue/issue
{
  "base_req": {
    "from": "darc1d22ccl8xpzzzldl28l9gs9htrgaatkaxjwskkl",
    "memo": "test issue",
    "chain_id": "darchub",
    "account_number": "0",
    "sequence": "1",
    "gas": "94337",
    "gas_adjustment": "1.2",
    "fees": [
      {
        "denom": "darc",
        "amount": "50"
      }
    ],
    "simulate": false
  },
  "denom": "knstl",
  "symbol": "KNSTL",
  "total_supply": "1000",
  "decimals": "3",
  "description": "xcvx"
}

Response
{
    "type": "cosmos-sdk/StdTx",
    "value": {
        "msg": [
            {
                "type": "issue/MsgIssue",
                "value": {
                    "owner": "darc1d22ccl8xpzzzldl28l9gs9htrgaatkaxjwskkl",
                    "issuer": "darc1d22ccl8xpzzzldl28l9gs9htrgaatkaxjwskkl",
                    "params": {
                        "denom": "knstl",
                        "symbol": "KNSTL",
                        "total_supply": "1000000",
                        "decimals": "3",
                        "description": "xcvx",
                        "burn_owner_disabled": false,
                        "burn_holder_disabled": false,
                        "burn_from_disabled": false,
                        "minting_finished": false,
                        "freeze_disabled": false
                    }
                }
            }
        ],
        "fee": {
            "amount": [
                {
                    "denom": "darc",
                    "amount": "50"
                }
            ],
            "gas": "94337"
        },
        "signatures": null,
        "memo": "test issue"
    }
}
*/
func issueCreateHandlerFn(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req issueCreateRequest
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

		var issueParams types.IssueParams
		if err := mapstructure.Decode(req, &issueParams); err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}
		issueParams.AddTotalSupply(&req.TotalSupply)

		msg := types.NewMsgIssueCreate(ownerAddr, ownerAddr, &issueParams)
		if err := msg.ValidateBasic(); err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		utils.WriteGenerateStdTxResponse(w, cliCtx, req.BaseReq, []sdk.Msg{msg})
	}
}
