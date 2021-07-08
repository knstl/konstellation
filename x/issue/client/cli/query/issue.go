package query

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/konstellation/konstellation/x/issue/types"
	"github.com/spf13/cobra"
)

// getCmdQueryIssue implements the query issue command.
func getQueryCmdIssue() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "find",
		Short: "Query issue by denom",
		Long:  "Query issue by denom",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			denom := args[0]

			ctx := cmd.Context()
			issueRequest := types.QueryIssueRequest{
				Denom: denom,
			}
			res, err := queryClient.QueryIssue(ctx, &issueRequest)
			if err != nil {
				return err
			}
			return clientCtx.PrintProto(res)
		},
	}

	return cmd
}
