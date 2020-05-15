// Copyright 2016 The Gosl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package plt

import (
	"testing"

	"github.com/cpmech/gosl/chk"
)

func Test_JserverPlotSimpleCurve(tst *testing.T) {

	// title of test
	verbose()
	chk.PrintTitle("JserverPlotSimpleCurve")

	// plot curves
	Init()
	Plot([]float64{0, 1, 2, 3}, []float64{0, 1, 2, 3})
	Plot([]float64{0, 1, 2, 3}, []float64{0, 1, 4, 9}, ParamsPlot{Marker: "*"})
}
