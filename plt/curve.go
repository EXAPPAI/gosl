// Copyright 2016 The Gosl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package plt

import (
	"bytes"

	"github.com/cpmech/gosl/io"
	"github.com/cpmech/gosl/utl"
)

// Curve defines the curve data
type Curve struct {
	Style         CurveStyle `json:"style"`         // line and marker arguments
	Label         string     `json:"label"`         // curve name or connection pair such as 'San Francisco -> Los Angeles'
	X             []float64  `json:"x"`             // x-coordinates
	Y             []float64  `json:"y"`             // y-coordinates
	Z             []float64  `json:"z,omitempty"`   // [optional] z-coordinates
	Kind          string     `json:"kind"`          // e.g. connection, city, fortress, base, mine, ...
	TagFirstPoint bool       `json:"tagFirstPoint"` // tag first point with label

	// for Python only
	Zindex float64 `json:"-"` // figure elevation (z-index)
	NoClip bool    `json:"-"` // turn clipping off
}

// Encode encodes Curve into JSON string
func (o *Curve) Encode() []byte {
	buf := new(bytes.Buffer)
	enc := utl.NewEncoder(buf, "json")
	enc.Encode(o)
	return buf.Bytes()
}

// PythonParams returns python options
func (o *Curve) PythonParams() (l string) {
	l = o.Style.PythonParams()
	if o.Label != "" {
		l += io.Sf(",label='%s'", o.Label)
	}
	if o.Zindex > 0 {
		l += io.Sf(",zorder=%d", o.Zindex)
	}
	if o.NoClip {
		l += ",clip_on=False"
	}
	return
}
