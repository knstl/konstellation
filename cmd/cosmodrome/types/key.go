package types

import (
	"errors"
)

type Key struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	Mnemonic string `json:"mnemonic"`
}

func (k *Key) GetName() string {
	return k.Name
}

func (k *Key) GetPassword() string {
	return k.Password
}

func (k *Key) GetMnemonic() string {
	return k.Mnemonic
}

type KeyStorage struct {
	Keys map[string]*Key `json:"keys"`
}

func (ks *KeyStorage) GetKey(address string) (*Key, error) {
	key, ex := ks.Keys[address]
	if !ex {
		return nil, errors.New("key does not exist")
	}

	return key, nil
}
