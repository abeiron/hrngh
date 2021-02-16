// Discord bindings for the Hrngh bot.
// Available at https://github.com/abeiron/hrngh

// Copyright 2020-2021, Undying Memory <abeiron@outlook.com>.  All rights reserved.
// Use of this source code is governed by the Microsoft Public License
// that can be found in the LICENSE file.

// Contains structures for emoji support.

package discord

// Emoji struct holds data related to Emoji's
type Emoji struct {
  ID            string   `json:"id"`
  Name          string   `json:"name"`
  Roles         []string `json:"roles"`
  User          *User    `json:"user"`
  RequireColons bool     `json:"require_colons"`
  Managed       bool     `json:"managed"`
  Animated      bool     `json:"animated"`
  Available     bool     `json:"available"`
}

// MessageFormat returns a correctly formatted Emoji for use in Message content and embeds
func (e *Emoji) MessageFormat() string {
  if e.ID != "" && e.Name != "" {
    if e.Animated {
      return "<a:" + e.APIName() + ">"
    }

    return "<:" + e.APIName() + ">"
  }

  return e.APIName()
}

// APIName returns an correctly formatted API name for use in the MessageReactions endpoints.
func (e *Emoji) APIName() string {
  if e.ID != "" && e.Name != "" {
    return e.Name + ":" + e.ID
  }
  if e.Name != "" {
    return e.Name
  }
  return e.ID
}
