package keys

import (
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client/keys"
)

// Commands registers a sub-tree of commands to interact with local private key storage.
func Commands() *cobra.Command {
	cmd := keys.Commands()
	cmd.AddCommand(exportKeyStoreCommand())

	return cmd
}
