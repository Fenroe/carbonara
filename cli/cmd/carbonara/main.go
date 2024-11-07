package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Fenroe/carbonara/cli/internal/commands"
	"github.com/Fenroe/carbonara/cli/internal/config"
	"github.com/Fenroe/carbonara/cli/internal/handlers"
	"github.com/Fenroe/carbonara/cli/internal/state"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	cfg, err := config.New()
	if err != nil {
		log.Fatal(err)
	}
	appState := state.State{
		Config: cfg,
		APIURL: cfg.GetAPIURL(),
		Client: &http.Client{},
	}
	hnds := make(map[string]func(*state.State, commands.Command) error)
	appCommands := commands.Commands{
		Handlers: hnds,
	}
	appCommands.Register("register", handlers.HandlerRegister)
	appCommands.Register("login", handlers.HandlerLogin)
	appCommands.Register("send", handlers.HandlerSend)
	appCommands.Register("sync", handlers.HandlerSync)
	appCommands.Register("logout", handlers.HandlerLogout)
	cliArgs := os.Args
	if len(cliArgs) < 2 {
		fmt.Println("type 'help' for a list of commands")
		os.Exit(0)
	} else {
		cliArgs = cliArgs[1:]
	}
	userCommand := commands.Command{
		Name: cliArgs[0],
		Args: cliArgs[1:],
	}
	err = appCommands.Run(&appState, userCommand)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
}
