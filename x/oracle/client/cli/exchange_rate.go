package cli

import (
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/konstellation/konstellation/x/oracle/types"
)

// NewTxCmd returns a root CLI command handler for all x/crisis transaction commands.
func NewExchangeRateCmd() *cobra.Command {
	txCmd := &cobra.Command{
		Use:   types.ModuleName,
		Short: "Exchange Rate subcommands", DisableFlagParsing: true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	txCmd.AddCommand(NewMsgExchangeRateCmd())

	return txCmd
}

// NewMsgVerifyInvariantTxCmd returns a CLI command handler for creating a
// MsgVerifyInvariant transaction.
func NewMsgExchangeRateCmd() *cobra.Command {
	cmd := &cobra.Command{}
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
