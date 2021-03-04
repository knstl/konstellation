package keys

import (
	//"bufio"
	//"github.com/cosmos/cosmos-sdk/client"
	"github.com/spf13/cobra"
	//"github.com/cosmos/cosmos-sdk/client/input"
	//"github.com/cosmos/cosmos-sdk/client/keys"
	//"github.com/konstellation/konstellation/crypto/keystore"
)

func exportKeyStoreCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "export-keystore <name>",
		Short: "Export account to key store",
		Long:  `Export account to key store in encrypted format.`,
		Args:  cobra.ExactArgs(1),
		RunE:  runExportKeyStoreCmd,
	}
}

func runExportKeyStoreCmd(cmd *cobra.Command, args []string) error {
	//kb, err := keys.NewKeyBaseFromHomeFlag()
	//if err != nil {
	//	return err
	//}

	//buf := bufio.NewReader(cmd.InOrStdin())
	//clientCtx, err := client.GetClientQueryContext(cmd)
	//if err != nil {
	//	return err
	//}
	//decryptPassword, err := input.GetPassword("Enter passphrase to decrypt your key:", buf)
	//if err != nil {
	//	return err
	//}
	//encryptPassword, err := input.GetPassword("Enter passphrase to encrypt the exported key:", buf)
	//if err != nil {
	//	return err
	//}
	//
	//ac, err := kb.ExportPrivateKeyObject(args[0], decryptPassword)
	//if err != nil {
	//	return err
	//}
	//
	//clientCtx.Keyring.ExportPrivateKeyObject()
	//
	//clientCtx.Keyring.
	//
	//encryptedKeyStore, err := keystore.NewKeyStoreV3(ac, args[0], []byte(encryptPassword))
	//
	//cmd.Println(string(encryptedKeyStore))
	//return err

	return nil
}
