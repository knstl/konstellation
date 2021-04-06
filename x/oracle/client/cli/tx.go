package cli

import (
	"errors"
	"github.com/spf13/cobra"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	//	"github.com/cosmos/cosmos-sdk/types/msgservice"

	"github.com/konstellation/konstellation/x/oracle/types"
)

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
	)

	return txCmd
}

func NewMsgSetExchangeRateCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "set-exchange-rate [allowed-address] [denom] [amount]",
		Short: "Set exchange rate",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			allowedAddress, denom, amount := args[0], args[1], args[2]
			if allowedAddress == "" {
				return errors.New("invalid address")
			}
			if denom == "" {
				return errors.New("invalid denom name")
			}
			if amount == "" {
				return errors.New("invalid amount")
			}
			amountInt, err := strconv.Atoi(amount)
			if err != nil {
				return errors.New("invalid amount")
			}

			rate := sdk.NewCoin("Darc", sdk.NewInt(int64(amountInt)))
			msg := types.NewMsgSetExchangeRate(&rate, allowedAddress)
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
		Use:   "del-exchange-rate [allowed-address] [denom]",
		Short: "Delete exchange rate",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			allowedAddress, denom := args[0], args[1]
			if allowedAddress == "" {
				return errors.New("invalid address")
			}
			if denom == "" {
				return errors.New("invalid denom name")
			}

			msg := types.NewMsgDeleteExchangeRate(denom, allowedAddress)
			svcMsgClientConn := &ServiceMsgClientConn{}
			msgClient := types.NewMsgClient(svcMsgClientConn)
			_, err = msgClient.DeleteExchangeRate(cmd.Context(), msg)
			if err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), svcMsgClientConn.GetMsgs()...)
		},
	}
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
