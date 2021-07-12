package query

import (
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"

	"github.com/konstellation/konstellation/x/issue/types"
)

// getQueryCmdAllowance implements the query issue command.
func getQueryCmdAllowance() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "allowance [owner] [spender] [denom]",
		Args:  cobra.ExactArgs(3),
		Short: "Query allowance",
		Long:  "Query the amount of tokens that an owner allowed to a spender",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			ctx := cmd.Context()
			allowanceRequest := types.QueryAllowanceRequest{
				Owner:   args[0],
				Spender: args[1],
				Denom:   args[2],
			}
			res, err := queryClient.QueryAllowance(ctx, &allowanceRequest)
			if err != nil {
				return err
			}
			return clientCtx.PrintProto(res)
		},
	}

	return cmd
}
