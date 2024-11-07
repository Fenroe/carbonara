package commands

import (
	"fmt"

	"github.com/Fenroe/carbonara/cli/internal/state"
)

type Command struct {
	Name string
	Args []string
}

type Commands struct {
	Handlers map[string]func(*state.State, Command) error
}

func (c *Commands) Register(name string, f func(*state.State, Command) error) {
	c.Handlers[name] = f
}

func (c *Commands) Run(s *state.State, cmd Command) error {
	if _, ok := c.Handlers[cmd.Name]; !ok {
		return fmt.Errorf("no command called '%v' was found", cmd.Name)
	}
	return c.Handlers[cmd.Name](s, cmd)
}
