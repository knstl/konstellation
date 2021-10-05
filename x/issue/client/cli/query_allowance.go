package cli

import (
	"context"
	sdk "github.com/cosmos/cosmos-sdk/types"

	//	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/konstellation/konstellation/x/issue/types"
	"github.com/spf13/cobra"
)

func CmdAllowances() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "allowances [owner] [denom]",
		Short: "Query allowances",
		Long:  "Query the amount of tokens that an owner allowed to all spender",
		Args:  cobra.RangeArgs(1, 2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			if _, err := sdk.AccAddressFromBech32(args[0]); err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllowanceRequest{
				Owner: args[0],
				Denom: args[1],
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

func CmdAllowance() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "allowance [owner] [spender] [denom]",
		Short: "shows a Allowance",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			if _, err := sdk.AccAddressFromBech32(args[0]); err != nil {
				return err
			}
			if _, err := sdk.AccAddressFromBech32(args[1]); err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllowanceRequest{
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
