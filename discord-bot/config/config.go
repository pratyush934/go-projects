package config

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

var (
	Token     string
	BotPrefix string

	config *configStruct
)

type configStruct struct {
	Token     string `json:"Token"`
	BotPrefix string `json:"BotPrefix"`
}

func ReadConfig() error {
	fmt.Println("Reading the config file")

	file, err := os.ReadFile("./config.json")

	if err != nil {
		log.Printf("ther is an error in ReadConfig file %v", err.Error())
		return err
	}

	fmt.Println(string(file))

	err = json.Unmarshal(file, &config)

	if err != nil {
		log.Printf("There is an error %v", err.Error())
		return err
	}

	Token = config.Token
	BotPrefix = config.BotPrefix

	return nil
}
