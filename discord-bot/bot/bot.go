package bot

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/pratyush934/go-projects/discord-bot/config"
	"log"
)

// BotId /* This is the ID for the Bot, it is very important to trace*/
var BotId string

/* When we create a session, it is important to store it somewhere so to make it accessible*/
var goBot *discordgo.Session //

func Start() {

	goBot, err := discordgo.New("Bot " + config.Token)

	if err != nil {
		log.Printf("there is an error while Starting, %v", err)
		return
	}

	user, err := goBot.User("@me")

	if err != nil {
		fmt.Println(err.Error())
	}

	BotId = user.ID

	goBot.AddHandler(messageHandler)

	err = goBot.Open()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("The Bot is running")

}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {

	if m.Author.ID == BotId {
		return
	}

	if m.Content == "ping" {
		_, err := s.ChannelMessageSend(m.ChannelID, "pong")

		if err != nil {
			fmt.Println(err.Error())
			return
		}
	}

}
