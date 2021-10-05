package cli

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/konstellation/konstellation/x/issue/types"
)

const (
	flagDecimals           = "decimals"
	flagBurnOwnerDisabled  = "burn-owner"
	flagBurnHolderDisabled = "burn-holder"
	flagBurnFromDisabled   = "burn-from"
	flagFreezeDisabled     = "freeze"
	flagMintDisabled       = "mint"
	flagDescription        = "description"
)

// GetTxCmd returns the transaction commands for this module
func GetTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("%s transactions subcommands", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	// this line is used by starport scaffolding # 1
	cmd.AddCommand(CmdIssueCreate())
	cmd.AddCommand(CmdFeatures())
	cmd.AddCommand(CmdEnableFeature())
	cmd.AddCommand(CmdDisableFeature())
	cmd.AddCommand(CmdDescription())
	cmd.AddCommand(CmdFreeze())
	cmd.AddCommand(CmdUnfreeze())
	cmd.AddCommand(CmdMint())
	cmd.AddCommand(CmdBurn())
	cmd.AddCommand(CmdBurnFrom())
	cmd.AddCommand(CmdApprove())
	cmd.AddCommand(CmdIncreaseAllowance())
	cmd.AddCommand(CmdDecreaseAllowance())
	cmd.AddCommand(CmdTransferOwnership())
	cmd.AddCommand(CmdTransferFrom())
	cmd.AddCommand(CmdTransfer())

	return cmd
}
