package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Fenroe/carbonara/cli/internal/config"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	cfg, err := config.New()
	if err != nil {
		log.Fatal(err)
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
