package config

import (
	"github.com/bwmarrin/discordgo"
)

// Discord -
//
// Represents the Discord configuration.
type Discord struct {
	token 	string 				`toml:"token"`
	secret 	string 				`toml:"secret"`
	status 	discordgo.Status 	`toml:"status"`
}

// DiscordConfig -
//
// Creates a new instance of the default Discord configuration.
func DiscordConfig() Discord {
	return Discord{
		"BOT_TOKEN",
		"BOT_SECRET",
		discordgo.StatusOnline,
	}
}

// Token -
// 
// Returns the Discord token.
func (d Discord) Token() string {
	return d.token
}

// Secret -
//
// Returns the Discord secret.
func (d Discord) Secret() string {
	return d.secret
}

// SetToken -
// `token`: The token being supplied to the bot.
//
// Sets the token for the Discord bot connection.
func (d Discord) SetToken(token string) (_ Discord, err error) {
	d.token = token

	return d, err
}

// SetSecret -
// `secret`: The secret key for the Discord bot account.
//
// Sets the secret key for the Discord bot connection.
func (d Discord) SetSecret(secret string) (_ Discord, err error) {
	d.secret = secret

	return d, err
}

// SetStatus -
// `status`: The status to be displayed for the bot in the Discord client.
//
// Sets the status for the bot account.
func (d Discord) SetDiscordStatus(status discordgo.Status) (_ Discord, err error) {
	d.status = status

	return d, err
}
