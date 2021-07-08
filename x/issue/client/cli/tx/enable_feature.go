package tx

import (
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"

	"github.com/konstellation/konstellation/x/issue/types"
)

// getTxCmdEnable feature
func getTxCmdEnableFeature() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "enable-feature [denom] [feature]",
		Args:  cobra.ExactArgs(2),
		Short: "Enable feature",
		Long:  "Enable feature for token",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			denom := args[0]
			feature := args[1]

			msg := types.NewMsgEnableFeature(clientCtx.GetFromAddress(), denom, feature)
			validateErr := msg.ValidateBasic()
			if validateErr != nil {
				return validateErr
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}

	return cmd
}
