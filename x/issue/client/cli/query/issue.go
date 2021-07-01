package query

import (
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/konstellation/kn-sdk/x/issue/query"
	"github.com/konstellation/kn-sdk/x/issue/types"
	"github.com/spf13/cobra"
)

// getCmdQueryIssue implements the query issue command.
func getQueryCmdIssue(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "find",
		Short: "Query issue by denom",
		Long:  "Query issue by denom",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			denom := args[0]

			// Query the issues
			res, _, err := cliCtx.QueryWithData(query.PathQueryIssue(denom), nil)
			if err != nil {
				return err
			}

			var issue types.CoinIssue
			cdc.MustUnmarshalJSON(res, &issue)
			return cliCtx.PrintOutput(&issue)
		},
	}

	return cmd
}
