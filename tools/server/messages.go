// Copyright 2016 The Gosl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

// Messages holds all messages
type Messages struct {
	all map[string][][]byte // plotName => messages
}

// NewMessages returns a new Messages struct
func NewMessages() (o *Messages) {
	return &Messages{
		all: make(map[string][][]byte),
	}
}

// Append appends message
func (o *Messages) Append(message []byte) {
	response := ParseResponse(message)
	if response == nil {
		return
	}
	if _, ok := o.all[response.PlotName]; ok {
		o.all[response.PlotName] = append(o.all[response.PlotName], message)
	} else {
		o.all[response.PlotName] = [][]byte{message}
	}
}

// Remove removes message
func (o *Messages) Remove(plotName string) {
	delete(o.all, plotName)
}

// SendAll sends all messages using dispatcher function
func (o *Messages) SendAll(dispatcher func(message []byte)) {
	for _, messages := range o.all {
		for _, message := range messages {
			dispatcher(message)
		}
	}
}
