package query

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/konstellation/konstellation/x/issue/types"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func pathQueryIssueAllowance(owner sdk.AccAddress, spender sdk.AccAddress, issueID string) string {
	return fmt.Sprintf("%s/%s/%s/%s/%s/%s", types.Custom, types.QuerierRoute, types.QueryAllowance, issueID, owner.String(), spender.String())
}

func getIssueAllowance(cliCtx context.CLIContext, owner sdk.AccAddress, spender sdk.AccAddress, issueID string) ([]byte, int64, error) {
	return cliCtx.QueryWithData(pathQueryIssueAllowance(owner, spender, issueID), nil)
}

// GetQueryCmdAllowance implements the query issue command.
func GetQueryCmdAllowance(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "allowance [owner] [spender] [issue-id]",
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

			issueID := args[2]
			if err := types.CheckIssueId(issueID); err != nil {
				return errors.Errorf(err.Error())
			}

			res, _, err := getIssueAllowance(cliCtx, owner, spender, issueID)
			if err != nil {
				return err
			}

			var approval types.Allowance
			cdc.MustUnmarshalJSON(res, &approval)

			return cliCtx.PrintOutput(approval)
		},
	}

	cmd.Flags().String(flagAddress, "", "Token owner address")
	cmd.Flags().String(flagStartIssueId, "", "Start issueId of issues")
	cmd.Flags().Int32(flagLimit, 30, "Query number of issue results per page returned")
	return cmd
}
