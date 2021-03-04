package utils

import (
	os2 "github.com/tendermint/tendermint/libs/os"
	"io/ioutil"
	"os"
	"path/filepath"
)

func WriteFile(name string, dir string, contents []byte) error {
	writePath := filepath.Join(dir)
	file := filepath.Join(writePath, name)

	err := os2.EnsureDir(writePath, 0700)
	if err != nil {
		return err
	}

	err = os2.WriteFile(file, contents, 0600)
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
