package cli

import (
	"context"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/konstellation/konstellation/x/issue/types"
	"github.com/spf13/cobra"
)

const (
	flagOwner = "owner"
	flagLimit = "limit"
)

func CmdListIssuesParams() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-issues-params --owner [owner-address] --limit [limit]",
		Short: "list all IssuesParams",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			owner, _ := cmd.Flags().GetString(flagOwner)
			limit, _ := cmd.Flags().GetString(flagLimit)
			limitNumber, err := strconv.ParseUint(limit, 10, 64)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			qp := types.NewIssuesParams(owner, int32(limitNumber))

			params := &types.QueryAllIssuesParamsRequest{
				Params: &qp,
			}

			res, err := queryClient.IssuesParamsAll(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

/*
func CmdShowIssuesParams() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-issues-params [id]",
		Short: "shows a IssuesParams",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			id, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			params := &types.QueryGetIssuesParamsRequest{
				Id: id,
			}

			res, err := queryClient.IssuesParams(context.Background(), params)
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
