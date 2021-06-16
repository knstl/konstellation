package cli

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"strconv"
	"strings"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"

	//	"github.com/cosmos/cosmos-sdk/types/msgservice"

	"github.com/konstellation/konstellation/x/oracle/types"
)

const RateUnit = 1000000000000000000

func NewTxCmd() *cobra.Command {
	txCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Exchange Rate subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	txCmd.AddCommand(
		NewMsgSetExchangeRateCmd(),
		NewMsgDeleteExchangeRateCmd(),
		NewMsgSetAdminAddrCmd(),
	)

	return txCmd
}

func NewMsgSetExchangeRateCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "set-exchange-rate [pair] [rate] [denoms]",
		Short: "Set exchange rate",
		Long:  `Set exchange rate for pair: kbtckusd 5000000000000 kbtc,kusd`,
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			pair, rateStr, denomsStr := args[0], args[1], args[2]
			if pair == "" {
				return types.ErrInvalidPair
			}

			denoms := strings.Split(denomsStr, ",")
			if len(denoms) != 2 {
				return types.ErrInvalidDenoms
			}

			rate, err := strconv.ParseUint(rateStr, 10, 64)
			if err != nil {
				return types.ErrInvalidRate
			}

			exchangeRate := types.ExchangeRate{
				Pair:   pair,
				Rate:   rate,
				Denoms: denoms,
			}

			msg := types.NewMsgSetExchangeRate(clientCtx.GetFromAddress(), &exchangeRate)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func NewMsgDeleteExchangeRateCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete-exchange-rate [pair]",
		Short: "Delete exchange rate",
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

			msg := types.NewMsgDeleteExchangeRate(clientCtx.GetFromAddress(), pair)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func NewMsgSetAdminAddrCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "set-admin-addr --add [address-to-add] --delete [address-to-delete]",
		Short: "Set Admin Address",
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

			msg := types.NewMsgSetAdminAddr(clientCtx.GetFromAddress(), addressesToAdd, addressesToDelete)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}
	cmd.Flags().StringSlice(FlagAdd, nil, "addr,addr2,addr3")
	cmd.Flags().StringSlice(FlagDelete, nil, "addr,addr2,addr3")

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
