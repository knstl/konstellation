package cli

import (
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"

	"github.com/konstellation/konstellation/x/oracle/types"
)

func GetQueryExchangeRateCmd() *cobra.Command {
	qCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Querying exchange rate subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	qCmd.AddCommand(
		GetCmdQueryExchangeRate(),
		GetCmdQueryExchangeRates(),
	)

	return qCmd
}

func GetCmdQueryExchangeRate() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "exchange-rate",
		Short: "Query exchange rate",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			// protobuf function and structs
			queryClient := types.NewQueryClient(clientCtx)

			pair := args[0]
			exchangeRate := &types.QueryExchangeRateRequest{
				Pair: pair,
			}

			res, err := queryClient.ExchangeRate(cmd.Context(), exchangeRate)

			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res.ExchangeRate)
		},
	}
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func GetCmdQueryExchangeRates() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "exchange-rates",
		Short: "Query all available exchange rates",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			// protobuf function and structs
			queryClient := types.NewQueryClient(clientCtx)

			exchangeRate := &types.QueryExchangeRateRequest{}
			res, err := queryClient.ExchangeRate(cmd.Context(), exchangeRate)

			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res.ExchangeRate)
		},
	}
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
