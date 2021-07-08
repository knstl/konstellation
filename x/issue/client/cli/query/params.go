package query

import (
	"github.com/konstellation/konstellation/x/issue/query"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"

	"github.com/konstellation/konstellation/x/issue/types"
)

// getQueryCmdParams implements the query issue command.
func getQueryCmdParams(cdc *codec.LegacyAmino) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "params",
		Short: "Query params",
		Long:  "Query issue module params",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := client.Context{}
			cliCtx := ctx.WithLegacyAmino(cdc)

			// Query the issues
			res, _, err := cliCtx.QueryWithData(query.PathParams(), nil)
			if err != nil {
				return err
			}

			var params types.Params
			cdc.MustUnmarshalJSON(res, &params)
			return cliCtx.PrintObjectLegacy(params)
		},
	}

	return cmd
}
