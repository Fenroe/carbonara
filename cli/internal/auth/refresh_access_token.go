package auth

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Fenroe/carbonara/cli/internal/state"
)

func refreshAccessToken(s *state.State) error {
	type response struct {
		AccessToken string `json:"access_token"`
	}

	url := fmt.Sprintf("%s/api/refresh", s.APIURL)
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return err
	}
	token := fmt.Sprintf("Bearer %v", s.Config.RefreshToken)
	req.Header.Set("Authorization", token)
	res, err := s.Client.Do(req)
	if err != nil {
		return err
	}
	if res.StatusCode == http.StatusUnauthorized {
		// Access Token has expired, user must re-authenticate
		s.Config.SetCredentials("", "")
		fmt.Println("Please log back in")
		return nil
	}
	defer res.Body.Close()
	var data response
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&data)
	if err != nil {
		return err
	}
	s.Config.SetCredentials(data.AccessToken, s.Config.RefreshToken)
	return nil
}
