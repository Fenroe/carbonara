package config

func (cfg *Config) CheckIfLoggedIn() bool {
	return cfg.RefreshToken != ""
}
