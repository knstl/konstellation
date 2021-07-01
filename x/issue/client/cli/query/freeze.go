package query

import (
	"github.com/konstellation/kn-sdk/x/issue/query"
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/konstellation/kn-sdk/x/issue/types"
)

// getQueryCmdAllowance implements the query issue command.
func getQueryCmdFreeze(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "freeze [denom] [holder]",
		Args:  cobra.ExactArgs(2),
		Short: "Query freeze",
		Long:  "Query freeze that an owner allowed to a spender",
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

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

			return cliCtx.PrintOutput(&freeze)
		},
	}

	return cmd
}
