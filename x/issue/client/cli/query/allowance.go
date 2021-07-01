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
func getQueryCmdAllowance(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "allowance [owner] [spender] [denom]",
		Args:  cobra.ExactArgs(3),
		Short: "Query allowance",
		Long:  "Query the amount of tokens that an owner allowed to a spender",
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			owner, err := sdk.AccAddressFromBech32(args[0])
			if err != nil {
				return err
			}
			spender, err := sdk.AccAddressFromBech32(args[1])
			if err != nil {
				return err
			}

			denom := args[2]

			res, _, err := cliCtx.QueryWithData(query.PathQueryIssueAllowance(owner, spender, denom), nil)
			if err != nil {
				return err
			}

			var approval types.Allowance
			cdc.MustUnmarshalJSON(res, &approval)

			return cliCtx.PrintOutput(approval)
		},
	}

	return cmd
}
