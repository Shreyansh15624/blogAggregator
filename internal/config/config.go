package config

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
)

type Config struct {
	DB_Url          string `json:"db_url"`
	CurrentUsername string `json:"current_user_name"`
}

// Global Declaration of the configuration file name as a 'const'
const configFileName = ".gatorconfig.json"

// Helper Functions
func getConfigFilePath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	configPath := filepath.Join(homeDir, configFileName)
	return configPath, nil
}

func write(cfg *Config) error {
	configPath, err := getConfigFilePath()
	if err != nil {
		return err
	}
	configByteData, err := json.Marshal(cfg)
	if err != nil {
		return err
	}

	err = os.WriteFile(configPath, configByteData, 0666)
	if err != nil {
		return err
	}
	return nil
}

// Config Methods
func (c *Config) SetUser(username string) error {
	if username == "" {
		return errors.New("username cannot be empty")
	}

	c.CurrentUsername = username

	err := write(c)
	if err != nil {
		return err
	}
	return nil
}

// Config Aid Functions
func Read() (Config, error) {
	var c Config
	configPath, err := getConfigFilePath()
	if err != nil {
		return Config{}, err
	}
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
