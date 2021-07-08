package query

import (
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"

	"github.com/konstellation/konstellation/x/issue/types"
)

// getQueryCmdIssuesAll implements the query issue command.
func getQueryCmdIssuesAll() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-all",
		Short: "Query all issues",
		Long:  "Query all or one of the account issue list, the limit default is 30",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			ctx := cmd.Context()
			allIssuesRequest := types.QueryAllIssuesRequest{}
			res, err := queryClient.QueryAllIssues(ctx, &allIssuesRequest)
			if err != nil {
				return err
			}
			return clientCtx.PrintProto(res)
		},
	}

	return cmd
}
