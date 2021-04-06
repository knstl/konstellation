package cli

import (
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"

	"github.com/konstellation/konstellation/x/oracle/types"
)

func GetQueryExchangeRateCmd() *cobra.Command {
	txCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Querying exchange rate subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	txCmd.AddCommand(
		GetCmdQueryExchangeRate(),
	)

	return txCmd
}

func GetCmdQueryExchangeRate() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "exchange-rate",
		Short: "Query exchange rate",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			// protobuf function and structs
			queryClient := types.NewQueryClient(clientCtx)

			exchangeRate := &types.QueryExchangeRateRequest{}
			res, err := queryClient.GetExchangeRate(cmd.Context(), exchangeRate)

			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res.ExchangeRate)
		},
	}
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
