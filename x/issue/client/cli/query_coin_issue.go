package cli

import (
	"context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/viper"

	//	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/konstellation/konstellation/x/issue/types"
	"github.com/spf13/cobra"
)

func CmdListCoinIssue() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "Query issue list",
		Long:  "Query all or one of the account issue list, the limit default is 30",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			if _, err := sdk.AccAddressFromBech32(viper.GetString(flagOwner)); err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryIssuesRequest{
				Pagination: pageReq,
				Owner:      viper.GetString(flagOwner),
			}

			res, err := queryClient.Issues(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddPaginationFlagsToCmd(cmd, cmd.Use)
	flags.AddQueryFlagsToCmd(cmd)
	cmd.Flags().String(flagOwner, "", "Token owner address")

	return cmd
}

func CmdListAllCoinIssue() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-all",
		Short: "Query all issues",
		Long:  "Query all or one of the account issue list, the limit default is 30",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllIssuesRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.IssuesAll(context.Background(), params)
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

			params := &types.QueryIssueRequest{
				Denom: args[0],
			}

			res, err := queryClient.Issue(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
