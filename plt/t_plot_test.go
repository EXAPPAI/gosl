// Copyright 2016 The Gosl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package plt

import (
	"math"
	"testing"

	"github.com/cpmech/gosl/utl"

	"github.com/cpmech/gosl/chk"
)

func Test_PlotSimpleCurve(tst *testing.T) {

	// title of test
	verbose()
	chk.PrintTitle("PlotSimpleCurve")

	x1 := utl.LinSpace(0, 2*math.Pi, 101)
	y1 := utl.GetMapped(x1, func(x float64) float64 { return math.Sin(x) })
	y2 := utl.GetMapped(x1, func(x float64) float64 { return math.Cos(x) })
	y3 := utl.GetMapped(x1, func(x float64) float64 { return math.Sin(4 * x) })
	y4 := utl.GetMapped(x1, func(x float64) float64 { return math.Cos(4 * x) })

	style1 := NewCurveStyle()
	style1.MarkerType = "+"
	style1.MarkerEvery = 5

	style2 := NewCurveStyle()
	style2.LineColor = "#f3d144"
	style2.MarkerColor = "#d555fb"
	style2.MarkerEvery = 2

	Begin("myplot1")
	Plot(x1, y1, "curve1", style1)
	// time.Sleep(6000 * time.Millisecond)
	Plot(x1, y2, "curve2", style2)

	// time.Sleep(3000 * time.Millisecond)

	Begin("myplot2")
	Plot(x1, y3, "curveX", style1)
	Plot(x1, y4, "curveY", style2)
}
