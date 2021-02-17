// WebSocket implementation for the Hrngh bot.
// Available at https://github.com/abeiron/hrngh

// Copyright 2020-2021, Undying Memory <abeiron@outlook.com>.  All rights reserved.
// Use of this source code is governed by the Microsoft Public License
// that can be found in the LICENSE file.

package ws

// Codec represents a symmetric pair of functions that implement a codec.
type Codec struct {
	Marshal   func(v interface{}) (data []byte, pt byte, err error)
	Unmarshal func(data []byte, pt byte, v interface{}) (err error)
}

// Send sends a 'v' marshalled by c.Marshal as a single frame to 'ws'.
func (c Codec) Send(ws *Conn, v interface{}) (err error) {
	data, pt, err := c.Marshal(v)
	if err != nil {
		return err
	}

	ws.wio.Lock()
	defer ws.wio.Unlock()

	w, err := ws.frameWriterFactory.NewFrameWriter(pt)
	if err != nil {
		return nil
	}

	_, err = w.Write(data)
	w.Close()

	return err
}

// Recv receives single frame from ws, unmarshalled by c.Unmarshal and stores
// in 'v'. The whole frame payload is read to an in-memory buffer; max size of
// payload is defined by ws.MaxPayloadBytes. If frame payload size exceeds
// limit, ErrFrameTooLarge is returned; in this case frame is not read off wire
// completely. The next call to Receive would read and discard leftover data of
// previous oversized frame before processing next frame.
func (c Codec) Recv(ws *Conn, v interface{}) (err error) {
	ws.rio.Lock()
	defer ws.rio.Unlock()

	if ws.frameReader != nil {
		_, err = io.Copy(ioutil.Discard, ws.frameReader)
		if err != nil {
			return err
		}

		ws.frameReader = nil
	}

again:
	frame, err := ws.frameReaderFactory.NewFrameReader()
	if err != nil {
		return err
	}

	frame, err = ws.frameHandler.HandleFrame(frame)
	if err != nil {
		return err
	}

	if frame == nil {
		goto again
	}

	maxPayloadBytes := ws.MaxPayloadBytes
	if maxPayloadBytes == 0 {
		maxPayloadBytes = DefaultMaxPayloadBytes
	}

	if hf, ok := frame.(*hybiFrameReader); ok && hf.header.Length > int64(maxPayloadBytes) {
		// Payload size exceeds limit, so there is no need to call Unmarshal.
		//
		// Set frameReader to current oversized frame so that
		// the next call to this function can drain any leftover
		// data before processing the next frame.
		ws.frameReader = frame
		return ErrFrameTooLarge
	}

	pt := frame.PayloadType()
	data, err := ioutil.ReadAll(frame)
	if err != nil {
		return err
	}

	return c.Unmarshal(data, pt, v)
}

func marshal(v interface{}) (msg []byte, pt byte, err error) {
	switch data := v.(type) {
	case string:
		return []byte(data), TextFrame, nil
	case []byte:
		return data, BinaryFrame, nil
	}

	return nil, UnknownFrame, ErrNotSupported
}

func unmarshal(msg []byte, pt byte, v interface{}) (err error) {
	switch data := v.(type) {
	case *string:
		*data = string(msg)
		return nil
	case *[]byte:
		*data = msg
		return nil
	}

	return ErrNotSupported
}

/*
Message is a codec to send/receive text/binary data in a frame on WebSocket connection.
To send/receive text frame, use string type.
To send/receive binary frame, use []byte type.
Trivial usage:
	import "github.com/abeiron/hrngh/api/ws"
	// receive text frame
	var message string
	ws.Message.Recv(ws, &message)
	// send text frame
	message = "hello"
	ws.Message.Send(ws, message)
	// receive binary frame
	var data []byte
	ws.Message.Recv(ws, &data)
	// send binary frame
	data = []byte{0, 1, 2}
	ws.Message.Send(ws, data)
*/
var Message = Codec{marshal, unmarshal}

func jsonMarshal(v interface{}) (msg []byte, payloadType byte, err error) {
	msg, err = json.Marshal(v)
	return msg, TextFrame, err
}

func jsonUnmarshal(msg []byte, payloadType byte, v interface{}) (err error) {
	return json.Unmarshal(msg, v)
}

/*
JSON is a codec to send/receive JSON data in a frame from a WebSocket connection.
Trivial usage:
	import "github.com/abeiron/hrngh/api/ws"
	type T struct {
		Msg string
		Count int
	}
	// receive JSON type T
	var data T
	ws.JSON.Receive(ws, &data)
	// send JSON type T
	ws.JSON.Send(ws, data)
*/
var JSON = Codec{jsonMarshal, jsonUnmarshal}
