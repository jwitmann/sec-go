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
	data, err := os.ReadFile("../../config/sec-keys.json")
	if err != nil {
		return nil, fmt.Errorf("read config/sec-keys.json: %w", err)
	}

	var keys TestKeys
	if err := json.Unmarshal(data, &keys); err != nil {
		return nil, fmt.Errorf("parse config/sec-keys.json: %w", err)
	}

	return &keys, nil
}
