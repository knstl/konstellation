package keybase

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/tendermint/crypto/bcrypt"
	tmcrypto "github.com/tendermint/tendermint/crypto"
)

func HashPassword(pass string) ([]byte, error) {
	saltBytes := tmcrypto.CRandBytes(16)
	passwordHash, err := bcrypt.GenerateFromPassword(saltBytes, []byte(pass), 2)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return nil, err
	}

	return passwordHash, nil
}

func SaveHashedPassword(dir string, passwordHash []byte) error {
	return ioutil.WriteFile(dir+"/keyhash", passwordHash, 0555)
}
