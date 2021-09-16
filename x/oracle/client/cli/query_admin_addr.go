package cli

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/konstellation/konstellation/x/oracle/types"
	"github.com/spf13/cobra"
)

func CmdListAdminAddr() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-admin-addr",
		Short: "list all AdminAddr",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllAdminAddrRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.AdminAddrAll(context.Background(), params)
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

/*
func CmdShowAdminAddr() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-admin-addr [id]",
		Short: "shows a AdminAddr",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			id, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			params := &types.QueryGetAdminAddrRequest{
				Id: id,
			}

			res, err := queryClient.AdminAddr(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
*/
