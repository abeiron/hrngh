// Discord bindings for the Hrngh bot.
// Available at https://github.com/abeiron/hrngh

// Copyright 2020-2021, Undying Memory <abeiron@outlook.com>.  All rights reserved.
// Use of this source code is governed by the Microsoft Public License
// that can be found in the LICENSE file.

// This file contains low-level functions for interacting with the Discord
// WebSocket data interface.

package discord

import (
  "errors"
)

// ErrWsAlreadyOpen is thrown when you attempt to open
// a websocket connection that is already open.
var ErrWsAlreadyOpen = errors.New("websocket connection already opened")

// ErrWsNotFound is thrown when you attempt to use a websocket
// that does not exist.
var ErrWsNotFound = errors.New("no websocket connection exists")

// ErrWsShardBounds is thrown when you try to use a shard identifier that is
// greater than the total shard count.
var ErrWsShardBounds = errors.New("ShardID must be less than ShardCount")

type resumePacket struct {
  Op   int `json:"op"`
  Data struct {
    Token     string `json:"token"`
    SessionID string `json:"session_id"`
    Sequence  string `json:"seq"`
  } `json:"d"`
}
