package cli

import (
	"github.com/spf13/cobra"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/konstellation/konstellation/x/oracle/types"
)

var _ = strconv.Itoa(0)

func CmdSetAdminAddr() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "set-admin-addr",
		Short: "Broadcast message SetAdminAddr --add [address-to-add] --delete [address-to-delete]",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			add, _ := cmd.Flags().GetStringSlice(FlagAdd)
			del, _ := cmd.Flags().GetStringSlice(FlagDelete)

			var addressesToAdd []*types.AdminAddr
			var addressesToDelete []*types.AdminAddr
			for _, a := range add {
				_, err := sdk.AccAddressFromBech32(a)
				if err != nil {
					return err
				}
				addressesToAdd = append(addressesToAdd, types.NewAdminAddr(a))
			}

			for _, a := range del {
				_, err := sdk.AccAddressFromBech32(a)
				if err != nil {
					return err
				}

				addressesToDelete = append(addressesToDelete, types.NewAdminAddr(a))
			}

			msg := types.NewMsgSetAdminAddr(clientCtx.GetFromAddress().String(), addressesToAdd, addressesToDelete)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
