package config

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

const configFileName = ".gatorconfig.json"

type Config struct {
	DbURL           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func Read() (Config, error) {

	cfg := Config{}

	path, err := getConfigFilePath()
	if err != nil {
		fmt.Println("Error: ", err)
		return cfg, err
	}

	// Open our jsonFile
	jsonFile, err := os.Open(path)
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
		return cfg, err
	}
	fmt.Println("Successfully Opened ", path)
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	byteValue, err := io.ReadAll(jsonFile)

	if err != nil {
		fmt.Print(err)
		return cfg, err
	}
	err = json.Unmarshal(byteValue, &cfg)
	if err != nil {
		fmt.Print(err)
		return Config{}, err
	}

	return cfg, nil
}

func SetUser(name string, cfg Config) error {
	cfg.CurrentUserName = name
	err := write(cfg)
	if err != nil {
		return err
	}
	return nil
}

func getConfigFilePath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	path := home + "/" + configFileName
	if _, err := os.Stat(path); err != nil {
		return "", err
	}
	return path, nil
}

func write(cfg Config) error {

	json, err := json.Marshal(cfg)
	if err != nil {
		return err
	}

	path, err := getConfigFilePath()
	if err != nil {
		return err
	}

	err = os.WriteFile(path, json, 0666)
	if err != nil {
		return nil
	}

	return nil
}
