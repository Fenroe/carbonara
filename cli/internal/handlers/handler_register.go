package handlers

import (
	"errors"
	"fmt"

	"github.com/Fenroe/carbonara/cli/internal/commands"
	"github.com/Fenroe/carbonara/cli/internal/state"
	"github.com/Fenroe/carbonara/cli/internal/util"
)

func HandlerRegister(s *state.State, _ commands.Command) error {
	type request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	url := fmt.Sprintf("%s/api/users", s.APIURL)
	email, err := util.ReadEmail("Enter your email")
	if err != nil {
		return err
	}
	password, err := util.ReadPassword("Choose a password")
	if err != nil {
		return err
	}
	confirmPassword, err := util.ReadPassword("Confirm your password")
	if err != nil {
		return errors.New("couldn't confirm password")
	}
	if password != confirmPassword {
		return errors.New("passwords don't match")
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
	fmt.Println("Account created successfully")
	return nil
}
