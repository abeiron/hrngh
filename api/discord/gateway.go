// Discord bindings for the Hrngh bot.
// Available at https://github.com/abeiron/hrngh

// Copyright 2020-2021, Undying Memory <abeiron@outlook.com>.  All rights reserved.
// Use of this source code is governed by the Microsoft Public License
// that can be found in the LICENSE file.

// Contains code relating to the Discord gateway.

package discord

// GatewayBotResponse stores the data for the gateway/bot response
type GatewayBotResponse struct {
  URL    string `json:"url"`
  Shards int    `json:"shards"`
}

// GatewayStatusUpdate is sent by the client to indicate a presence or status update
// https://discord.com/developers/docs/topics/gateway#update-status-gateway-status-update-structure
type GatewayStatusUpdate struct {
  Since  int      `json:"since"`
  Game   Activity `json:"game"`
  Status string   `json:"status"`
  AFK    bool     `json:"afk"`
}

type GatewayResponse struct {
  Url string `json:"url"`
  Shards int `json:"shards"`
}
