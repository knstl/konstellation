package query

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/cosmos/cosmos-sdk/client"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/konstellation/konstellation/x/issue/types"
)

const (
	flagOwner = "owner"
	flagLimit = "limit"
)

// getCmdQueryIssues implements the query issue command.
func getQueryCmdIssues() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "Query issue list",
		Long:  "Query all or one of the account issue list, the limit default is 30",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			address, err := sdk.AccAddressFromBech32(viper.GetString(flagOwner))
			if err != nil {
				return err
			}
			qp := types.NewIssuesParams(
				address.String(),
				int32(viper.GetInt(flagLimit)),
			)

			ctx := cmd.Context()
			issuesRequest := types.QueryIssuesRequest{Params: &qp}
			res, err := queryClient.QueryIssues(ctx, &issuesRequest)
			if err != nil {
				return err
			}
			return clientCtx.PrintProto(res)
		},
	}

	cmd.Flags().String(flagOwner, "", "Token owner address")
	cmd.Flags().Int32(flagLimit, 30, "Query number of issue results per page returned")
	return cmd
}
