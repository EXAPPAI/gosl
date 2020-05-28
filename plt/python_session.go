// Copyright 2016 The Gosl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package plt

import (
	"bytes"

	"github.com/cpmech/gosl/io"
)

// pythonSession holds data for python (matplotlib) plotting session
type pythonSession struct {
	buffer    *bytes.Buffer // buffer holding Python commands
	bufferEa  *bytes.Buffer // buffer holding Python extra artists commands
	created3d bool          // flag indicating that Python axes3d ('AX3D') has been created
}

// newPythonSession creates new pythonSession
func newPythonSession() (o *pythonSession) {
	return &pythonSession{
		buffer:    new(bytes.Buffer),
		bufferEa:  new(bytes.Buffer),
		created3d: false,
	}
}

// write writes commands to buffer
func (o *pythonSession) write(commands string, extraArtists ...string) {
	io.Ff(o.buffer, commands)
	if len(extraArtists) > 0 {
		ea := extraArtists[0]
		io.Ff(o.bufferEa, ea)
	}
}

// uid returns unique identifier
func (o *pythonSession) uid() int {
	return o.buffer.Len()
}

// plot plots x-y series
func (o *pythonSession) plot(curve *Curve) {
	uid := o.uid()
	sx := io.Sf("x%d", uid)
	sy := io.Sf("y%d", uid)
	pythonGen2Arrays(o.buffer, sx, sy, curve.X, curve.Y)
	io.Ff(o.buffer, "plt.plot(%s,%s", sx, sy)
	io.Ff(o.buffer, ")\n")
}

// plotOne plots one point
func (o *pythonSession) plotOne(x, y float64, style ...CurveStyle) {
	io.Ff(o.buffer, "plt.plot(%23.15e,%23.15e", x, y)
	io.Ff(o.buffer, ")\n")
}
