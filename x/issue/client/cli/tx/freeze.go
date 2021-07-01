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

// getTxCmdFreeze implements burn function
func getTxCmdFreeze(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "freeze [denom] [holder_address] [op]",
		Args:  cobra.ExactArgs(3),
		Short: "Freeze tokens in holder",
		Long:  "Freeze tokens in holder \n Operations: in, out, in-out",
		RunE: func(cmd *cobra.Command, args []string) error {
			txBldr := auth.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			holder, err := sdk.AccAddressFromBech32(args[1])
			if err != nil {
				return err
			}

			denom := args[0]
			op := args[2]

			msg := types.NewMsgFreeze(cliCtx.GetFromAddress(), holder, denom, op)
			validateErr := msg.ValidateBasic()
			if validateErr != nil {
				return validateErr
			}

			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}

	return cmd
}
