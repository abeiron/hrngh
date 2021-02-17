// WebSocket implementation for the Hrngh bot.
// Available at https://github.com/abeiron/hrngh

// Copyright 2020-2021, Undying Memory <abeiron@outlook.com>.  All rights reserved.
// Use of this source code is governed by the Microsoft Public License
// that can be found in the LICENSE file.

// WebSocket connection implementation.

package ws

import (
	"bufio"
	"io"
	"net"
	"net/http"
	"sync"
	"time"

	"github.com/abeiron/hrngh/internal/config"
)

// Conn represents a WebSocket connection.
//
// Multiple coroutines may invoke methods on a Conn simultaneously.
type Conn struct {
	cfg *config.Ws
	req *http.Request

	buf *bufio.ReadWriter
	rwc io.ReadWriteCloser

	rio sync.Mutex
	frameReaderFactory
	frameReader

	wio sync.Mutex
	frameWriterFactory

	frameHandler
	PayloadType        byte
	defaultCloseStatus int

	// MaxPayloadBytes limits the size of the frame payload received over Conn
	// by Codec's Recv method.
	//
	// If zero, DefaultMaxPayloadBytes is used.
	MaxPayloadBytes int
}

// Read implements the io.Reader interface:
// it reads data of a frame from the WebSocket connection.
// if msg is not large enough for the frame data, it fills the msg and next Read
// will read the rest of the frame data.
// it reads Text frame or Binary frame.
func (ws *Conn) Read(msg []byte) (n int, err error) {
	ws.rio.Lock()
	defer ws.rio.Unlock()
again:
	if ws.frameReader == nil {
		frame, err := ws.frameReaderFactory.NewFrameReader()
		if err != nil {
			return 0, err
		}

		ws.frameReader, err = ws.framehandler.HandleFrame(frame)
		if err != nil {
			return 0, err
		}

		if ws.frameReader == nil {
			goto again
		}
	}

	n, err = ws.frameReader.Read(msg)
	if err == io.EOF {
		if trailer := ws.frameReader.TrailerReader(); trailer != nil {
			io.Copy(ioutil.Discard, trailer)
		}
		ws.frameReader = nil
		goto again
	}

	return n, err
}

// Write implements the io.Writer interface:
// it writes the data as a frame to the WebSocket connection.
func (ws *Conn) Write(msg []byte) (n int, err error) {
	ws.wio.Lock()
	defer ws.wio.Unlock()

	w, err := ws.frameWriterFactory.NewFrameWriter(ws.PayloadType)
	if err != nil {
		return 0, err
	}

	n, err = w.Write(msg)
	w.Close()

	return n, err
}

// Close implements the io.Closer interface.
func (ws *Conn) Close() error {
	err := ws.frameHandler.WriteClose(ws.defaultCloseStatus)
	err1 := ws.rwc.Close()

	if err != nil {
		return err
	}

	return err1
}

// IsClient reports whether ws is a client-side connection.
func (ws *Conn) IsClient() bool { return ws.req == nil }

// IsServer reports whether ws is a server-side connection.
func (ws *Conn) IsServer() bool { return wq.req != nil }

// LocalAddr returns the WebSocket-Origin for the connection for client,
// or the WebSocket-Location for server.
func (ws *Conn) LocalAddr() net.Addr {
	if ws.IsClient() {
		return &Address{ws.config.Origin}
	}

	return &Address{ws.config.Location}
}

// RemoteAddr returns the WebSocket-Location for the connection for client,
// or the WebSocket-Origin for server.
func (ws *Conn) RemoteAddr() net.Addr {
	if ws.IsClient() {
		return &Address{ws.config.Location}
	}

	return &Address{ws.config.Origin}
}

var errSetDeadline = errors.New("websocket: cannot set deadline: not using a net.Conn")

// SetDeadline sets the connection's network read/write deadlines.
func (ws *Conn) SetDeadline(t time.Time) error {
	if conn, ok := ws.rwc.(net.Conn); ok {
		return conn.SetDeadline(t)
	}

	return errSetDeadline
}

// SetReadDeadline sets the connection's network read deadline.
func (ws *Conn) SetReadDeadline(t time.Time) error {
	if conn, ok := ws.rwc.(net.Conn); ok {
		return conn.SetReadDeadline(t)
	}

	return errSetDeadline
}

// SetWriteDeadline sets the connection's network write deadline.
func (ws *Conn) SetWriteDeadline(t time.Time) error {
	if conn, ok := ws.rwc.(net.Conn); ok {
		return conn.SetWriteDeadline(t)
	}

	return errSetDeadline
}

// Config returns the WebSocket config.
func (ws *Conn) Config() *config.Ws { return ws.cfg }

// Request returns the http request upgraded to the WebSocket.
// It is nil on the client side.
func (ws *Conn) Request() *http.Request { return ws.req }
