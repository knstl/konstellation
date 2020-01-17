package query

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/konstellation/konstellation/x/issue/types"
	"github.com/spf13/cobra"
)

func pathQueryIssuesAll() string {
	return fmt.Sprintf("%s/%s/%s", types.Custom, types.QuerierRoute, types.QueryIssuesAll)
}

func getIssuesListAll(cliCtx context.CLIContext) ([]byte, int64, error) {
	return cliCtx.QueryWithData(pathQueryIssuesAll(), nil)
}

// getQueryCmdIssuesAll implements the query issue command.
func getQueryCmdIssuesAll(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-all",
		Short: "Query all issues ",
		Long:  "Query all or one of the account issue list, the limit default is 30",
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			// Query the issues
			res, _, err := getIssuesListAll(cliCtx)
			if err != nil {
				return err
			}

			var issues types.CoinIssues
			cdc.MustUnmarshalJSON(res, &issues)
			return cliCtx.PrintOutput(issues)
		},
	}

	return cmd
}
