package tx

import (
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/konstellation/konstellation/x/issue/types"
)

// getTxCmdDecreaseAllowance Decreases the allowance granted to `spender` by the caller.
func getTxCmdDecreaseAllowance() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "decrease-allowance [spender] [amount]",
		Args:  cobra.ExactArgs(2),
		Short: "Decreases the allowance granted to `spender` by the caller.",
		Long:  "Decreases the allowance granted to `spender` by the caller.",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			spender, err := sdk.AccAddressFromBech32(args[0])
			if err != nil {
				return err
			}

			coins, err := sdk.ParseCoinsNormalized(args[1])
			if err != nil {
				return err
			}

			msg := types.NewMsgDecreaseAllowance(clientCtx.GetFromAddress(), spender, coins)
			validateErr := msg.ValidateBasic()
			if validateErr != nil {
				return validateErr
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}

	return cmd
}
