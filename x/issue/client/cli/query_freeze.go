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

func CmdListFreeze() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "freezes [holder]",
		Short: "Query freezes",
		Long:  "Query the amount of tokens that an owner allowed to all spender",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			if _, err := sdk.AccAddressFromBech32(args[0]); err != nil {
				return err
			}

			params := &types.QueryFreezesRequest{
				Holder: args[0],
			}

			res, err := queryClient.Freezes(context.Background(), params)
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

func CmdListAllFreeze() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "freezes-all [denom]",
		Short: "Query all freezes",
		Long:  "Query all freezed account for denom",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllFreezesRequest{
				Denom: args[0],
			}

			res, err := queryClient.FreezesAll(context.Background(), params)
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
		Use:   "freeze [denom] [holder]",
		Short: "Query freeze",
		Long:  "Query freeze that an owner allowed to a spender",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryFreezeRequest{
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
