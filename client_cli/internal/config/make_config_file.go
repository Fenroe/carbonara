package config

import (
	"fmt"
	"os"
	"path/filepath"
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
	// Create the directory if it doesn't exist
	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		fmt.Println("Failed to create config directory:", err)
		return Config{}, err
	}
	// Now create or open the config file for writing
	file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("Failed to open or create config file:", err)
		return Config{}, err
	}
	defer file.Close()
	cfg := Config{
		AccessToken:  "",
		RefreshToken: "",
	}
	writeToConfigFile(cfg)
	return cfg, nil
}
