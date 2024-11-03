package config

func (cfg *Config) SetCredentials(accessToken, refreshToken string) error {
	cfg.AccessToken = accessToken
	cfg.RefreshToken = refreshToken
	err := writeToConfigFile(*cfg)
	return err
}
