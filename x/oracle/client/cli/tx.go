package cli

import (
	"errors"
	"github.com/spf13/cobra"
	"strconv"

	//"github.com/cosmos/cosmos-sdk/client/flags"
	//"github.com/cosmos/cosmos-sdk/types/msgservice"
	//"github.com/cosmos/cosmos-sdk/client/tx"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
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

/*
TODO: need to update protobuf way
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
			// No ServiceMsgClient struct in msgservice
			svcMsgClientConn := &msgservice.ServiceMsgClientConn{}
			// No idea to generate NewMsgClient function from protobuf
			msgClient := types.NewMsgClient(svcMsgClientConn)
			_, err = msgClient.VerifyInvariant(cmd.Context(), msg)
			if err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), svcMsgClientConn.GetMsgs()...)
		},
	}
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
*/

func NewMsgSetExchangeRateCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "set-exchange-rate [allowed-address] [denom] [amount]",
		Short: "Set exchange rate",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
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

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			rate := sdk.NewCoin("Darc", sdk.NewInt(int64(amountInt)))
			msg := types.NewMsgSetExchangeRate(&rate, allowedAddress)
			err = msg.ValidateBasic()
			if err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}
	return cmd
}

func NewMsgDeleteExchangeRateCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "del-exchange-rate [allowed-address] [denom]",
		Short: "Delete exchange rate",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			allowedAddress, denom := args[0], args[1]
			if allowedAddress == "" {
				return errors.New("invalid address")
			}
			if denom == "" {
				return errors.New("invalid denom name")
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			msg := types.NewMsgDeleteExchangeRate(denom, allowedAddress)
			err = msg.ValidateBasic()
			if err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}
	return cmd
}
