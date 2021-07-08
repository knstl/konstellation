package query

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/konstellation/konstellation/x/issue/types"
	"github.com/spf13/cobra"
)

// getQueryCmdParams implements the query issue command.
func getQueryCmdParams() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "params",
		Short: "Query params",
		Long:  "Query issue module params",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			ctx := cmd.Context()
			paramsRequest := types.QueryParamsRequest{}
			res, err := queryClient.QueryParams(ctx, &paramsRequest)
			if err != nil {
				return err
			}
			return clientCtx.PrintProto(res)
		},
	}
	return cmd
}
