package query

import (
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/konstellation/konstellation/x/issue/types"
)

// getQueryCmdFreezes implements the query issue command.
func getQueryCmdFreezes() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "freezes [denom]",
		Args:  cobra.ExactArgs(1),
		Short: "Query freezes",
		Long:  "Query the amount of tokens that an owner allowed to all spender",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			denom := args[0]

			ctx := cmd.Context()
			freezesRequest := types.QueryFreezesRequest{
				Denom: denom,
			}
			res, err := queryClient.QueryFreezes(ctx, &freezesRequest)
			if err != nil {
				return err
			}
			return clientCtx.PrintProto(res)
		},
	}

	return cmd
}
