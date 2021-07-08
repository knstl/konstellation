package query

import (
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"

	"github.com/konstellation/konstellation/x/issue/types"
)

// GetQueryCmd returns the transaction commands for this module
func GetQueryCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Querying commands for the issue module",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(
		getQueryCmdIssue(),
		getQueryCmdIssues(),
		getQueryCmdIssuesAll(),
		getQueryCmdAllowance(),
		getQueryCmdAllowances(),
		getQueryCmdParams(),
		getQueryCmdFreeze(),
		getQueryCmdFreezes(),
	)

	return cmd
}
