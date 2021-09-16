package cli

import (
	"context"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/konstellation/konstellation/x/issue/types"
	"github.com/spf13/cobra"
)

func CmdListAllowanceList() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-allowance-list",
		Short: "list all AllowanceList",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllAllowanceListRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.AllowanceListAll(context.Background(), params)
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

func CmdShowAllowanceList() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-allowance-list [id]",
		Short: "shows a AllowanceList",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			id, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			params := &types.QueryGetAllowanceListRequest{
				Id: id,
			}

			res, err := queryClient.AllowanceList(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
