package cli

import (
	"context"
	//	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/konstellation/konstellation/x/issue/types"
	"github.com/spf13/cobra"
)

func CmdListCoinIssue() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-coin-issue",
		Short: "list all CoinIssue",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllCoinIssueRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.CoinIssueAll(context.Background(), params)
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

func CmdShowCoinIssue() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-coin-issue [denom]",
		Short: "shows a CoinIssue",
		Long:  "Query issue by denom",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryGetCoinIssueRequest{
				Denom: args[0],
			}

			res, err := queryClient.CoinIssue(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
