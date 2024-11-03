package main

import "fmt"

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
