package main

import (
	"blogAggregator/internal/config"
	"fmt"
	"os"
)

type state struct {
	cfg *config.Config
}

func main() {
	configData, err := config.Read()
	if err != nil {
		fmt.Println(err)
	}

	userArgs := os.Args[1:]
	if len(userArgs) < 2 {
		fmt.Println("not enough arguments!")
		os.Exit(1)
	}
	// fmt.Println(userArgs)

	userCmds := command{
		name:      userArgs[0],
		arguments: userArgs[1:],
	}

	programState := state{
		cfg: &configData,
	}

	commands := commands{
		commandsMap: map[string]func(*state, command) error{
			"login": handlerLogin,
		},
	}

	commands.run(&programState, userCmds)

	os.Exit(0)

}
