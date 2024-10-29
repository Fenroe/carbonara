package config

import "github.com/Fenroe/carbonarapi/internal/database"

type Config struct {
	Greeting  string
	DB        *database.Queries
	JWTSecret string
}
