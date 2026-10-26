package main

import (
	"blogAggregator/internal/config"
	"fmt"
)

func main() {
	config, err := config.Read()
	if err != nil {
		fmt.Println(err)
	}
	// fmt.Println(config.CurrentUsername)
	config.SetUser("Shreyansh")
	fmt.Println(config.DB_Url)
	fmt.Println(config.CurrentUsername)
}
