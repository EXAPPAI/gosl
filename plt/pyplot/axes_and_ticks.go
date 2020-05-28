// Copyright 2016 The Gosl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package plt

/*
// DoubleYscale duplicates y-scale
func DoubleYscale(ylabelOrEmpty string) {
	if !argsControl.UsePython {
		return
	}
	io.Ff(&pythonBuffer, "plt.gca().twinx()\n")
	if ylabelOrEmpty != "" {
		io.Ff(&pythonBuffer, "plt.gca().set_ylabel('%s')\n", ylabelOrEmpty)
	}
}

// SetXlog sets x-scale to be log
func SetXlog() {
	if !argsControl.UsePython {
		return
	}
	io.Ff(&pythonBuffer, "plt.gca().set_xscale('log')\n")
}

// SetYlog sets y-scale to be log
func SetYlog() {
	if !argsControl.UsePython {
		return
	}
	io.Ff(&pythonBuffer, "plt.gca().set_yscale('log')\n")
}

// SetNoXtickLabels hides labels of x-ticks but keep ticks
func SetNoXtickLabels() {
	if !argsControl.UsePython {
		return
	}
	io.Ff(&pythonBuffer, "plt.gca().tick_params(labelbottom='off')\n")
}

// SetNoYtickLabels hides labels of y-ticks but keep ticks
func SetNoYtickLabels() {
	if !argsControl.UsePython {
		return
	}
	io.Ff(&pythonBuffer, "plt.gca().tick_params(labelleft='off')\n")
}

// SetTicksXlist sets x-axis ticks with given list
func SetTicksXlist(values []float64) {
	if !argsControl.UsePython {
		return
	}
	io.Ff(&pythonBuffer, "plt.xticks(%v)\n", floats2list(values))
}

// SetTicksYlist sets y-ayis ticks with given list
func SetTicksYlist(values []float64) {
	if !argsControl.UsePython {
		return
	}
	io.Ff(&pythonBuffer, "plt.yticks(%v)\n", floats2list(values))
}

// SetXnticks sets number of ticks along x
func SetXnticks(num int) {
	if !argsControl.UsePython {
		return
	}
	if num == 0 {
		io.Ff(&pythonBuffer, "plt.gca().get_xaxis().set_ticks([])\n")
	} else {
		io.Ff(&pythonBuffer, "plt.gca().get_xaxis().set_major_locator(tck.MaxNLocator(%d))\n", num)
	}
}

// SetYnticks sets number of ticks along y
func SetYnticks(num int) {
	if !argsControl.UsePython {
		return
	}
	if num == 0 {
		io.Ff(&pythonBuffer, "plt.gca().get_yaxis().set_ticks([])\n")
	} else {
		io.Ff(&pythonBuffer, "plt.gca().get_yaxis().set_major_locator(tck.MaxNLocator(%d))\n", num)
	}
}

// SetTicksX sets ticks along x
func SetTicksX(majorEvery, minorEvery float64, majorFmt string) {
	if !argsControl.UsePython {
		return
	}
	uid := pythonGenUID()
	if majorEvery > 0 {
		io.Ff(&pythonBuffer, "majorLocator%d = tck.MultipleLocator(%g)\n", uid, majorEvery)
		io.Ff(&pythonBuffer, "nticks%d = (plt.gca().axis()[1] - plt.gca().axis()[0]) / %g\n", uid, majorEvery)
		io.Ff(&pythonBuffer, "if nticks%d < majorLocator%d.MAXTICKS * 0.9:\n", uid, uid)
		io.Ff(&pythonBuffer, "    plt.gca().xaxis.set_major_locator(majorLocator%d)\n", uid)
	}
	if minorEvery > 0 {
		io.Ff(&pythonBuffer, "minorLocator%d = tck.MultipleLocator(%g)\n", uid, minorEvery)
		io.Ff(&pythonBuffer, "nticks%d = (plt.gca().axis()[1] - plt.gca().axis()[0]) / %g\n", uid, minorEvery)
		io.Ff(&pythonBuffer, "if nticks%d < minorLocator%d.MAXTICKS * 0.9:\n", uid, uid)
		io.Ff(&pythonBuffer, "    plt.gca().xaxis.set_minor_locator(minorLocator%d)\n", uid)
	}
	if majorFmt != "" {
		io.Ff(&pythonBuffer, "majorFormatter%d = tck.FormatStrFormatter(r'%s')\n", uid, majorFmt)
		io.Ff(&pythonBuffer, "plt.gca().xaxis.set_major_formatter(majorFormatter%d)\n", uid)
	}
}

// SetTicksY sets ticks along y
func SetTicksY(majorEvery, minorEvery float64, majorFmt string) {
	if !argsControl.UsePython {
		return
	}
	uid := pythonGenUID()
	if majorEvery > 0 {
		io.Ff(&pythonBuffer, "majorLocator%d = tck.MultipleLocator(%g)\n", uid, majorEvery)
		io.Ff(&pythonBuffer, "nticks%d = (plt.gca().axis()[1] - plt.gca().axis()[0]) / %g\n", uid, majorEvery)
		io.Ff(&pythonBuffer, "if nticks%d < majorLocator%d.MAXTICKS * 0.9:\n", uid, uid)
		io.Ff(&pythonBuffer, "    plt.gca().yaxis.set_major_locator(majorLocator%d)\n", uid)
	}
	if minorEvery > 0 {
		io.Ff(&pythonBuffer, "minorLocator%d = tck.MultipleLocator(%g)\n", uid, minorEvery)
		io.Ff(&pythonBuffer, "nticks%d = (plt.gca().axis()[1] - plt.gca().axis()[0]) / %g\n", uid, minorEvery)
		io.Ff(&pythonBuffer, "if nticks%d < minorLocator%d.MAXTICKS * 0.9:\n", uid, uid)
		io.Ff(&pythonBuffer, "    plt.gca().yaxis.set_minor_locator(minorLocator%d)\n", uid)
	}
	if majorFmt != "" {
		io.Ff(&pythonBuffer, "majorFormatter%d = tck.FormatStrFormatter(r'%s')\n", uid, majorFmt)
		io.Ff(&pythonBuffer, "plt.gca().yaxis.set_major_formatter(majorFormatter%d)\n", uid)
	}
}

// SetScientificX sets scientific notation for ticks along x-axis
func SetScientificX(minOrder, maxOrder int) {
	if !argsControl.UsePython {
		return
	}
	uid := pythonGenUID()
	io.Ff(&pythonBuffer, "fmt%d = plt.ScalarFormatter(useOffset=True)\n", uid)
	io.Ff(&pythonBuffer, "fmt%d.set_powerlimits((%d,%d))\n", uid, minOrder, maxOrder)
	io.Ff(&pythonBuffer, "plt.gca().xaxis.set_major_formatter(fmt%d)\n", uid)
}

// SetScientificY sets scientific notation for ticks along y-axis
func SetScientificY(minOrder, maxOrder int) {
	if !argsControl.UsePython {
		return
	}
	uid := pythonGenUID()
	io.Ff(&pythonBuffer, "fmt%d = plt.ScalarFormatter(useOffset=True)\n", uid)
	io.Ff(&pythonBuffer, "fmt%d.set_powerlimits((%d,%d))\n", uid, minOrder, maxOrder)
	io.Ff(&pythonBuffer, "plt.gca().yaxis.set_major_formatter(fmt%d)\n", uid)
}

// SetTicksNormal sets normal ticks
func SetTicksNormal() {
	if !argsControl.UsePython {
		return
	}
	io.Ff(&pythonBuffer, "plt.gca().ticklabel_format(useOffset=False)\n")
}

// SetTicksRotationX sets the rotation angle of x-ticks
func SetTicksRotationX(degree float64) {
	if !argsControl.UsePython {
		return
	}
	io.Ff(&pythonBuffer, "plt.setp(plt.gca().get_xticklabels(), rotation=%g)\n", degree)
}

// SetTicksRotationY sets the rotation angle of y-ticks
func SetTicksRotationY(degree float64) {
	if !argsControl.UsePython {
		return
	}
	io.Ff(&pythonBuffer, "plt.setp(plt.gca().get_yticklabels(), rotation=%g)\n", degree)
}

// ReplaceAxes substitutes axis frame (see Axes in gosl.py)
//   ex: xDel, yDel := 0.04, 0.04
func ReplaceAxes(xi, yi, xf, yf, xDel, yDel float64, xLab, yLab string, paramsArrow ParamsPlot, paramsText ParamsText) {
	if !argsControl.UsePython {
		return
	}
	io.Ff(&pythonBuffer, "plt.axis('off')\n")
	Arrow(xi, yi, xf, yi, paramsArrow)
	Arrow(xi, yi, xi, yf, paramsArrow)
	Text(xf, yi-xDel, xLab, paramsText)
	Text(xi-yDel, yf, yLab, paramsText)
}

// AxHline adds horizontal line to axis
func AxHline(y float64, params ParamsPlot) {
	if !argsControl.UsePython {
		return
	}
	io.Ff(&pythonBuffer, "plt.axhline(%g", y)
	// updateBufferAndClose(&pythonBuffer, args, false, false)
}

// AxVline adds vertical line to axis
func AxVline(x float64, params ParamsPlot) {
	if !argsControl.UsePython {
		return
	}
	io.Ff(&pythonBuffer, "plt.axvline(%g", x)
	// updateBufferAndClose(&pythonBuffer, args, false, false)
}
*/
