// WebSocket implementation for the Hrngh bot.
// Available at https://github.com/abeiron/hrngh

// Copyright 2020-2021, Undying Memory <abeiron@outlook.com>.  All rights reserved.
// Use of this source code is governed by the Microsoft Public License
// that can be found in the LICENSE file.

// This file contains configuration code for the 'ws' package.

package config

import (
	"crypto/tls"
	"net"
	"net/http"
	"net/url"
)

// Ws is a WebSocket configuration
type Ws struct {
	// A WebSocket server address.
	Location *url.Url

	// A WebSocket client origin
	Origin *url.Url

	// WebSocket subprotocols.
	Protocol []string

	// WebSocket protocol version
	Version int

	// TLS configuration for secure WebSocket communication. (wss).
	TlsConfig *tls.Config

	// Additional header fields to be sent in WebSocket in initial handshake.
	Header http.Header

	// Dialer used when opening WebSocket connections.
	Dialer *net.Dialer

	handshakeData map[string]string
}
