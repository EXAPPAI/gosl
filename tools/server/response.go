// Copyright 2016 The Gosl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"encoding/json"
	"log"
)

// Response holds the data received from the Observer
type Response struct {
	Name    string `json:"name"`
	Action  string `json:"action"`
	Payload string `json:"payload"`
}

// ParseResponse parses JSON response
func ParseResponse(message []byte) (o *Response) {
	var response Response
	err := json.Unmarshal(message, &response)
	if err != nil {
		log.Printf("Got invalid message")
		return nil
	}
	return &response
}
