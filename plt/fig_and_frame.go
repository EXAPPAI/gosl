// Copyright 2016 The Gosl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package plt

import "github.com/cpmech/gosl/io"

// HideBorders hides frame borders. "params" is optional (only one is considered)
func HideBorders(params ...ParamsFrame) {
	if len(params) < 1 {
		return
	}
	if !paramsControl.UsePython {
		return
	}
	hide := pythonGetHideList(params[0])
	if hide != "" {
		io.Ff(&pythonBuffer, "for spine in %s: plt.gca().spines[spine].set_visible(0)\n", hide)
	}
}

// HideAllBorders hides all frame borders
func HideAllBorders() {
	if !paramsControl.UsePython {
		return
	}
	io.Ff(&pythonBuffer, "for spine in ['left','right','bottom','top']: plt.gca().spines[spine].set_visible(0)\n")
}

// HideTRborders hides top and right borders
func HideTRborders() {
	if !paramsControl.UsePython {
		return
	}
	io.Ff(&pythonBuffer, "for spine in ['right','top']: plt.gca().spines[spine].set_visible(0)\n")
}

// SupTitle sets subplot title
func SupTitle(txt string, params ...ParamsText) {
	if !paramsControl.UsePython {
		return
	}
	uid := pythonGenUID()
	io.Ff(&pythonBuffer, "st%d = plt.suptitle(r'%s'", uid, txt)
	// updateBufferAndClose(&pythonBuffer, params, false, false)
	io.Ff(&pythonBuffer, "addToEA(st%d)\n", uid)
}

// Title sets title
func Title(txt string, params ...ParamsText) {
	if !paramsControl.UsePython {
		return
	}
	io.Ff(&pythonBuffer, "plt.title(r'%s'", txt)
	// updateBufferAndClose(&pythonBuffer, params, false, false)
}
