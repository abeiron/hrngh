package main

import (
	"fmt"

	"github.com/abeiron/hrngh/config"
	"github.com/bwmarrin/discordgo"
)

func main() {
	var cfg = config.AppConfig()
	var token = cfg.Discord().Token()

	var client, err = discordgo.New("Bot " + token)
	if err != nil {
		fmt.Printf("error creating Discord session: %s", err)
		panic(err)
	}

	defer s.Close()

	var err = cfg.SaveConfig()
	if err != nil {
		fmt.Printf("error saving configuration: %d", err)
		panic(err)
	}
}
