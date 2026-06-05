package testutil

import (
	"encoding/json"
	"fmt"
	"os"
)

type TestKeys struct {
	Primary   string `json:"primary"`
	Secondary string `json:"secondary"`
}

func LoadTestKeys() (*TestKeys, error) {
	paths := []string{
		"config/sec-keys.json",
		"../../config/sec-keys.json",
		"../../../config/sec-keys.json",
	}

	var data []byte
	var err error
	for _, path := range paths {
		data, err = os.ReadFile(path)
		if err == nil {
			break
		}
	}
	if err != nil {
		return nil, fmt.Errorf("read config/sec-keys.json: %w", err)
	}

	var keys TestKeys
	if err := json.Unmarshal(data, &keys); err != nil {
		return nil, fmt.Errorf("parse config/sec-keys.json: %w", err)
	}

	return &keys, nil
}
