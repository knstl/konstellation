package tx

import (
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"

	"github.com/konstellation/kn-sdk/x/issue/types"
)

// getTxCmdDisable feature
func getTxCmdDisableFeature(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "disable-feature [denom] [feature]",
		Args:  cobra.ExactArgs(2),
		Short: "Disable feature",
		Long:  "Disable feature for token",
		RunE: func(cmd *cobra.Command, args []string) error {
			txBldr := auth.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			denom := args[0]
			feature := args[1]

			msg := types.NewMsgDisableFeature(cliCtx.GetFromAddress(), denom, feature)
			validateErr := msg.ValidateBasic()
			if validateErr != nil {
				return validateErr
			}

			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}

	return cmd
}
