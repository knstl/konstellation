package cli

import (
	"github.com/spf13/cobra"

	"github.com/konstellation/konstellation/x/issue/client/cli/tx"
)

// GetTxCmd returns the transaction commands for this module
func GetTxCmd() *cobra.Command {
	return tx.GetTxCmd()
}
