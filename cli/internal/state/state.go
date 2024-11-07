package state

import (
	"net/http"

	"github.com/Fenroe/carbonara/cli/internal/config"
)

type State struct {
	Config config.Config
	APIURL string
	Client *http.Client
}
