package config

import (
	"fmt"
	"os"
	"path/filepath"
)

func getConfigFilePath() (string, error) {
	configDir, err := os.UserConfigDir() // Returns ~/.config or equivalent
	if err != nil {
		fmt.Println("Failed to get config directory:", err)
		return "", err
	}
	// Define the path for your app's config file
	path := filepath.Join(configDir, "carbonara", "config.json")
	return path, nil
}
