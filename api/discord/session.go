// Discord bindings for the Hrngh bot.
// Available at https://github.com/abeiron/hrngh

// Copyright 2020-2021, Undying Memory <abeiron@outlook.com>.  All rights reserved.
// Use of this source code is governed by the Microsoft Public License
// that can be found in the LICENSE file.

// This file contains code related to the Session struct.

package discord

import (
  "errors"
  "net/http"
  "sync"
  "time"
)

// Represents the status of a member.
type Status string

const (
  StatusOnline        Status = "online"
  StatusIdle          Status = "idle"
  StatusDoNotDisturb  Status = "dnd"
  StatusOffline       Status = "offline"
  StatusInvisible     Status = "invisible"
)

// Session holds the information pertaining to the current shard's session.
type Session struct {
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

  // Sharding
  ShardId int
  ShardCount int

  // Should state tracking be enabled?
  // 
  // State tracking is the best way for getting the users'
  // active guilds and the members of the guilds.
  StateEnabled bool

  // Whether WebSocket data is ready.
  // May be deprecated soon.
  DataReady bool

  // Stores the correct status of the WebSocket connection.
  status bool

  // The HTTP client used for REST requests.
  Client *http.Client
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

// IdentifyProperties contains the "properties" portion of an Identify packet
//
// https://discord.com/developers/docs/topics/gateway#identify-identify-connection-properties
type IdentifyProperties struct {
  OS              string `json:"$os"`
  Browser         string `json:"$browser"`
  Device          string `json:"$device"`
  Referer         string `json:"$referer"`
  ReferringDomain string `json:"$referring_domain"`
}

