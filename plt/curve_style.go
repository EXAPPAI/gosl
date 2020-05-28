// Copyright 2016 The Gosl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package plt

import "github.com/cpmech/gosl/io"

// CurveStyle holds definitions for the style of curves
type CurveStyle struct {
	// lines
	LineColor string  `json:"lineColor"` // color
	LineAlpha float64 `json:"lineAlpha"` // alpha (0, 1]. A<1e-14 => A=1.0
	LineStyle string  `json:"lineStyle"` // style
	LineWidth float64 `json:"lineWidth"` // width

	// markers
	MarkerType      string  `json:"markerType"`      // type
	MarkerImg       string  `json:"markerImg"`       // image filename
	MarkerColor     string  `json:"markerColor"`     // color
	MarkerAlpha     float64 `json:"markerAlpha"`     // alpha (0, 1]
	MarkerSize      float64 `json:"markerSize"`      // size; when using images, set markerSize=0 to use the image width
	MarkerEvery     float64 `json:"markerEvery"`     // mark-every
	MarkerLineColor string  `json:"markerLineColor"` // edge color
	MarkerLineWidth float64 `json:"markerLineWidth"` // edge width
	MarkerLineStyle string  `json:"markerLineStyle"` // edge style
	MarkerIsVoid    bool    `json:"markerIsVoid"`    // void marker (draw edge only)
}

// DefaultCurveStyle defines the default style
var DefaultCurveStyle = CurveStyle{
	// lines
	LineColor: "#b33434",
	LineAlpha: 0.7,
	LineStyle: "-",
	LineWidth: 3,

	// markers
	MarkerType:      "o",
	MarkerImg:       "",
	MarkerColor:     "#4c4deb",
	MarkerAlpha:     1,
	MarkerSize:      0,
	MarkerEvery:     0,
	MarkerLineColor: "#ffffff",
	MarkerLineWidth: 2,
	MarkerLineStyle: "none",
	MarkerIsVoid:    false,
}

// PythonParams returns curve style as python options
func (o CurveStyle) PythonParams() (l string) {
	// lines
	if o.LineColor != "" {
		l += io.Sf(",color='%s'", o.LineColor)
	}
	if o.LineAlpha > 0 {
		l += io.Sf(",alpha=%g", o.LineAlpha)
	}
	if o.LineStyle != "" {
		l += io.Sf(",linestyle='%s'", o.LineStyle)
	}
	if o.LineWidth > 0 {
		l += io.Sf(",lw=%g", o.LineWidth)
	}

	// markers
	if o.MarkerType != "" {
		l += io.Sf(",marker='%s'", o.MarkerType)
	}
	if o.MarkerSize > 0 {
		l += io.Sf(",ms=%d", o.MarkerSize)
	}
	if o.MarkerEvery > 0 {
		l += io.Sf(",markevery=%d", o.MarkerEvery)
	}
	if o.MarkerLineColor != "" {
		l += io.Sf(",markeredgecolor='%s'", o.MarkerLineColor)
	}
	if o.MarkerLineWidth > 0 {
		l += io.Sf(",mew=%g", o.MarkerLineWidth)
	}
	if o.MarkerIsVoid {
		l += ",markerfacecolor='none'"
		if o.MarkerLineColor == "" {
			l += io.Sf(",markeredgecolor='%s'", o.LineColor)
		}
	}
	return
}
