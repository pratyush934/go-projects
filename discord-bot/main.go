package main

import (
	"fmt"
	"github.com/pratyush934/go-projects/discord-bot/bot"
	"github.com/pratyush934/go-projects/discord-bot/config"
)

func main() {
	err := config.ReadConfig()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	bot.Start()

	<-make(chan struct{})

	return
}
