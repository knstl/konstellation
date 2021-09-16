package cli

import (
	"github.com/spf13/cobra"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/konstellation/konstellation/x/issue/types"
)

var _ = strconv.Itoa(0)

func CmdIncreaseAllowance() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "increase-allowance [spender] [amount]",
		Short: "Broadcast message IncreaseAllowance",
		Long:  "Increases the allowance granted to `spender` by the caller.",
		Args:  cobra.ExactArgs(2),
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

			msg := types.NewMsgIncreaseAllowance(clientCtx.GetFromAddress(), spender, coins)
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
