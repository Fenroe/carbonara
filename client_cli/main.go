package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Fenroe/carbonara/internal/config"
	"github.com/joho/godotenv"
)

type state struct {
	config config.Config
	apiurl string
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

func handlerRegister(s *state, cmd command) error {
	type request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	url := fmt.Sprintf("%s/api/users", s.apiurl)
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your email")
	email, err := reader.ReadString('\n')
	if err != nil {
		return errors.New("couldn't get email")
	}
	fmt.Print("Choose a password")
	password, err := reader.ReadString('\n')
	if err != nil {
		return errors.New("couldn't get password")
	}
	fmt.Print("Confirm your password")
	confirmPassword, err := reader.ReadString('\n')
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
	jsonData, err := json.Marshal(body)
	if err != nil {
		return errors.New("couldn't marshal json")
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return errors.New("couldn't create HTTP request")
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	fmt.Print("Account created successfully")
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
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your email")
	email, err := reader.ReadString('\n')
	if err != nil {
		return errors.New("couldn't get email")
	}
	password, err := reader.ReadString('\n')
	if err != nil {
		return errors.New("couldn't get password")
	}
	body := request{
		Email:    email,
		Password: password,
	}
	jsonData, err := json.Marshal(body)
	if err != nil {
		return errors.New("couldn't marshal json")
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return errors.New("couldn't create HTTP request")
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	res, err := client.Do(req)
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
	}
	handlers := make(map[string]func(*state, command) error)
	appCommands := commands{
		handlers: handlers,
	}
	appCommands.register("register", handlerRegister)
	appCommands.register("login", handlerLogin)
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

/*
	err = clipboard.Init()
	text := string(clipboard.Read(clipboard.FmtText))
	if text != "" {
		fmt.Println(text)
		return
	}
*/
