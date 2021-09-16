package cli

import (
	"github.com/spf13/cobra"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/konstellation/konstellation/x/oracle/types"
)

var _ = strconv.Itoa(0)

func CmdDeleteExchangeRate() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete-exchange-rate [pair]",
		Short: "Broadcast message DeleteExchangeRate",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			pair := args[0]
			if pair == "" {
				return types.ErrInvalidPair
			}

			msg := types.NewMsgDeleteExchangeRate(clientCtx.GetFromAddress().String(), pair)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
