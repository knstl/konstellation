package tx

import (
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/konstellation/konstellation/x/issue/types"
)

// getTxCmdApprove Sets `amount` as the allowance of `spender` over the caller's tokens.
func getTxCmdApprove() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "approve [spender] [amount]",
		Args:  cobra.ExactArgs(2),
		Short: "Sets `amount` as the allowance of `spender` over the caller's tokens.",
		Long:  "Sets `amount` as the allowance of `spender` over the caller's tokens.",
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

			msg := types.NewMsgApprove(clientCtx.GetFromAddress(), spender, coins)
			validateErr := msg.ValidateBasic()
			if validateErr != nil {
				return validateErr
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}

	return cmd
}
