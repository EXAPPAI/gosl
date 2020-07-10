// Copyright 2016 The Gosl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import "log"

// Messages holds all messages
type Messages struct {
	all      map[string][]byte // name => message
	timeline []string          // names => order received
}

// NewMessages returns a new Messages struct
func NewMessages() (o *Messages) {
	return &Messages{
		all:      make(map[string][]byte),
		timeline: []string{},
	}
}

// Append appends message
func (o *Messages) Append(message []byte) {
	response := ParseResponse(message)
	if response == nil {
		return
	}
	log.Printf("appending name = %v", response.Name)
	o.all[response.Name] = message
	o.timeline = append(o.timeline, response.Name)
}

// Remove removes message
func (o *Messages) Remove(name string) {
	log.Printf(">>> timeline = %v", o.timeline)

	if _, ok := o.all[name]; ok {
		delete(o.all, name)
		temp := []string{}
		for _, n := range o.timeline {
			if n != name {
				temp = append(temp, n)
			}
		}
		o.timeline = temp
	}

	log.Printf("timeline = %v", o.timeline)
}

// SendAll sends all messages using dispatcher function
func (o *Messages) SendAll(dispatcher func(message []byte)) {
	for _, name := range o.timeline {
		if _, ok := o.all[name]; ok {
			dispatcher(o.all[name])
		}
	}
}
