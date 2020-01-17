package cli

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/konstellation/konstellation/x/issue/client/cli/query"
	"github.com/spf13/cobra"
)

// GetQueryCmd returns the transaction commands for this module
func GetQueryCmd(cdc *codec.Codec) *cobra.Command {
	return query.GetQueryCmd(cdc)
}
