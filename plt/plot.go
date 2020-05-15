// Copyright 2016 The Gosl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package plt

import (
	"github.com/cpmech/gosl/chk"
	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
)

// Plot plots curve
func Plot(x, y []float64, style ...CurveStyle) {
	var s CurveStyle
	if len(style) > 0 {
		s = style[0]
	} else {
		s = DefaultCurveStyle
	}
	// curves.List = append(curves.List, Curve{X: x, Y: y, Style: s})
	curve := &Curve{X: x, Y: y, Style: s}
	json := curve.Encode()
	err := wsutil.WriteClientMessage(connection, ws.OpText, json)
	if err != nil {
		chk.Panic("cannot send curve data to server")
	}

}
