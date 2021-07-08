package query

import (
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	sdk "github.com/cosmos/cosmos-sdk/types"

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

			owner, err := sdk.AccAddressFromBech32(args[0])
			if err != nil {
				return err
			}
			spender, err := sdk.AccAddressFromBech32(args[1])
			if err != nil {
				return err
			}

			denom := args[2]

			ctx := cmd.Context()
			allowanceRequest := types.QueryAllowanceRequest{
				Owner:   types.AccAddress(owner),
				Spender: types.AccAddress(spender),
				Denom:   denom,
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
