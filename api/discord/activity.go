// Discord bindings for the Hrngh bot.
// Available at https://github.com/abeiron/hrngh

// Copyright 2020-2021, Undying Memory <abeiron@outlook.com>.  All rights reserved.
// Use of this source code is governed by the Microsoft Public License
// that can be found in the LICENSE file.

package discord

// Activity defines the Activity sent with GatewayStatusUpdate.
//
// https://discord.com/developers/docs/topics/gateway#activity-object
type Activity struct {
	Name string       `json:"name"`
	Type ActivityType `json:"type"`
	Url  string       `json:"url, omitempty"`
}

// ActivityType is the type of activity in the Activity struct.
//
// See ActivityType* consts.
//
// https://discord.com/developers/docs/topics/gateway#activity-object-activity-types
type ActivityType int

// Valid ActivityType values.
const (
	ActivityTypeGame ActivityType = iota
	ActivityTypeStreaming
	ActivityTypeListening
	ActivityTypeWatching
	ActivityTypeCustom = 4
)

// An Assets struct contains assets and labels used in the rich presence "playing .." Game
type Assets struct {
	LargeImageID string `json:"large_image,omitempty"`
	SmallImageID string `json:"small_image,omitempty"`
	LargeText    string `json:"large_text,omitempty"`
	SmallText    string `json:"small_text,omitempty"`
}
