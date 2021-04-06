package cli

import (
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/konstellation/konstellation/x/oracle/types"
)

func GetQueryExchangeRateCmd(cdc *codec.AminoCodec) *cobra.Command {
	txCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Querying exchange rate subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	txCmd.AddCommand(
		GetCmdQueryExchangeRate(cdc),
	)

	return txCmd
}

/*
TODO: need to update protobuf way
func GetCmdQueryExchangeRate(cdc *codec.AminoCodec) *cobra.Command {
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

			params := &types.QueryParamsRequest{}
			res, err := queryClient.Params(cmd.Context(), params)

			if err != nil {
				return err
			}

			return clientCtx.PrintProto(&res.Params)
		},
	}
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
*/

func GetCmdQueryExchangeRate(cdc *codec.AminoCodec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "exchange-rate",
		Short: "Query exchange rate",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			res, _, err := clientCtx.QueryWithData(types.QueryExchangeRate, nil)
			if err != nil {
				return err
			}
			var out sdk.Coin
			cdc.MustUnmarshalJSON(res, &out)
			return clientCtx.PrintObjectLegacy(out)
		},
	}
	return cmd
}
