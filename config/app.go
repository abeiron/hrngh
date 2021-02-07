package config

import (
	"fmt"
	"os"

	"github.com/pelletier/go-toml"
)

// App -
//
// Represents the customizable end-user configuration for the application.
type App struct {
	discord Discord
	tag string
	ver string
}

// AppConfig -
//
// Creates a new instance of the default app configuration.
func AppConfig() App {
	return App{
		DiscordConfig(),
		"am a bot",
		"0.1.0",
	}
}

// AppTag -
//
// Returns the application tagline.
func (app App) AppTag() string {
	return app.tag
}

// AppVer -
//
// Returns the application version.
func (app App) AppVer() string {
	return app.ver
}

// Discord -
//
// Returns the configuration of the Discord bot.
func (app *App) Discord() *Discord {
	return &app.discord
}

// SaveConfig -
//
// Saves the configuration to a file.
func (app App) SaveConfig() error {
	str, err := toml.Marshal(app)

	if err != nil {
		fmt.Printf("error marshalling application configuration: %d", err)
		return err
	}

	f, err := os.Create("config.toml")

	if err != nil {
		fmt.Printf("error creating configuration: %d", err)
		return err
	}

	defer f.Close()



	_, err = f.Write(str)
	if err != nil {
		fmt.Printf("error writing to file: %d", err)
		return err
	}

	f.Sync()



	return nil
}
