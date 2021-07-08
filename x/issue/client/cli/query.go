package cli

import (
	"github.com/spf13/cobra"

	"github.com/konstellation/konstellation/x/issue/client/cli/query"
)

// GetQueryCmd returns the transaction commands for this module
func GetQueryCmd() *cobra.Command {
	return query.GetQueryCmd()
}
