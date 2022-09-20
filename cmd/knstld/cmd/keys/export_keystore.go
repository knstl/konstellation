package keys

import (
	"bufio"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/input"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/crypto"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/konstellation/konstellation/crypto/keystore"
	"github.com/spf13/cobra"
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

	var blankcodec codec.Codec
	inBuf := bufio.NewReader(cmd.InOrStdin())
	keyringBackend, _ := cmd.Flags().GetString(flags.FlagKeyringBackend)
	clientCtx, err := client.GetClientQueryContext(cmd)
	if err != nil {
		return err
	}
	kb, err := keyring.New(sdk.KeyringServiceName(), keyringBackend, clientCtx.HomeDir, inBuf, blankcodec)
	if err != nil {
		return err
	}

	decryptPassword, err := input.GetPassword("Enter passphrase to decrypt your key:", inBuf)
	if err != nil {
		return err
	}
	encryptPassword, err := input.GetPassword("Enter passphrase to encrypt the exported key:", inBuf)
	if err != nil {
		return err
	}

	ac, err := kb.ExportPrivKeyArmor(args[0], decryptPassword)
	if err != nil {
		return err
	}
	pk, _, err := crypto.UnarmorDecryptPrivKey(ac, decryptPassword)
	if err != nil {
		return err
	}

	encryptedKeyStore, err := keystore.NewKeyStoreV3(pk, args[0], []byte(encryptPassword))
	if err != nil {
		return err
	}

	cmd.Println(string(encryptedKeyStore))

	return nil
}
