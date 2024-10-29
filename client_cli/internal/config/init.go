package config

import (
	"encoding/json"
	"errors"
	"os"
)

func Init() (Config, error) {
	path, err := getConfigFilePath()
	if err != nil {
		return Config{}, err
	}
	_, err = os.Stat(path)
	if err != nil {
		return Config{}, errors.New("couldn't find config file")
	}
	file, err := os.Open(path)
	if err != nil {
		return Config{}, errors.New("couldn't open config file")
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	cfg := Config{}
	err = decoder.Decode(&cfg)
	if err != nil {
		return Config{}, err
	}
	return cfg, nil
}
