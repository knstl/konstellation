package keys

import (
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
)

// Commands registers a sub-tree of commands to interact with local private key storage.
func Commands(defaultNodeHome string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "keys2",
		Short: "Additional commands for keys",
		Long:  `Additional commands for keys for Knstld`,
	}
	cmd.AddCommand(exportKeyStoreCommand())

	cmd.PersistentFlags().String(flags.FlagKeyringBackend, flags.DefaultKeyringBackend, "Select keyring's backend (os|file|test)")

	return cmd
}
