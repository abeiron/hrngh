package config

import "github.com/abeiron/hrngh/api/discord"

// Discord represents the Discord configuration.
type Discord struct {
	token 	string 				`hcl:"token"`
	secret 	string 				`hcl:"secret"`
	status 	discord.Status 	`hcl:"status"`
}

// DiscordConfig creates a new instance of the default Discord configuration.
func DiscordConfig() *Discord {
	return *Discord{
		"BOT_TOKEN",
		"BOT_SECRET",
		discord.StatusOnline,
	}
}

// Token returns the Discord token.
func (d *Discord) Token() string {
	return (*d).token
}

// Secret returns the Discord secret.
func (d *Discord) Secret() string {
	return (*d).secret
}

// SetToken
//
// `token`: The token being supplied to the bot.
//
// Sets the token for the Discord bot connection.
func (d *Discord) SetToken(token string) (_ *Discord, err error) {
	d.token = token

	return d, err
}

// SetSecret
//
// `secret`: The secret key for the Discord bot account.
//
// Sets the secret key for the Discord bot connection.
func (d *Discord) SetSecret(secret string) (_ *Discord, err error) {
	d.secret = secret

	return d, err
}

// SetStatus
//
// `status`: The status to be displayed for the bot in the Discord client.
//
// Sets the status for the bot account.
func (d *Discord) SetStatus(status discord.Status) (_ *Discord, err error) {
	d.status = status

	return d, err
}
