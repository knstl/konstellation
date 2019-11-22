package keybase

import (
	"fmt"

	clkeys "github.com/cosmos/cosmos-sdk/client/keys"
	"github.com/cosmos/cosmos-sdk/crypto/keys"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// GenerateSaveCoinKey returns the address of a public key, along with the secret
// phrase to recover the private key.

func SaveCoinKey(clientRoot, keyName, keyPass, keyMnemonic string, overwrite bool) (sdk.AccAddress, string, error) {
	// get the keystore from the client
	keybase, err := clkeys.NewKeyBaseFromDir(clientRoot)
	if err != nil {
		return []byte{}, "", err
	}

	// ensure no overwrite
	if !overwrite {
		_, err := keybase.Get(keyName)
		if err == nil {
			return []byte{}, "", fmt.Errorf(
				"key already exists, overwrite is disabled (clientRoot: %s)", clientRoot)
		}
	}

	var info keys.Info
	if keyMnemonic == "" {
		// generate a private key, with recovery phrase
		info, keyMnemonic, err = keybase.CreateMnemonic(keyName, keys.English, keyPass, keys.Secp256k1)
	} else {
		// account 0 "Account number for HD derivation"
		// index 0 "Address index number for HD derivation"
		info, err = keybase.CreateAccount(keyName, keyMnemonic, "", keyPass, 0, 0)
	}

	if err != nil {
		return []byte{}, "", err
	}

	return sdk.AccAddress(info.GetPubKey().Address()), keyMnemonic, nil
}
