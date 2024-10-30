package config

import (
	"encoding/json"
	"errors"
	"os"
)

/*
Returns a new Config struct based on the config file in the user's home directory.

If no such file exists, instead creates a config file and returns a default Config struct.
*/
func New() (Config, error) {
	path, err := getConfigFilePath()
	if err != nil {
		return Config{}, err
	}
	_, err = os.Stat(path)
	if err != nil {
		return makeConfigFile()
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
