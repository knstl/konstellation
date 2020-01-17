package query

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/konstellation/konstellation/x/issue/types"
	"github.com/spf13/cobra"
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
		getQueryCmdIssues(cdc),
		getQueryCmdIssuesAll(cdc),
		getQueryCmdAllowance(cdc),
	)

	return cmd
}
