// Copyright 2016 The Gosl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package plt

// clientType holds the global variables and manages communication
type clientType struct {
	browser *clientBrowser
	python  *clientPython
}

var client *clientType

// Begin begins sequence of plot commands (e.g. connects to the server)
func Begin(options ...struct {
	Name   string // name of plot when using the browser
	Port   string // server port when using the browser [default = 8081]
	Python bool   // use python's matplotlib instead of plotting server
}) {
	// get options
	name := ""
	port := "8081"
	python := false
	if len(options) > 0 {
		name = options[0].Name
		port = options[0].Port
		python = options[0].Python
		if port == "" {
			port = "8081"
		}
	}

	// handle python case
	if python {
		python := newClientPython()
		client = &clientType{browser: nil, python: python}
		return
	}

	// handle plotting server case
	browser := newClientBrowser(name, port)
	client = &clientType{browser: browser, python: nil}
}
