package cli

import (
	"context"
	//	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/konstellation/konstellation/x/issue/types"
	"github.com/spf13/cobra"
)

func CmdListFreeze() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-freeze [denom]",
		Short: "list all Freeze",
		Long:  "Query the amount of tokens that an owner allowed to all spender",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllFreezeRequest{
				Denom: args[0],
			}

			res, err := queryClient.FreezeAll(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddPaginationFlagsToCmd(cmd, cmd.Use)
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdShowFreeze() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-freeze [denom] [holder]",
		Short: "shows a Freeze",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryGetFreezeRequest{
				Denom:  args[0],
				Holder: args[1],
			}

			res, err := queryClient.Freeze(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
