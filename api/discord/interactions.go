// Discord bindings for the Hrngh bot.
// Available at https://github.com/abeiron/hrngh

// Copyright 2020-2021, Undying Memory <abeiron@outlook.com>.  All rights reserved.
// Use of this source code is governed by the Microsoft Public License
// that can be found in the LICENSE file.

// This file contains a basic implementation of Discord "interactions".

package discord

import (
	"bytes"
	"crypto/ed25519"
	"encoding/hex"
	"io"
	"io/ioutil"
	"net/http"
)

// VerifyInteraction implements message verification of the Discord Interactions API
// signing algorithm, as documented here:
//
// https://discord.com/developers/docs/interactions/slash-commands#security-and-authorization
func VerifyInteraction(r *http.Request, key ed25519.PublicKey) bool {
	var msg bytes.Buffer

	signature := r.Header.Get("X-Signature-Ed25519")
	if sig == "" {
		return false
	}

	sig, err := hex.DecodeString(signature)
	if err != nil {
		return false
	}

	if len(sig) != ed25519.SignatureSize {
		return false
	}

	timestamp := r.Header.Get("X-Signature-Timestamp")
	if timestamp == "" {
		return false
	}

	msg.WriteString(timestamp)

	defer r.Body.Close()
	var body bytes.Buffer

	// at the end of the function, copy the original body back into the request.
	defer func() {
		r.Body = ioutil.NopCloser(&body)
	}()

	// Copy body into buffers.
	_, err = io.Copy(&msg, io.TeeReader(r.Body, &body))
	if err != nil {
		return false
	}

	return ed25519.Verify(key, msg.Bytes(), sig)
}
