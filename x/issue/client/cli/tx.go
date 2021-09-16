package cli

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	// "github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/konstellation/konstellation/x/issue/types"
)

var (
	DefaultRelativePacketTimeoutTimestamp = uint64((time.Duration(10) * time.Minute).Nanoseconds())
)

const (
	flagPacketTimeoutTimestamp = "packet-timeout-timestamp"

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
	cmd.AddCommand(CmdUnfreeze())

	cmd.AddCommand(CmdBurnFrom())

	cmd.AddCommand(CmdBurn())

	cmd.AddCommand(CmdMint())

	cmd.AddCommand(CmdDecreaseAllowance())

	cmd.AddCommand(CmdIncreaseAllowance())

	cmd.AddCommand(CmdIncreaseAllowance())

	cmd.AddCommand(CmdApprove())

	cmd.AddCommand(CmdTransferOwnership())

	cmd.AddCommand(CmdTransferFrom())

	cmd.AddCommand(CmdTransfer())

	cmd.AddCommand(CmdFeatures())

	cmd.AddCommand(CmdEnableFeature())

	cmd.AddCommand(CmdDisableFeature())

	cmd.AddCommand(CmdDescription())

	cmd.AddCommand(CmdIssueCreate())

	return cmd
}
