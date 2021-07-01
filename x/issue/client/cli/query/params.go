package query

import (
	"github.com/konstellation/kn-sdk/x/issue/query"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"

	"github.com/konstellation/kn-sdk/x/issue/types"
)

// getQueryCmdParams implements the query issue command.
func getQueryCmdParams(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "params",
		Short: "Query params",
		Long:  "Query issue module params",
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			// Query the issues
			res, _, err := cliCtx.QueryWithData(query.PathParams(), nil)
			if err != nil {
				return err
			}

			var params types.Params
			cdc.MustUnmarshalJSON(res, &params)
			return cliCtx.PrintOutput(params)
		},
	}

	return cmd
}
