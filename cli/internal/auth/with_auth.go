package auth

import (
	"errors"
	"net/http"

	"github.com/Fenroe/carbonara/cli/internal/state"
)

func WithAuth(s *state.State, handler func(string) (*http.Response, error)) (*http.Response, error) {
	// Initial request with the current access token
	res, err := handler(s.Config.AccessToken)
	if err != nil {
		return nil, err
	}

	// Check if token refresh is needed
	if res.StatusCode == http.StatusUnauthorized {
		// Refresh the access token
		err = refreshAccessToken(s)
		if err != nil {
			return nil, err // Return if unable to refresh
		}

		// Retry the request with the new access token
		res, err = handler(s.Config.AccessToken)
		if err != nil {
			return nil, err
		}

		// Check if the second attempt is also unauthorized
		if res.StatusCode == http.StatusUnauthorized {
			return nil, errors.New("unauthorized: access token is invalid after refresh attempt")
		}
	}

	return res, nil
}
