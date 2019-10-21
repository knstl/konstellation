package utils

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/tendermint/tendermint/libs/common"
)

func WriteFile(name string, dir string, contents []byte) error {
	writePath := filepath.Join(dir)
	file := filepath.Join(writePath, name)

	err := common.EnsureDir(writePath, 0700)
	if err != nil {
		return err
	}

	err = common.WriteFile(file, contents, 0600)
	if err != nil {
		return err
	}

	return nil
}

func ReadFile(name string) ([]byte, error) {
	file, err := os.Open(name)
	defer file.Close()
	if err != nil {
		return nil, err
	}

	return ioutil.ReadAll(file)
}
