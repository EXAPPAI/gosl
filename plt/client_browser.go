// Copyright 2016 The Gosl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package plt

import (
	"bytes"
	"context"
	"net"
	"time"

	"github.com/cpmech/gosl/chk"
	"github.com/cpmech/gosl/utl"
	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
)

// clientBrowser holds information of the current plotting session
type clientBrowser struct {
	name       string
	connection net.Conn
}

// newClientBrowser creates new clientBrowser
func newClientBrowser(name string, port string) (o *clientBrowser) {
	dur, err := time.ParseDuration("1m")
	if err != nil {
		chk.Panic("INTERNAL ERROR: cannot parse duration")
	}
	ctx, cancel := context.WithTimeout(context.Background(), dur)
	defer cancel()
	connection, _, _, err := ws.DefaultDialer.Dial(ctx, "ws://localhost:"+port+"/provider")
	if err != nil {
		chk.Panic("cannot connect to plotting server")
	}
	return &clientBrowser{name, connection}
}

// encode encodes Session into JSON string
func (o *response) encode() []byte {
	buf := new(bytes.Buffer)
	enc := utl.NewEncoder(buf, "json")
	enc.Encode(o)
	return buf.Bytes()
}

// send sends message to server
func (o *clientBrowser) send(message []byte) {
	err := wsutil.WriteClientMessage(o.connection, ws.OpText, message)
	if err != nil {
		chk.Panic("cannot send message to server")
	}
}

// plot plots x-y series
func (o *clientBrowser) plot(curve *Curve) {
	data := &response{o.name, "plot", string(curve.Encode())}
	o.send(data.encode())
}
