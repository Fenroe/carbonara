package config

func (cfg *Config) ClearCredentials() {
	cfg.SetCredentials("", "")
}
