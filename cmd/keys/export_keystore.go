package keys

import (
	"bufio"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client/input"
	"github.com/cosmos/cosmos-sdk/client/keys"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/konstellation/konstellation/crypto/keystore"
)

func exportKeyStoreCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "export-keystore <name>",
		Short: "Export account to key store",
		Long:  `Export account to key store in encrypted format.`,
		Args:  cobra.ExactArgs(1),
		RunE:  runExportKeyStoreCmd,
	}
	return cmd
}

func runExportKeyStoreCmd(cmd *cobra.Command, args []string) error {
	kb, err := keys.NewKeyBaseFromHomeFlag()
	if err != nil {
		return err
	}

	buf := bufio.NewReader(cmd.InOrStdin())
	decryptPassword, err := input.GetPassword("Enter passphrase to decrypt your key:", buf)
	if err != nil {
		return err
	}
	encryptPassword, err := input.GetPassword("Enter passphrase to encrypt the exported key:", buf)
	if err != nil {
		return err
	}

	ac, err := kb.ExportPrivateKeyObject(args[0], decryptPassword)
	if err != nil {
		return err
	}

	encryptedKeyStore, err := keystore.NewKeyStoreV3(
		sdk.AccAddress(ac.PubKey().Address()).String(),
		ac.Bytes(),
		[]byte(encryptPassword),
	)

	cmd.Println(string(encryptedKeyStore))
	return err
}
