package config

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

func Read() Config {

	cfg := Config{}

	path, err := getConfigFilePath()
	if err != nil {
		fmt.Println("Error: ", err)
		return cfg
	}

	// Open our jsonFile
	jsonFile, err := os.Open(path)
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
		return cfg
	}
	fmt.Println("Successfully Opened ", path)
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	byteValue, err := io.ReadAll(jsonFile)

	if err != nil {
		fmt.Print(err)
	}
	err = json.Unmarshal(byteValue, &cfg)
	if err != nil {
		fmt.Print(err)
	}

	return cfg
}

func SetUser(name string, cfg Config) {
	cfg.CurrentUserName = name
	write(cfg)
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
