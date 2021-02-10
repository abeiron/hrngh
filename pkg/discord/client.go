package discord

import (
	"errors"
	"sync"
)

type Status string

const (
	StatusOnline 		Status = "online"
	StatusIdle			Status = "idle"
	StatusDoNotDisturb 	Status = "dnd"
	StatusOffline		Status = "offline"
	StatusInvisible		Status = "invisible"
)

type Client struct {
	sync.RWMutex

	// Generally configurable settings.

	// Identify is sent during the initial handshake with the Discord gateway.
	//
	// https://discord.com/developers/docs/topics/gateway#identify
	Identify Identify

	MFA bool

	LogLevel int

	// Should the client reconnect the websocket on error?
	ReconnectOnError bool

	ShardId int
	ShardCount int

	// Should state tracking be enabled?
	// 
	// State tracking is the best way for getting the users'
	// active guilds and the members of the guilds.
	StateEnabled bool
}

// Identify is sent during initial handshake with the Discord gateway.
//
// https://discord.com/developers/docs/topics/gateway#identify
type Identify struct {
	Token string `json:"token"`
	Properties IdentifyProperties `json:"properties"`
	Compress bool `json:"compress"`
	LargeThreshold int `json:"large_threshold"`
	Shard *[2]int `json:"shard,omitempty"`
	Presence GatewayStatusUpdate `json:"presence,omitempty"`
	GuildSubscriptions bool `json:"guild_subscriptions"`
	Intents Intent `json:"intents"`
}


