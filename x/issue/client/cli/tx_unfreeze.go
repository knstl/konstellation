package cli

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/konstellation/konstellation/x/issue/types"
	"github.com/spf13/cobra"
)

func CmdUnfreeze() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "unfreeze [denom] [holder_address] [op]",
		Short: "Unfreeze tokens in holder",
		Long:  "Unfreeze tokens in holder \n Operations: in, out, in-out",
		Args:  cobra.ExactArgs(3),
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

			msg := types.NewMsgUnfreeze(clientCtx.GetFromAddress(), holder, denom, op)
			validateErr := msg.ValidateBasic()
			if validateErr != nil {
				return validateErr
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
