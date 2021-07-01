package cli

import (
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/codec"

	"github.com/konstellation/kn-sdk/x/issue/client/cli/tx"
)

// GetTxCmd returns the transaction commands for this module
func GetTxCmd(cdc *codec.Codec) *cobra.Command {
	return tx.GetTxCmd(cdc)
}
