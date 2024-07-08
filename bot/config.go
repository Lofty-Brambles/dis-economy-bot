package bot

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	Token string `json:"token"`
}

func LoadConfig(path string) (*Config, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read config: invalid config file-path")
	}

	defer file.Close()

	var config Config
	if err := json.NewDecoder(file).Decode(&config); err != nil {
		return nil, err
	}

	return &config, nil
}
