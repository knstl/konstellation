package tx

import (
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"

	"github.com/konstellation/konstellation/x/issue/types"
)

func getTxCmdDescription() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "description [denom] [description]",
		Args:  cobra.ExactArgs(2),
		Short: "Change issue's description",
		Long:  "Change issue's description",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			denom := args[0]
			description := args[1]

			msg := types.NewMsgDescription(clientCtx.GetFromAddress(), denom, description)
			validateErr := msg.ValidateBasic()
			if validateErr != nil {
				return validateErr
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}

	return cmd
}
