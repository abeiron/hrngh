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

	defer client.Close()

	var err0 = cfg.SaveConfig()
	if err0 != nil {
		fmt.Printf("error saving configuration: %d", err0)
		panic(err0)
	}
}
