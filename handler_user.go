package main

import (
	"blogAggregator/internal/database"
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.arguments) == 0 {
		return errors.New("empty command arguments! 'login' expects a single argument to be passed!")
	}

	if cmd.arguments[0] == "unknown" {
		return errors.New("rejected login for user with name 'unknown'")
	}

	_, err := s.db.GetUser(context.Background(), cmd.arguments[0])
	if err != nil {
		// fmt.Println("this user doesn't exists in the database, hence login blocked!\n error:", err)
		return err
	}

	err = s.cfg.SetUser(cmd.arguments[0])
	if err != nil {
		return err
	}
	fmt.Println("User has been set!")

	return nil
}

func handlerRegister(s *state, cmd command) error {
	if len(cmd.arguments) == 0 {
		return errors.New("empty command arguments! 'register' expects a single argument to be passed!")
	}

	// Checking if user already exists...
	userInDB, err := s.db.GetUser(context.Background(), cmd.arguments[0])
	if err != nil {
		// Inserting the user into the PostgreSQL DB
		userEntry := database.CreateUserParams{
			ID:        uuid.New(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Name:      cmd.arguments[0],
		}

		userInDB, err = s.db.CreateUser(context.Background(), userEntry)
		if err != nil {
			return err
		}

		fmt.Println("user was created!\nUser Details:")
		fmt.Printf("ID: %s\nCreated on: %v\nUpdated on: %v\nName: %v\n", userInDB.ID.String(), userInDB.CreatedAt, userInDB.UpdatedAt, userInDB.Name)

		// Setting the user
		err = s.cfg.SetUser(cmd.arguments[0])
		if err != nil {
			return err
		}
	} else {
		// fmt.Println("a user with this username already exists in the database.\ndetails:")
		// fmt.Printf("ID: %s\nCreated on: %v\nUpdated on: %v\nName: %v\n", userInDB.ID.String(), userInDB.CreatedAt, userInDB.UpdatedAt, userInDB.Name)
		return errors.New("a user with this username already exists in the database!")
	}
	return nil
}

func handlerReset(s *state, cmd command) error {
	if len(cmd.arguments) > 0 {
		return errors.New("the reset command shouldn't have any arguements!")
	}
	err := s.db.ResetUsers(context.Background())
	if err != nil {
		return err
	}
	return nil
}
