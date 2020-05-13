// Copyright 2016 The Gosl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package plt

import (
	"testing"

	"github.com/cpmech/gosl/chk"
)

func Test_InitDefault(tst *testing.T) {

	verbose()
	chk.PrintTitle("InitDefault")

	Init()
	if paramsControl.UsePython {
		tst.Errorf("UsePython should be false")
	}
	chk.String(tst, paramsControl.TmpDir, "/tmp/gosl/plt")
	chk.String(tst, paramsControl.OutDir, "/tmp/gosl/plt")
	if paramsControl.FigEps {
		tst.Errorf("FigEps should be false")
	}
	chk.Int(tst, "FigDpi", paramsControl.FigDpi, 150)
	chk.Float64(tst, "FigProp", 1e-10, paramsControl.FigProp, 0.75)
	chk.Float64(tst, "FigWidth", 1e-10, paramsControl.FigWidth, 400)
}

func Test_InitWithParams(tst *testing.T) {

	verbose()
	chk.PrintTitle("InitWithParams")

	Init(ParamsControl{UsePython: true})
	if !paramsControl.UsePython {
		tst.Errorf("UsePython should be true")
	}
	chk.String(tst, paramsControl.TmpDir, "/tmp/gosl/plt")
	chk.String(tst, paramsControl.OutDir, "/tmp/gosl/plt")
	if paramsControl.FigEps {
		tst.Errorf("FigEps should be false")
	}
	chk.Int(tst, "FigDpi", paramsControl.FigDpi, 150)
	chk.Float64(tst, "FigProp", 1e-10, paramsControl.FigProp, 0.75)
	chk.Float64(tst, "FigWidth", 1e-10, paramsControl.FigWidth, 400)
}
