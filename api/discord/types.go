// Discord bindings for the Hrngh bot.
// Available at https://github.com/abeiron/hrngh

// Copyright 2020-2021, Undying Memory <abeiron@outlook.com>.  All rights reserved.
// Use of this source code is governed by the Microsoft Public License
// that can be found in the LICENSE file.

// This file contains some useful types, currently only a wrapper for a Discord Timestamp.

package discord

import (
	"encoding/json"
	"net/http"
	"time"
)

// Timestamp stores a timestamp as sent by the Discord API.
type Timestamp string

// Parse parses a timestamp into a time.Time object.
//
// The only time this can fail is if Discord changes the timestamp format.
func (t Timestamp) Parse() (time.Time, error) {
	return time.Parse(time.RFC3339, string(t))
}

// RESTError stores error information about a request with a bad response code.
// Message is not always present, there are cases where api calls can fail
// without returning a json message.
type RESTError struct {
	Request      *http.Request
	Response     *http.Response
	ResponseBody []byte

	Message *APIErrorMessage // Message may be nil.
}

func newRestError(req *http.Request, resp *http.Response, body []byte) *RESTError {
	restErr := &RESTError{
		Request:      req,
		Response:     resp,
		ResponseBody: body,
	}

	// Attempt to decode the error and assume no message was provided if it fails
	var msg *APIErrorMessage
	err := json.Unmarshal(body, &msg)
	if err == nil {
		restErr.Message = msg
	}

	return restErr
}

func (r RESTError) Error() string {
	return "HTTP " + r.Response.Status + ", " + string(r.ResponseBody)
}
