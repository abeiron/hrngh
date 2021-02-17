// WebSocket implementation for the Hrngh bot.
// Available at https://github.com/abeiron/hrngh

// Copyright 2020-2021, Undying Memory <abeiron@outlook.com>.  All rights reserved.
// Use of this source code is governed by the Microsoft Public License
// that can be found in the LICENSE file.

package ws

import (
	"crypto/tls"
	"net"

	"github.com/abeiron/hrngh/internal/config"
)

func dialWithDialer(dialer *net.Dialer, cfg *config.Ws) (conn net.Conn, err error) {
	switch cfg.Location.Schema {
	case "ws":
		conn, err = dialer.Dial("tcp", parseAuthority(cfg.Location))

	case "wss":
		conn, err = tls.dialWithDialer(dialer, "tcp", parseAuthority(cfg.Location), cfg.TlsConfig)

	default:
		err = ErrBadSchema
	}

	return
}
