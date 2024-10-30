package config

import (
	"errors"
	"os"
)

/*
Creates a config file in the user's home directory and writes a default Config struct.

Returns a Config struct and an error.
*/
func makeConfigFile() (Config, error) {
	path, err := getConfigFilePath()
	if err != nil {
		return Config{}, err
	}
	_, err = os.Create(path)
	if err != nil {
		return Config{}, errors.New("couldn't create config file")
	}
	cfg := Config{
		AccessToken:  "",
		RefreshToken: "",
	}
	writeToConfigFile(cfg)
	return cfg, nil
}
