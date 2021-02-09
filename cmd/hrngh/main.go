package main

import (
	"fmt"

	"github.com/abeiron/hrngh/internal/config"
)

func main() {
	var cfg = config.AppConfig()
	var token = cfg.Discord().Token()

	var err = cfg.SaveConfig()
	if err != nil {
		fmt.Printf("error saving configuration: %d", err)
		panic(err)
	}
}
