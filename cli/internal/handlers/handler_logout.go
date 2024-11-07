package handlers

import (
	"fmt"
	"net/http"

	"github.com/Fenroe/carbonara/cli/internal/commands"
	"github.com/Fenroe/carbonara/cli/internal/state"
)

func HandlerLogout(s *state.State, _ commands.Command) error {
	if !s.Config.CheckIfLoggedIn() {
		fmt.Println("You aren't currently logged in.")
		return nil
	}
	url := fmt.Sprintf("%v/api/revoke", s.APIURL)
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", s.Config.RefreshToken))
	res, err := s.Client.Do(req)
	if err != nil {
		return err
	}
	if res.StatusCode < 200 || res.StatusCode >= 300 {
		return fmt.Errorf("unexpected HTTP status: %s", res.Status)
	}
	s.Config.ClearCredentials()
	fmt.Println("Successfully logged out")
	return nil
}
