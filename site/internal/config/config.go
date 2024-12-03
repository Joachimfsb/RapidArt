package util

import (
	"encoding/json"
	"os"
	"path/filepath"
	"rapidart/internal/consts"
)

var Config config

type config struct {
	Server struct {
		Port string `json:"port"`
		Host string `json:"host"`
	} `json:"server"`
	Database struct {
		Url      string `json:"url"`
		Db       string `json:"db"`
		Username string `json:"user"`
		Password string `json:"pass"`
	} `json:"database"`
}

func InitializeConfig() error {
	// Read config
	bytes, err := os.ReadFile(filepath.Join(consts.CONFIG_DIR, "config.json"))
	if err != nil {
		return err
	}
	// Decode
	err2 := json.Unmarshal(bytes, &Config)
	if err2 != nil {
		return err2
	}

	return nil
}
