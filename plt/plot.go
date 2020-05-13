// Copyright 2016 The Gosl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package plt

import "github.com/cpmech/gosl/io"

// Plot plots x-y series
func Plot(x, y []float64, params ...ParamsPlot) (sx, sy string) {
	if !paramsControl.UsePython {
		return
	}
	uid := pythonGenUID()
	sx = io.Sf("x%d", uid)
	sy = io.Sf("y%d", uid)
	pythonGen2Arrays(&pythonBuffer, sx, sy, x, y)
	io.Ff(&pythonBuffer, "plt.plot(%s,%s", sx, sy)
	pythonPutParams(params)
	return
}

// PlotOne plots one point @ (x,y)
func PlotOne(x, y float64, params ...ParamsPlot) {
	if !paramsControl.UsePython {
		return
	}
	io.Ff(&pythonBuffer, "plt.plot(%23.15e,%23.15e", x, y)
	// updateBufferAndClose(&pythonBuffer, params, false, false)
}
