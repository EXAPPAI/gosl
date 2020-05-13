// Copyright 2016 The Gosl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package plt

import (
	"testing"

	"github.com/cpmech/gosl/chk"
)

func Test_HideBorders(tst *testing.T) {

	verbose()
	chk.PrintTitle("HideBorders")

	Init(ParamsControl{UsePython: true})
	Save("HideBorders")
}
