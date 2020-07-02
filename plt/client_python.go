// Copyright 2016 The Gosl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package plt

import (
	"bytes"

	"github.com/cpmech/gosl/io"
)

// clientPython holds data for python (matplotlib) plotting session
type clientPython struct {
	buffer    *bytes.Buffer // buffer holding Python commands
	bufferEa  *bytes.Buffer // buffer holding Python extra artists commands
	created3d bool          // flag indicating that Python axes3d ('AX3D') has been created
}

// newClientPython creates new clientPython
func newClientPython() (o *clientPython) {
	return &clientPython{
		buffer:    new(bytes.Buffer),
		bufferEa:  new(bytes.Buffer),
		created3d: false,
	}
}

// write writes commands to buffer
func (o *clientPython) write(commands string, extraArtists ...string) {
	io.Ff(o.buffer, commands)
	if len(extraArtists) > 0 {
		ea := extraArtists[0]
		io.Ff(o.bufferEa, ea)
	}
}

// uid returns unique identifier
func (o *clientPython) uid() int {
	return o.buffer.Len()
}

// plot plots x-y series
func (o *clientPython) plot(curve *Curve) {
	uid := o.uid()
	sx := io.Sf("x%d", uid)
	sy := io.Sf("y%d", uid)
	pythonGen2Arrays(o.buffer, sx, sy, curve.X, curve.Y)
	io.Ff(o.buffer, "plt.plot(%s,%s", sx, sy)
	io.Ff(o.buffer, curve.pythonParams())
	io.Ff(o.buffer, ")\n")
}

// plotOne plots one point
func (o *clientPython) plotOne(x, y float64, style ...CurveStyle) {
	io.Ff(o.buffer, "plt.plot(%23.15e,%23.15e", x, y)
	if len(style) > 0 {
		io.Ff(o.buffer, style[0].pythonParams())
	}
	io.Ff(o.buffer, ")\n")
}
