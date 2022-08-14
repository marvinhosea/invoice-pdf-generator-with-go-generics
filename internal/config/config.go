package config

import (
	"errors"
	"os"
)

type UniDocConfig struct {
	Key string
}

func GetUniDocCred() (*UniDocConfig, error) {
	key, ok := os.LookupEnv("UNIDOC_LICENSE_API_KEY")
	if !ok {
		return nil, errors.New("uni doc key not found")
	}

	return &UniDocConfig{Key: key}, nil
}
