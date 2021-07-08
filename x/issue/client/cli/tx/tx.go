package tx

import (
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"

	"github.com/konstellation/konstellation/x/issue/types"
)

// GetTxCmd returns the transaction commands for this module
func GetTxCmd() *cobra.Command {
	txCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Issue transaction subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	txCmd.AddCommand(
		getTxCmdCreate(),
		getTxCmdTransfer(),
		getTxCmdApprove(),
		getTxCmdIncreaseAllowance(),
		getTxCmdDecreaseAllowance(),
		getTxCmdTransferFrom(),
		getTxCmdMint(),
		getTxCmdBurn(),
		getTxCmdBurnFrom(),
		getTxCmdTransferOwnership(),
		getTxCmdDisableFeature(),
		getTxCmdEnableFeature(),
		getTxCmdFeatures(),
		getTxCmdDescription(),
		getTxCmdFreeze(),
		getTxCmdUnfreeze(),
	)

	return txCmd
}
