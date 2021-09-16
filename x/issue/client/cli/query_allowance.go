package cli

import (
	"context"
	//	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/konstellation/konstellation/x/issue/types"
	"github.com/spf13/cobra"
)

func CmdListAllowance() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-allowance [owner] [denom]",
		Short: "list all Allowance",
		Long:  "Query the amount of tokens that an owner allowed to all spender",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllAllowanceRequest{
				Owner: args[0],
				Denom: args[1],
			}

			res, err := queryClient.AllowanceAll(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdShowAllowance() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-allowance [owner] [spender] [denom]",
		Short: "shows a Allowance",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryGetAllowanceRequest{
				Owner:   args[0],
				Spender: args[1],
				Denom:   args[2],
			}

			res, err := queryClient.Allowance(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
