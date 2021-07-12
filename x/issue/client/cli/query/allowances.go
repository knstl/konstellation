package query

import (
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"

	"github.com/konstellation/konstellation/x/issue/types"
)

// getQueryCmdAllowances implements the query issue command.
func getQueryCmdAllowances() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "allowances [owner] [denom]",
		Args:  cobra.ExactArgs(2),
		Short: "Query allowances",
		Long:  "Query the amount of tokens that an owner allowed to all spender",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			ctx := cmd.Context()
			allowancesRequest := types.QueryAllowancesRequest{
				Owner: args[0],
				Denom: args[1],
			}
			res, err := queryClient.QueryAllowances(ctx, &allowancesRequest)
			if err != nil {
				return err
			}
			return clientCtx.PrintProto(res)
		},
	}

	return cmd
}
