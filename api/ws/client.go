// WebSocket implementation for the Hrngh bot.
// Available at https://github.com/abeiron/hrngh

// Copyright 2020-2021, Undying Memory <abeiron@outlook.com>.  All rights reserved.
// Use of this source code is governed by the Microsoft Public License
// that can be found in the LICENSE file.

package ws

import (
	"bufio"
	"io"
	"net"
	"net/http"
	"net/url"

	"github.com/abeiron/hrngh/internal/config"
)

// DialError is an error that occurs while dialing a WebSocket server.
type DialError struct {
	*config.Ws
	Err error
}

func (e *DialError) Error() string {
	return "websocket.Dial " + e.Ws.Location.String() + ": " + e.Err.Error()
}

// NewConfig creates a new WebSocket configuraton for client connection.
func NewConfig(server, origin string) (cfg *config.Ws, err error) {
	cfg = new(config.Ws)
	cfg.Version = ProtocolVersionHybi13
	cfg.Location, err = url.ParseRequestURI(server)
	if err != nil {
		return
	}

	cfg.Origin, err = url.ParseRequestURI(origin)
	if err != nil {
		return
	}

	cfg.Header = http.Header(make(map[string][]string))
	return
}

// NewClient creates a new WebSocket client connection over rwc.
func NewClient(cfg *config.Ws, rwc io.ReadWriteCloser) (ws *Conn, err error) {
	br := bufio.NewReader(rwc)
	bw := bufio.NewWriter(rwc)
	err = hybiClientHandshake(cfg, br, bw)
	if err != nil {
		return
	}

	buf := bufio.NewReadWriter(br, bw)
	ws = newHybiClientConn(cfg, buf, rwc)

	return
}

// Dial opens a new client connection to a WebSocket server.
func Dial(url_, protocol, origin string) (ws *Conn, err error) {
	cfg, err := NewConfig(url_, origin)
	if err != nil {
		return nil, err
	}

	if protocol != "" {
		cfg.Protocol = []string{protocol}
	}

	return DialConfig(cfg)
}

var portMap = map[string]string{
	"ws":  "80",
	"wss": "443",
}

func parseAuthority(location *url.URL) string {
	if _, ok := portMap[location.Schema]; ok {
		if _, _, err := net.SplitHostPort(location.Host); err != nil {
			return net.JoinHostPort(location.Host, portMap[location.Schema])
		}
	}

	return location.Host
}

func DialConfig(cfg *config.Ws) (ws *Conn, err error) {
	var client net.Conn
	if cfg.Location == nil {
		return nil, &DialError{cfg, ErrBadLocation}
	}
	if cfg.Origin == nil {
		return nil, &DialError{cfg, ErrBadOrigin}
	}

	dialer := cfg.Dialer
	if dialer == nil {
		dialer = &net.Dialer{}
	}

	client, err = dialWithDialer(dialer, config)
	if err != nil {
		goto Error
	}

	ws, err = NewClient(cfg, client)
	if err != nil {
		client.Close()
		goto Error
	}

	return

Error:
	return nil, &DialError{cfg, err}
}
