package config

import (
	"encoding/json"
	"errors"
	"os"
)

type Config struct {
	DB_Url          string `json:"db_url"`
	CurrentUsername string `json:"current_user_name"`
}

func Read() (Config, error) {
	var c Config
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return Config{}, err
	}
	configPath := homeDir + "/.gatorconfig.json"
	configByteData, err := os.ReadFile(configPath)
	if err != nil {
		return Config{}, err
	}

	err = json.Unmarshal(configByteData, &c)
	if err != nil {
		return Config{}, err
	}

	return c, nil
}

func (c *Config) SetUser(username string) error {
	if username == "" {
		return errors.New("username cannot be empty")
	}

	homeDir, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	configPath := homeDir + "/.gatorconfig.json"

	c.CurrentUsername = username
	configByteData, err := json.Marshal(c)
	if err != nil {
		return err
	}

	err = os.WriteFile(configPath, configByteData, 0666)
	if err != nil {
		return err
	}
	return nil
}
