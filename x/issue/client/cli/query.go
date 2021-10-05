package cli

import (
	"fmt"
	// "strings"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	// "github.com/cosmos/cosmos-sdk/client/flags"
	// sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/konstellation/konstellation/x/issue/types"
)

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd(queryRoute string) *cobra.Command {
	// Group issue queries under a subcommand
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("Querying commands for the %s module", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	// this line is used by starport scaffolding # 1

	cmd.AddCommand(CmdListIssues())
	cmd.AddCommand(CmdListAllIssues())
	cmd.AddCommand(CmdShowIssue())

	cmd.AddCommand(CmdAllowances())
	cmd.AddCommand(CmdAllowance())

	//cmd.AddCommand(CmdListCoinIssueDenoms())
	//cmd.AddCommand(CmdShowCoinIssueDenoms())

	//cmd.AddCommand(CmdListCoinIssueList())
	//cmd.AddCommand(CmdShowCoinIssueList())

	cmd.AddCommand(CmdListIssuesParams())
	//cmd.AddCommand(CmdShowIssuesParams())

	cmd.AddCommand(CmdListIssueParams())
	cmd.AddCommand(CmdShowIssueParams())

	//cmd.AddCommand(CmdListIssueFeatures())
	//cmd.AddCommand(CmdShowIssueFeatures())

	cmd.AddCommand(CmdListParams())
	//cmd.AddCommand(CmdShowParams())

	//cmd.AddCommand(CmdListIssues())
	//cmd.AddCommand(CmdShowIssues())

	//cmd.AddCommand(CmdListCoinIssueDenom())
	//cmd.AddCommand(CmdShowCoinIssueDenom())

	cmd.AddCommand(CmdListFreeze())
	cmd.AddCommand(CmdShowFreeze())

	//cmd.AddCommand(CmdListCoins())
	//cmd.AddCommand(CmdShowCoins())

	cmd.AddCommand(CmdListAllowanceList())
	cmd.AddCommand(CmdShowAllowanceList())

	cmd.AddCommand(CmdListAddress())
	cmd.AddCommand(CmdShowAddress())

	return cmd
}
