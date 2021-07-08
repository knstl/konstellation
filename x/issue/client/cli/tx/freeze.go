package tx

import (
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/konstellation/konstellation/x/issue/types"
)

// getTxCmdFreeze implements burn function
func getTxCmdFreeze() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "freeze [denom] [holder_address] [op]",
		Args:  cobra.ExactArgs(3),
		Short: "Freeze tokens in holder",
		Long:  "Freeze tokens in holder \n Operations: in, out, in-out",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			holder, err := sdk.AccAddressFromBech32(args[1])
			if err != nil {
				return err
			}

			denom := args[0]
			op := args[2]

			msg := types.NewMsgFreeze(clientCtx.GetFromAddress(), holder, denom, op)
			validateErr := msg.ValidateBasic()
			if validateErr != nil {
				return validateErr
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}

	return cmd
}
