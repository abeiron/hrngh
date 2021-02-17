// Discord bindings for the Hrngh bot.
// Available at https://github.com/abeiron/hrngh

// Copyright 2020-2021, Undying Memory <abeiron@outlook.com>.  All rights reserved.
// Use of this source code is governed by the Microsoft Public License
// that can be found in the LICENSE file.

// The code in this file works to implement rudimentary support for Discord Voice.

package discord

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/abeiron/hrngh/api/ws"
	"golang.org/x/crypto/nacl/secretbox"
)

// ------------------------------------------------------------------------------------------------
// Code related to both VoiceConnection Websocket and UDP connections.
// ------------------------------------------------------------------------------------------------

type VoiceConnection struct {
	sync.RWMutex

	LogLevel     int
	Ready        bool // if true, bot is ready to send/receive audio.
	UserID       string
	GuildID      string
	ChannelID    string
	deaf         bool
	mute         bool
	speaking     bool
	reconnecting bool // if true, voice connection is attempting to reconnect.

	OpusSend chan []byte  // Chan for sending opus audio
	OpusRecv chan *Packet // Chan for receiving opus audio

	wsConn  *ws.Conn
	wsMutex sync.Mutex
	udpConn *net.UDPConn
	session *Session

	sessionID string
	token     string
	endpoint  string

	// Used to send a close signal to any open coroutines.
	close chan struct{}

	// Used to allow blocking until connected.
	connected chan bool

	// Used to pass the session identifier from onVoiceStateUpdate
	sessionRecv chan string

	op4 voiceOP4
	op2 voiceOP2

	voiceSpeakingUpdatehandlers []VoiceSpeakingUpdateHandler
}

// VoiceSpeakingUpdateHandler type provides a functiond definition for the
// VoiceSpeakingUpdate event.
type VoiceSpeakingUpdateHandler func(vc *VoiceConnection, vs *VoiceSpeakingUpdate)

// Speaking sends a speaking notification to Discord over the voice websocket.
// This must be sent as true prior to sending audio and should be set to false
// once finished sending audio.
//  b : Send true if speaking and false if not.
func (v *VoiceConnection) Speaking(b bool) (err error) {
	v.log(LogDebug, "called (%t)", b)

	type voiceSpeakingData struct {
		Speaking bool `json:"speaking"`
		Delay    int  `json:"delay"`
	}

	type voiceSpeakingOp struct {
		Op   int               `json:"op"` // Always 5
		Data voiceSpeakingData `json:"d"`
	}

	if v.wsConn == nil {
		return fmt.Errorf("no VoiceConnection websocket")
	}

	data := voiceSpeakingOp{5, voiceSpeakingData{b, 0}}
	v.wsMutex.Lock()
	err = v.wsConn.WriteJSON(data)
	v.wsMutex.Unlock()

	v.Lock()
	defer v.Unlock()
	if err != nil {
		v.speaking = false
		v.log(LogError, "Speaking() write json error, %s", err)
		return
	}

	v.speaking = b

	return
}

// ChangeChannel sends Discord a request to change channels within a Guild.
//
// Note: this function may or may not be used.
// May simply use "VoiceChannelJoin" instead.
func (v *VoiceConnection) ChangeChannel(channelID string, mute, deaf bool) (err error) {
	v.log(LogInformation, "called")

	data := voiceChannelJoinOp{4, voiceChannelJoinData{&v.GuildID, &channelID, mute, deaf}}
	v.wsMutex.Lock()
	err = v.session.wsConn.WriteJSON(data)
	v.wsMutex.Unlock()

	if err != nil {
		return
	}

	v.ChannelID = channelID
	v.deaf = deaf
	v.mute = mute
	v.speaking = false

	return
}

// TODO: LINE 139 // CHECK NOTES // PAGE 49934, LINE 19-399
