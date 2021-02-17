// WebSocket implementation for the Hrngh bot.
// Available at https://github.com/abeiron/hrngh

// Copyright 2020-2021, Undying Memory <abeiron@outlook.com>.  All rights reserved.
// Use of this source code is governed by the Microsoft Public License
// that can be found in the LICENSE file.

package ws

import (
	"bufio"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/abeiron/hrngh/internal/config"
)

const (
	ProtocolVersionHybi13    = 13
	ProtocolVersionHybi      = ProtocolVersionHybi13
	SupportedProtocolVersion = "13"

	ContinuationFrame = 0
	TextFrame         = 1
	BinaryFrame       = 2
	CloseFrame        = 8
	PingFrame         = 9
	PongFrame         = 10
	UnknownFrame      = 255

	DefaultMaxPayloadBytes = 32 << 20 // 32MB
)

// ProtocolError represents WebSocket protocol errors.
type ProtocolError struct {
	Error string
}

func (err *ProtocolError) Error() string { return err.Error }

var (
	ErrBadProtocolVer   = &ProtocolError{"bad protocol version"}
	ErrBadSchema        = &ProtocolError{"bad schema"}
	ErrBadStatus        = &ProtocolError{"bad status"}
	ErrBadUpgrade       = &ProtocolError{"missing or bad upgrade"}
	ErrBadOrigin        = &ProtocolError{"missing or bad WebSocket-Origin"}
	ErrBadLocation      = &ProtocolError{"missing or bad WebSocket-Location"}
	ErrBadProtocol      = &ProtocolError{"missing or bad WebSocket-Protocol"}
	ErrBadVersion       = &ProtocolError{"missing or bad WebSocket-Version"}
	ErrChallengeResp    = &ProtocolError{"mismatch challenge/response"}
	ErrBadFrame         = &ProtocolError{"bad frame"}
	ErrBadFrameBoundary = &ProtocolError{"not on frame boundary"}
	ErrNotWebSocket     = &ProtocolError{"not websocket protocol"}
	ErrBadReqMethod     = &ProtocolError{"bad method"}
	ErrUnsupported      = &ProtocolError{"unsupported"}
)

// ErrFrameTooLarge is returned by Codec's Receive method if payload size
// exceeds limitation set by Conn.MaxPayloadBytes
var ErrFrameTooLarge = errors.New("websocket: frame payload size exceeds limitation")

// Address is an implementation of net.Addr for WebSocket
type Address struct {
	*url.Url
}

// Network returns the network type for a WebSocket: "websocket".
func (addr *Address) Network() string { return "websocket" }

// serverHandshaker is an interface to handle WebSocket server-side handshake.
type serverHandshaker interface {
	// ReadHandshake reads handshake request message from client.
	// Returns http response code and error, if any.
	ReadHandshake(buf *bufio.Reader, req *http.Request) (code int, err error)

	// AcceptHandshake accepts the client handshake request and sends
	// handshake request back to client.
	AcceptHandshake(buf *bufio.Writer) (err error)

	// NewServerConn creates a new WebSocket connection.
	NewServerConn(buf *bufio.ReadWriter, rwc io.ReadWriteCloser, req *http.Request) (conn *Conn)
}

// frameReader is an interface to read a WebSocket frame.
type frameReader interface {
	// Reader is to read payload of the frame.
	io.Reader

	// PayloadType returns the type of the payload.
	PayloadType() byte

	// HeaderReader returns a header to read the header of the frame.
	HeaderReader() io.Reader

	// TrailerReader returns a reader to read the trailer of the frame.
	// If it returns nil, there is no trailer in the frame.
	TrailerReader() io.Reader

	// Len returns the total length of the frame, including header and trailer.
	Len() int
}

// frameReaderFactory is an interface to create new frame readers.
type frameReaderFactory interface {
	NewFrameReader() (r frameReader, err error)
}

// frameWriter is an interface to write a WebSocket frame.
type frameWriter interface {
	// Writer is to write to the payload of the frame.
	io.WriteCloser
}

// frameWriterFactory is an interface to create new frame writers.
type frameWriterFactory interface {
	NewFrameWriter(payloadType byte) (w frameWriter, err error)
}

type frameHandler interface {
	HandleFrame(frame frameReader) (r frameReader, err error)
	WriteClose(status int) (err error)
}
