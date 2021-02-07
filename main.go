package main

import (
	"fmt"

	"github.com/abeiron/hrngh/config"
	"github.com/bwmarrin/discordgo"
)

// Token contains the app token
var Token string

func main() {
	cfg := config.AppConfig()
	Token = cfg.Discord().Token()

	s, err := discordgo.New("Bot " + Token)
	if err != nil {
		fmt.Printf("error creating Discord session: %s", err)
		return
	}

	defer s.Close()

	err := cfg.SaveConfig()
	if err != nil {
		fmt.Printf("error saving configuration: %d", err)
		return
	} else {
		
	}
}
