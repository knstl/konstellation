package query

import (
	"github.com/konstellation/konstellation/x/issue/query"
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/konstellation/konstellation/x/issue/types"
)

// getQueryCmdAllowance implements the query issue command.
func getQueryCmdFreeze(cdc *codec.LegacyAmino) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "freeze [denom] [holder]",
		Args:  cobra.ExactArgs(2),
		Short: "Query freeze",
		Long:  "Query freeze that an owner allowed to a spender",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := client.Context{}
			cliCtx := ctx.WithLegacyAmino(cdc)

			holder, err := sdk.AccAddressFromBech32(args[1])
			if err != nil {
				return err
			}

			denom := args[0]

			res, _, err := cliCtx.QueryWithData(query.PathQueryIssueFreeze(denom, holder), nil)
			if err != nil {
				return err
			}

			var freeze types.Freeze
			cdc.MustUnmarshalJSON(res, &freeze)

			return cliCtx.PrintObjectLegacy(&freeze)
		},
	}

	return cmd
}
