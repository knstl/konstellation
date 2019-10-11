package keystore

import (
	"encoding/json"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/google/uuid"
)

const (
	version = 3
	scryptN = 8192
	scryptP = 1
)

type keyStore struct {
	Address string              `json:"address"`
	Crypto  keystore.CryptoJSON `json:"crypto"`
	Id      string              `json:"id"`
	Version int                 `json:"version"`
}

func NewKeyStoreV3(address string, data []byte, encryptPassword []byte) ([]byte, error) {
	cryptoStruct, err := keystore.EncryptDataV3(data, encryptPassword, scryptN, scryptP)
	if err != nil {
		return nil, err
	}

	return json.Marshal(keyStore{
		address,
		cryptoStruct,
		uuid.New().String(),
		version,
	})
}
