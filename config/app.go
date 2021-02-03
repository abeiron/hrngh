package config

import "github.com/bwmarrin/discordgo"

// App -
//
// Represents the customizable end-user configuration for the application.
type App struct {
	token  string
	status discordgo.Status
}

// NewAppConfig -
//
// Creates a new instance of the default app configuration.
func NewAppConfig() App {
	return App{
		"BOT TOKEN",
		discordgo.StatusOnline,
	}
}

// AppToken -
//
// Returns the string representation of the Discord app token.
func (app App) AppToken() string {
	return app.token
}
