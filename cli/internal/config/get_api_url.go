package config

import (
	"os"
)

func (cfg *Config) GetAPIURL() string {
	devURL := os.Getenv("DEV_API_URL")
	if devURL != "" {
		return devURL
	}
	return "https://carbonarapi.fenmain.com"
}
