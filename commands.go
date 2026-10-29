package main

import "fmt"

type command struct {
	name      string
	arguments []string
}

type commands struct {
	commandsMap map[string]func(*state, command) error
}

func (c *commands) run(s *state, cmd command) error {
	handler, ok := c.commandsMap[cmd.name]
	if !ok {
		return nil
	}

	err := handler(s, cmd)
	if err != nil {
		return err
	}
	return nil
}

func (c *commands) register(name string, f func(s *state, cmd command) error) {
	if _, ok := c.commandsMap[name]; ok {
		fmt.Println("command already exists!")
	} else {
		c.commandsMap[name] = f
	}
}
