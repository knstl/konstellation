package tx

import (
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"

	"github.com/konstellation/kn-sdk/x/issue/types"
)

// GetTxCmd returns the transaction commands for this module
func GetTxCmd(cdc *codec.Codec) *cobra.Command {
	txCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Issue transaction subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	for _, c := range client.PostCommands(
		getTxCmdCreate(cdc),
		getTxCmdTransfer(cdc),
		getTxCmdApprove(cdc),
		getTxCmdIncreaseAllowance(cdc),
		getTxCmdDecreaseAllowance(cdc),
		getTxCmdTransferFrom(cdc),
		getTxCmdMint(cdc),
		getTxCmdBurn(cdc),
		getTxCmdBurnFrom(cdc),
		getTxCmdTransferOwnership(cdc),
		getTxCmdDisableFeature(cdc),
		getTxCmdEnableFeature(cdc),
		getTxCmdFeatures(cdc),
		getTxCmdDescription(cdc),
		getTxCmdFreeze(cdc),
		getTxCmdUnfreeze(cdc),
	) {
		_ = c.MarkFlagRequired(client.FlagFrom)
		txCmd.AddCommand(c)
	}

	return txCmd
}
