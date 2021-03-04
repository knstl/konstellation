package utils

import (
	"encoding/json"
)

func ReadJson(name string, obj interface{}) error {
	bytes, err := ReadFile(name)
	if err != nil {
		return err
	}

	return json.Unmarshal(bytes, &obj)
}
