package query

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/konstellation/konstellation/x/issue/types"
	"github.com/spf13/cobra"
)

func pathQueryIssueAllowance(owner sdk.AccAddress, spender sdk.AccAddress, denom string) string {
	return fmt.Sprintf("%s/%s/%s/%s/%s/%s", types.Custom, types.QuerierRoute, types.QueryAllowance, denom, owner.String(), spender.String())
}

func getIssueAllowance(cliCtx context.CLIContext, owner sdk.AccAddress, spender sdk.AccAddress, issueID string) ([]byte, int64, error) {
	return cliCtx.QueryWithData(pathQueryIssueAllowance(owner, spender, issueID), nil)
}

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

			res, _, err := getIssueAllowance(cliCtx, owner, spender, denom)
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
