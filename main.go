package main

import (
	"blogAggregator/internal/config"
	"blogAggregator/internal/database"
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

type state struct {
	db  *database.Queries
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

	db, err := sql.Open("postgres", configData.DB_Url)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	dbQueries := database.New(db)
	programState := state{
		db:  dbQueries,
		cfg: &configData,
	}

	commands := commands{
		commandsMap: map[string]func(*state, command) error{
			"login":    handlerLogin,
			"register": handlerRegister,
		},
	}

	err = commands.run(&programState, userCmds)
	if err != nil {
		fmt.Println("there has been an error:\n", err)
	}

	os.Exit(0)

}
