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

func CmdTransferFrom() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "transfer-from [from_address] [to_address] [amount]",
		Short: "Broadcast message TransferFrom",
		Long:  "Moves `amount` tokens from `sender` to `recipient` using the allowance mechanism. `amount` is then deducted from the caller's allowance.",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			fromAddr, err := sdk.AccAddressFromBech32(args[0])
			if err != nil {
				return err
			}

			toAddr, err := sdk.AccAddressFromBech32(args[1])
			if err != nil {
				return err
			}

			coins, err := sdk.ParseCoinsNormalized(args[2])
			if err != nil {
				return err
			}

			msg := types.NewMsgTransferFrom(clientCtx.GetFromAddress(), fromAddr, toAddr, coins)
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
