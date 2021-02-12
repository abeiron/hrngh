// Discord bindings for the Hrngh bot.
// Available at https://github.com/abeiron/hrngh

// Copyright 2020-2021, Undying Memory <abeiron@outlook.com>.  All rights reserved.
// Use of this source code is governed by the Microsoft Public License
// that can be found in the LICENSE file.

// This file contains code related to Discord "intents"

package discord

// Intent is the type of a Gateway Intent
// https://discord.com/developers/docs/topics/gateway#gateway-intents
type Intent int

// Constants for the different bit offsets of intents
const (
  IntentsGuilds Intent = 1 << iota
  IntentsGuildMembers
  IntentsGuildBans
  IntentsGuildEmojis
  IntentsGuildIntegrations
  IntentsGuildWebhooks
  IntentsGuildInvites
  IntentsGuildVoiceStates
  IntentsGuildPresences
  IntentsGuildMessages
  IntentsGuildMessageReactions
  IntentsGuildMessageTyping
  IntentsDirectMessages
  IntentsDirectMessageReactions
  IntentsDirectMessageTyping

  IntentsAllWithoutPrivileged = IntentsGuilds |
    IntentsGuildBans |
    IntentsGuildEmojis |
    IntentsGuildIntegrations |
    IntentsGuildWebhooks |
    IntentsGuildInvites |
    IntentsGuildVoiceStates |
    IntentsGuildMessages |
    IntentsGuildMessageReactions |
    IntentsGuildMessageTyping |
    IntentsDirectMessages |
    IntentsDirectMessageReactions |
    IntentsDirectMessageTyping
  IntentsAll = IntentsAllWithoutPrivileged |
    IntentsGuildMembers |
    IntentsGuildPresences
  IntentsNone Intent = 0
)

// MakeIntent used to help convert a gateway intent value for use in the Identify structure;
// this was useful to help support the use of a pointer type when intents were optional.
// This is now a no-op, and is not necessary to use.
func MakeIntent(intents Intent) Intent {
  return intents
}
