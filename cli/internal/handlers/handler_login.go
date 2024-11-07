package handlers

import (
	"encoding/json"
	"fmt"

	"github.com/Fenroe/carbonara/cli/internal/commands"
	"github.com/Fenroe/carbonara/cli/internal/state"
	"github.com/Fenroe/carbonara/cli/internal/util"
)

func HandlerLogin(s *state.State, _ commands.Command) error {
	type request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	type response struct {
		AccessToken  string `json:"access_token"`
		RefreshToken string `json:"refresh_token"`
	}

	url := fmt.Sprintf("%s/api/login", s.APIURL)
	email, err := util.ReadEmail("Enter your email")
	if err != nil {
		return err
	}
	password, err := util.ReadPassword("Choose a password")
	if err != nil {
		return err
	}
	body := request{
		Email:    email,
		Password: password,
	}
	res, err := util.DoPostRequest(s, body, url)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	var userData response
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&userData)
	if err != nil {
		return err
	}
	fmt.Println("Log in successful")
	return s.Config.SetCredentials(userData.AccessToken, userData.RefreshToken)
}
