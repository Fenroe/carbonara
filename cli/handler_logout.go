package main

import (
	"fmt"
	"net/http"
)

func handlerLogout(s *state, _ command) error {
	if !s.config.CheckIfLoggedIn() {
		fmt.Println("You aren't currently logged in.")
		return nil
	}
	url := fmt.Sprintf("%v/api/revoke", s.apiurl)
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", s.config.RefreshToken))
	res, err := s.client.Do(req)
	if err != nil {
		return err
	}
	if res.StatusCode < 200 || res.StatusCode >= 300 {
		return fmt.Errorf("unexpected HTTP status: %s", res.Status)
	}
	s.config.ClearCredentials()
	fmt.Println("Successfully logged out")
	return nil
}
