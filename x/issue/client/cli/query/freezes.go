package query

import (
	"github.com/konstellation/kn-sdk/x/issue/query"
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/konstellation/kn-sdk/x/issue/types"
)

// getQueryCmdFreezes implements the query issue command.
func getQueryCmdFreezes(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "freezes [denom]",
		Args:  cobra.ExactArgs(1),
		Short: "Query freezes",
		Long:  "Query the amount of tokens that an owner allowed to all spender",
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			denom := args[0]

			res, _, err := cliCtx.QueryWithData(query.PathQueryIssueFreezes(denom), nil)
			if err != nil {
				return err
			}

			var freezes types.AddressFreezes
			cdc.MustUnmarshalJSON(res, &freezes)

			return cliCtx.PrintOutput(&freezes)
		},
	}

	return cmd
}
