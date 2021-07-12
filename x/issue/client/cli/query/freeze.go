package query

import (
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"

	"github.com/konstellation/konstellation/x/issue/types"
)

// getQueryCmdAllowance implements the query issue command.
func getQueryCmdFreeze() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "freeze [denom] [holder]",
		Args:  cobra.ExactArgs(2),
		Short: "Query freeze",
		Long:  "Query freeze that an owner allowed to a spender",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			ctx := cmd.Context()
			freezeRequest := types.QueryFreezeRequest{
				Denom:  args[0],
				Holder: args[1],
			}
			res, err := queryClient.QueryFreeze(ctx, &freezeRequest)
			if err != nil {
				return err
			}
			return clientCtx.PrintProto(res)
		},
	}

	return cmd
}
