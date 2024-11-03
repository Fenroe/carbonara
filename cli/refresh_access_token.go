package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func refreshAccessToken(s *state) error {
	type response struct {
		AccessToken string `json:"access_token"`
	}

	url := fmt.Sprintf("%s/api/refresh", s.apiurl)
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return err
	}
	token := fmt.Sprintf("Bearer %v", s.config.RefreshToken)
	req.Header.Set("Authorization", token)
	res, err := s.client.Do(req)
	if err != nil {
		return err
	}
	if res.StatusCode == http.StatusUnauthorized {
		// Access Token has expired, user must re-authenticate
		s.config.SetCredentials("", "")
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
	s.config.SetCredentials(data.AccessToken, s.config.RefreshToken)
	return nil
}
