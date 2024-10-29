package config

type Config struct {
	JWT          string `json:"jwt"`
	RefreshToken string `json:"refresh_token"`
	APIURL       string `json:"api_url"`
}
