package tx

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/konstellation/konstellation/x/issue/types"
	"github.com/spf13/cobra"
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
		getTxCmdMintTo(cdc),
		getTxCmdBurn(cdc),
		getTxCmdBurnFrom(cdc),
	) {
		_ = c.MarkFlagRequired(client.FlagFrom)
		txCmd.AddCommand(c)
	}

	return txCmd
}
