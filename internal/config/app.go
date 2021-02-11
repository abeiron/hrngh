package config

import (
	"errors"
	"fmt"
	"os"

	"github.com/hashicorp/hcl/v2"
)

// App represents the customizable end-user configuration for the application.
type App struct {
	db Database		`hcl:"database"`
	discord Discord `hcl:"discord"`
	tag string 		`hcl:"tag"`
	ver string 		`hcl:"ver"`
}

// AppConfig creates a new instance of the default app configuration.
func AppConfig() *App {
	return App{
		DiscordConfig(),
		"am a bot",
		"0.1.0",
	}
}

// AppTag returns the application tagline.
func (app *App) AppTag() string {
	return (*app).tag
}

// AppVer returns the application version.
func (app *App) AppVer() string {
	return (*app).ver
}

// Discord returns the configuration of the Discord bot.
func (app *App) Discord() Discord {
	return (*app).discord
}

// Database returns the configuration of the Database instance.
func (app *App) Database() Database {
	return (*app).db
}

// SaveConfig saves the current instance of App to a file.
func (app *App) SaveConfig() error {
	str, err := toml.Marshal(&app)
	if err != nil {
		fmt.Printf("error marshalling application configuration: %d", err)
		return err
	}

	f, err0 := os.Create("config.toml")
	if err0 != nil {
		fmt.Printf("error creating configuration: %d", err0)
		return err0
	}

	defer f.Close()



	_, err1 := f.Write(str)
	if err1 != nil {
		fmt.Printf("error writing to file: %d", err1)
		return err1
	}

	f.Sync()



	return nil
}

// LoadConfig loads the configuration from its file.
func (app *App) LoadConfig() (_ *App, e error) {
	if app != AppConfig() {
		return app, errors.New("default configuration not present")
	}

	f, err := os.Open("config.toml")
	if err != nil {
		fmt.Printf("error opening file: %d", err)
		return app, err
	}

	defer f.Close()



	str := new(string)
	_, err0 := f.Read([]byte(*str))
	if err0 != nil {
		fmt.Printf("error reading from file: %d", err0)
		return app, err0
	}

	config, err1 := toml.Load(string(*str))
	if err1 != nil {
		fmt.Printf("error loading config from file: %d", err1)
		return app, err1
	}

	err2 := config.Unmarshal(&app)
	if err2 != nil {
		fmt.Printf("error unmarshalling configuration: %d", err2)
		return app, err2
	}



	return app, nil
}
