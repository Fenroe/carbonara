package config

import "github.com/Fenroe/carbonarapi/internal/database"

type Config struct {
	Greeting string
	Queries  *database.Queries
}
