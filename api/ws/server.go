// WebSocket implementation for the Hrngh bot.
// Available at https://github.com/abeiron/hrngh

// Copyright 2020-2021, Undying Memory <abeiron@outlook.com>.  All rights reserved.
// Use of this source code is governed by the Microsoft Public License
// that can be found in the LICENSE file.

// This file contains code related to serving a websocket.

package ws

import (
	"bufio"
	"fmt"
	"io"
	"net/http"

	"github.com/abeiron/hrngh/internal/config"
)

func newServerConn(rwc io.ReadWriteCloser, buf *bufio.ReadWriter, config *config.Ws, handshake func(*config.Ws, *http.Request) error) (conn *Conn, err error) {
	var hs serverHandshaker = &hybiServerHandshaker{Config: config}

}
