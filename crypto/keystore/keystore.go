package keystore

import (
	"encoding/json"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/google/uuid"
)

const (
	version = 3
	scryptN = 8192
	scryptP = 1
)

type keyStore struct {
	PublicKey string              `json:"public_key"`
	Address   string              `json:"address"`
	Crypto    keystore.CryptoJSON `json:"crypto"`
	Id        string              `json:"id"`
	Name      string              `json:"name"`
	Version   int                 `json:"version"`
}

func NewKeyStoreV3(pk cryptotypes.PrivKey, name string, encryptPassword []byte) ([]byte, error) {
	pubkey, err := sdk.Bech32ifyPubKey(sdk.Bech32PubKeyTypeAccPub, pk.PubKey())
	if err != nil {
		return nil, err
	}
	address := sdk.AccAddress(pk.PubKey().Address()).String()
	cryptoStruct, err := keystore.EncryptDataV3(pk.Bytes(), encryptPassword, scryptN, scryptP)
	if err != nil {
		return nil, err
	}

	return json.Marshal(keyStore{
		pubkey,
		address,
		cryptoStruct,
		uuid.New().String(),
		name,
		version,
	})
}
