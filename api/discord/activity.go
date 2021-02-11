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
	Name string 		`json:"name"`
	Type ActivityType 	`json:"type"`
	Url string 			`json:"url, omitempty"`
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
