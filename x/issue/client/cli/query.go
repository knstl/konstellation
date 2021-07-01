package cli

import (
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/codec"

	"github.com/konstellation/kn-sdk/x/issue/client/cli/query"
)

// GetQueryCmd returns the transaction commands for this module
func GetQueryCmd(cdc *codec.Codec) *cobra.Command {
	return query.GetQueryCmd(cdc)
}
