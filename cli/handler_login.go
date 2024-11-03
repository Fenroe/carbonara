package main

import (
	"encoding/json"
	"fmt"
)

func handlerLogin(s *state, _ command) error {
	type request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	type response struct {
		AccessToken  string `json:"access_token"`
		RefreshToken string `json:"refresh_token"`
	}

	url := fmt.Sprintf("%s/api/login", s.apiurl)
	email, err := readEmail("Enter your email")
	if err != nil {
		return err
	}
	password, err := readPassword("Choose a password")
	if err != nil {
		return err
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
	var userData response
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&userData)
	if err != nil {
		return err
	}
	fmt.Println("Log in successful")
	return s.config.SetCredentials(userData.AccessToken, userData.RefreshToken)
}
