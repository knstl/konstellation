package query

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/konstellation/konstellation/x/issue/query"
	"github.com/konstellation/konstellation/x/issue/types"
	"github.com/spf13/cobra"
)

// getCmdQueryIssue implements the query issue command.
func getQueryCmdIssue(cdc *codec.LegacyAmino) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "find",
		Short: "Query issue by denom",
		Long:  "Query issue by denom",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := client.Context{}
			cliCtx := ctx.WithLegacyAmino(cdc)

			denom := args[0]

			// Query the issues
			res, _, err := cliCtx.QueryWithData(query.PathQueryIssue(denom), nil)
			if err != nil {
				return err
			}

			var issue types.CoinIssue
			cdc.MustUnmarshalJSON(res, &issue)
			return cliCtx.PrintObjectLegacy(&issue)
		},
	}

	return cmd
}
