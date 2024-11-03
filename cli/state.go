package main

import (
	"net/http"

	"github.com/Fenroe/carbonara/cli/internal/config"
)

type state struct {
	config config.Config
	apiurl string
	client *http.Client
}
