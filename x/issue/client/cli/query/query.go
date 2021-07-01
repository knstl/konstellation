package query

import (
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"

	"github.com/konstellation/kn-sdk/x/issue/types"
)

// GetQueryCmd returns the transaction commands for this module
func GetQueryCmd(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Querying commands for the issue module",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(
		getQueryCmdIssue(cdc),
		getQueryCmdIssues(cdc),
		getQueryCmdIssuesAll(cdc),
		getQueryCmdAllowance(cdc),
		getQueryCmdAllowances(cdc),
		getQueryCmdParams(cdc),
		getQueryCmdFreeze(cdc),
		getQueryCmdFreezes(cdc),
	)

	return cmd
}
