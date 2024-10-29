package config

import (
	"errors"
	"fmt"
	"os"
)

func getConfigFilePath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", errors.New("couldn't find home directory")
	}
	path := fmt.Sprintf("%v.carbonaraconfig.json", homeDir)
	return path, nil
}
