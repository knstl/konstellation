package cli

import (
	"errors"
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

func NewExchangeRateCmd() *cobra.Command {
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
		Use:   "set-exchange-rate [allowed-address] [denom] [rate]",
		Short: "Set exchange rate",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			allowedAddress, denom, rateStr := args[0], args[1], args[2]
			if allowedAddress == "" {
				return errors.New("invalid address")
			}
			if denom == "" {
				return errors.New("invalid denom name")
			}
			if rateStr == "" {
				return errors.New("invalid rate")
			}
			rate, err := strconv.ParseFloat(rateStr, 32)
			if err != nil {
				return errors.New("invalid rate")
			}

			exchangeRate := types.ExchangeRate{
				Denom: "udarc",
				Rate:  uint64(rate * float64(RateUnit)),
			}
			msg := types.NewMsgSetExchangeRate(&exchangeRate, allowedAddress)
			svcMsgClientConn := &ServiceMsgClientConn{}
			msgClient := types.NewMsgClient(svcMsgClientConn)
			_, err = msgClient.SetExchangeRate(cmd.Context(), &msg)
			if err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), svcMsgClientConn.GetMsgs()...)
		},
	}
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func NewMsgDeleteExchangeRateCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "del-exchange-rate [sender]",
		Short: "Delete exchange rate",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			sender := args[0]
			if sender == "" {
				return errors.New("invalid sender")
			}

			msg := types.NewMsgDeleteExchangeRate(sender)
			svcMsgClientConn := &ServiceMsgClientConn{}
			msgClient := types.NewMsgClient(svcMsgClientConn)
			_, err = msgClient.DeleteExchangeRate(cmd.Context(), &msg)
			if err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), svcMsgClientConn.GetMsgs()...)
		},
	}
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func NewMsgSetAdminAddrCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "set-admin-addr [sender] [address-to-add] [address-to-delete]",
		Short: "Set Admin Address",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			sender, add, del := args[0], args[1], args[2]
			if sender == "" {
				return errors.New("invalid sender")
			}
			var addressesToAdd []string
			if add != "" {
				addressesToAdd = strings.Split(add, ",")
			}
			var addressesToDelete []string
			if del != "" {
				addressesToDelete = strings.Split(del, ",")
			}
			msg := types.NewMsgSetAdminAddr(sender, addressesToAdd, addressesToDelete)
			svcMsgClientConn := &ServiceMsgClientConn{}
			msgClient := types.NewMsgClient(svcMsgClientConn)
			_, err = msgClient.SetAdminAddr(cmd.Context(), &msg)
			if err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), svcMsgClientConn.GetMsgs()...)
		},
	}
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
