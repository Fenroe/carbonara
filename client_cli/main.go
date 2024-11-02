package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Fenroe/carbonara/internal/config"
	"github.com/joho/godotenv"
	"golang.design/x/clipboard"
)

type state struct {
	config config.Config
	apiurl string
	client *http.Client
}

type command struct {
	name string
	args []string
}

type commands struct {
	handlers map[string]func(*state, command) error
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.handlers[name] = f
}

func (c *commands) run(s *state, cmd command) error {
	if _, ok := c.handlers[cmd.name]; !ok {
		return fmt.Errorf("no command called '%v' was found", cmd.name)
	}
	return c.handlers[cmd.name](s, cmd)
}

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

func doPostRequest[T any](body T, url string) (*http.Response, error) {
	jsonData, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if res.StatusCode < 200 || res.StatusCode >= 300 {
		return res, fmt.Errorf("unexpected HTTP status: %s", res.Status)
	}
	return res, err
}

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
	res, err := doPostRequest(body, url)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	fmt.Println("Account created successfully")
	return nil
}

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
	res, err := doPostRequest(body, url)
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

func handlerLogout(s *state, _ command) error {
	s.config.SetCredentials("", "")
	return nil
}

func handlerSend(s *state, _ command) error {
	type request struct {
		Content string `json:"content"`
	}

	url := fmt.Sprintf("%s/api/clips", s.apiurl)
	err := clipboard.Init()
	if err != nil {
		return err
	}
	content := string(clipboard.Read(clipboard.FmtText))
	body := request{
		Content: content,
	}
	jsonData, err := json.Marshal(body)
	if err != nil {
		return errors.New("couldn't marshal json")
	}

	// Define a function to send the request
	sendRequest := func(token string) (*http.Response, error) {
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
		if err != nil {
			return nil, errors.New("couldn't create HTTP request")
		}
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", token))

		client := &http.Client{}
		return client.Do(req)
	}

	// First attempt with the current access token
	res, err := sendRequest(s.config.AccessToken)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	// Check if the token needs to be refreshed
	if res.StatusCode == http.StatusUnauthorized {
		// Attempt to refresh the access token
		err = refreshAccessToken(s)
		if err != nil {
			return err // Return if unable to refresh
		}

		// Retry with the new access token
		res, err = sendRequest(s.config.AccessToken)
		if err != nil {
			return err
		}
		defer res.Body.Close()

		// Check if the second attempt is also unauthorized
		if res.StatusCode == http.StatusUnauthorized {
			return errors.New("unauthorized: access token is invalid after refresh attempt")
		}
	}

	fmt.Println("Your clipboard data has been sent")
	return nil
}

func handlerSync(_ *state, _ command) error {
	return nil
}

func main() {
	godotenv.Load()
	cfg, err := config.New()
	if err != nil {
		log.Fatal(err)
	}
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	appState := state{
		config: cfg,
		apiurl: cfg.GetAPIURL(),
		client: &http.Client{},
	}
	handlers := make(map[string]func(*state, command) error)
	appCommands := commands{
		handlers: handlers,
	}
	appCommands.register("register", handlerRegister)
	appCommands.register("login", handlerLogin)
	appCommands.register("send", handlerSend)
	appCommands.register("sync", handlerSync)
	appCommands.register("logout", handlerLogout)
	cliArgs := os.Args
	if len(cliArgs) < 2 {
		fmt.Println("type 'help' for a list of commands")
		os.Exit(0)
	} else {
		cliArgs = cliArgs[1:]
	}
	userCommand := command{
		name: cliArgs[0],
		args: cliArgs[1:],
	}
	err = appCommands.run(&appState, userCommand)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
}
