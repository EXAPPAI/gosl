// Copyright 2016 The Gosl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package plt

import (
	"testing"
	"time"

	"github.com/cpmech/gosl/chk"
)

func Test_PlotSimpleCurve(tst *testing.T) {

	// title of test
	verbose()
	chk.PrintTitle("PlotSimpleCurve")

	x1 := []float64{0, 1, 2, 3}
	y1 := []float64{0, 1, 2, 3}
	y2 := []float64{0, 1, 4, 9}

	Begin("myplot1")
	Plot(x1, y1, "curve1")
	Plot(x1, y2, "curve2", CurveStyle{MarkerType: "*"})

	time.Sleep(3000 * time.Millisecond)

	Begin("myplot2")
	Plot(x1, y1, "curveX")
	Plot(x1, y2, "curveY", CurveStyle{MarkerType: "*"})
}
