// Copyright 2016 The Gosl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package plt

// Plot plots curve
func Plot(x, y []float64, label string, style ...CurveStyle) {
	var s CurveStyle
	if len(style) > 0 {
		s = style[0]
	} else {
		s = DefaultCurveStyle
	}
	curve := &Curve{
		Style:         s,
		Label:         label,
		X:             x,
		Y:             y,
		Kind:          "",
		TagFirstPoint: false,
		Zindex:        0,
		NoClip:        false,
	}
	if client.browser == nil {
		client.python.plot(curve)
	} else {
		client.browser.plot(curve)
	}
}
