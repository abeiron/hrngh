package main

import (
	"fmt"

	"github.com/abeiron/hrngh/config"
	"github.com/bwmarrin/discordgo"
)

// Token contains the app token
var Token string

func main() {
	cfg := config.NewAppConfig()
	Token = cfg.AppToken()

	s, err := discordgo.New("Bot " + Token)
	if err != nil {
		fmt.Println("error creating Discord session: ", err)
		return
	}

	defer s.Close()
}
