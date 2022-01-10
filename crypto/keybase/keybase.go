package keybase

import (
	hd "github.com/cosmos/cosmos-sdk/crypto/hd"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"io"
)

type expectedKey interface {
	GetName() string
	GetPassword() string
	GetMnemonic() string
}

// SaveCoinKey returns the address of a public key, along with the secret
// phrase to recover the private key.
func SaveCoinKey(dir, keyringBackend, algoStr string, key expectedKey, overwrite bool, testKeyring bool, inBuf io.Reader) (addr sdk.AccAddress, keyMnemonic string, err error) {
	// TODO appName?

	kr, err := keyring.New(sdk.KeyringServiceName(), keyringBackend, dir, inBuf)
	if err != nil {
		return []byte{}, "", err
	}

	// for development simplifying
	//if testKeyring {
	//	passwdHash, err := HashPassword(key.GetPassword())
	//	if err != nil {
	//		return nil, "", err
	//	}
	//
	//	if err := SaveHashedPassword(dir, passwdHash); err != nil {
	//		return nil, "", err
	//	}
	//}

	keyringAlgos, _ := kr.SupportedAlgorithms()
	algo, err := keyring.NewSigningAlgoFromString(algoStr, keyringAlgos)
	if err != nil {
		return nil, "", err
	}

	var info keyring.Info
	path := hd.CreateHDPath(118, 0, 0).String()
	if key.GetMnemonic() == "" {
		// generate a private key, with recovery phrase
		info, keyMnemonic, err = kr.NewMnemonic(key.GetName(), keyring.English, path, key.GetPassword(), algo)
	} else {
		// account 0 "Account number for HD derivation"
		// index 0 "Address index number for HD derivation"
		info, err = kr.NewAccount(key.GetName(), key.GetMnemonic(), "", path, algo)
		keyMnemonic = key.GetMnemonic()
	}

	if err != nil {
		return []byte{}, "", err
	}

	return sdk.AccAddress(info.GetPubKey().Address()), keyMnemonic, nil
}
