package discord

import (
	"errors"
)

type Status = int

const (
	StatusOnline 		Status = 0
	StatusIdle			Status = 1
	StatusDoNotDisturb 	Status = 2
	StatusOffline		Status = 3
	StatusInvisible		Status = 4
)

type Client struct {
	token string
	secret string
}

// New creates a new instance of a Discord client.
//
// `token`: Represents the bot token required for connecting to the Discord API.
func New(token string) (c Client, e error) {

	return c, nil
}
