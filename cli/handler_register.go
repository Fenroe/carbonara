package main

import (
	"errors"
	"fmt"
)

func handlerRegister(s *state, _ command) error {
	type request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	url := fmt.Sprintf("%s/api/users", s.apiurl)
	email, err := readEmail("Enter your email")
	if err != nil {
		return err
	}
	password, err := readPassword("Choose a password")
	if err != nil {
		return err
	}
	confirmPassword, err := readPassword("Confirm your password")
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
	res, err := doPostRequest(s, body, url)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	fmt.Println("Account created successfully")
	return nil
}
