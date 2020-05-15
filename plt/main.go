// Copyright 2016 The Gosl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package plt

// default control arguments
var paramsControl = &ParamsControl{
	UsePython: false,
	TmpDir:    "/tmp/gosl/plt",
	OutDir:    "/tmp/gosl/plt",
	FigEps:    false,
	FigDpi:    150,
	FigProp:   0.75,
	FigWidth:  400,
}

// flag indicating that the system has been initialized
var initialized = false

// Init initializes the system. "params" is optional (only one is considered)
func Init(params ...ParamsControl) {
	if len(params) > 0 {
		p := &params[0]
		paramsControl.UsePython = p.UsePython
		paramsControl.FigEps = p.FigEps
		if p.TmpDir != "" {
			paramsControl.TmpDir = p.TmpDir
		}
		if p.OutDir != "" {
			paramsControl.OutDir = p.OutDir
		}
		if p.FigDpi > 0 {
			paramsControl.FigDpi = p.FigDpi
		}
		if p.FigProp > 0 {
			paramsControl.FigProp = p.FigProp
		}
		if p.FigWidth > 0 {
			paramsControl.FigWidth = p.FigWidth
		}
	}
	if paramsControl.UsePython {
		pythonInit()
	} else {
		jserverInit()
	}
	initialized = true
}

// Save saves figure
func Save(fnkey string) {
	if paramsControl.UsePython {
		pythonSave(fnkey, false, false)
		return
	}
}

/*
// Annotate adds annotation to plot
func Annotate(x, y float64, txt string, params *A) {
	if !paramsControl.UsePython {
		return
	}
	l := "plt.annotate(r'%s',xy=(%g,%g)"
	if params != nil {
		addToCmd(&l, params.AxCoords, "xycoords='axes fraction'")
		addToCmd(&l, params.FigCoords, "xycoords='figure fraction'")
	}
	io.Ff(&pythonBuffer, l, txt, x, y)
	updateBufferAndClose(&pythonBuffer, params, false, false)
}

// AnnotateXlabels sets text of xlabels
func AnnotateXlabels(x float64, txt string, params *A) {
	if !paramsControl.UsePython {
		return
	}
	fsz := 7.0
	if params != nil {
		if params.Fsz > 0 {
			fsz = params.Fsz
		}
	}
	io.Ff(&pythonBuffer, "plt.annotate('%s', xy=(%g, -%g-3), xycoords=('data', 'axes points'), va='top', ha='center', size=%g", txt, x, fsz, fsz)
	updateBufferAndClose(&pythonBuffer, params, false, false)
}


// Text adds text to plot
func Text(x, y float64, txt string, params *A) {
	if !paramsControl.UsePython {
		return
	}
	io.Ff(&pythonBuffer, "plt.text(%g,%g,r'%s'", x, y, txt)
	updateBufferAndClose(&pythonBuffer, params, false, false)
}

// Cross adds a vertical and horizontal lines @ (x0,y0) to plot (i.e. large cross)
func Cross(x0, y0 float64, params *A) {
	if !paramsControl.UsePython {
		return
	}
	cl, ls, lw, z := "black", "dashed", 1.2, 0
	if params != nil {
		if params.C != "" {
			cl = params.C
		}
		if params.Lw > 0 {
			lw = params.Lw
		}
		if params.Ls != "" {
			ls = params.Ls
		}
		if params.Z > 0 {
			z = params.Z
		}
	}
	io.Ff(&pythonBuffer, "plt.axvline(%g, color='%s', linestyle='%s', linewidth=%g, zorder=%d)\n", x0, cl, ls, lw, z)
	io.Ff(&pythonBuffer, "plt.axhline(%g, color='%s', linestyle='%s', linewidth=%g, zorder=%d)\n", y0, cl, ls, lw, z)
}

// SplotGap sets gap between subplots
func SplotGap(w, h float64) {
	if !paramsControl.UsePython {
		return
	}
	io.Ff(&pythonBuffer, "plt.subplots_adjust(wspace=%g, hspace=%g)\n", w, h)
}

// Subplot adds/sets a subplot
func Subplot(i, j, k int) {
	if !paramsControl.UsePython {
		return
	}
	io.Ff(&pythonBuffer, "plt.subplot(%d,%d,%d)\n", i, j, k)
}

// SubplotI adds/sets a subplot with given indices in I
func SubplotI(I []int) {
	if !paramsControl.UsePython {
		return
	}
	if len(I) != 3 {
		return
	}
	io.Ff(&pythonBuffer, "plt.subplot(%d,%d,%d)\n", I[0], I[1], I[2])
}

// SetHspace sets horizontal space between subplots
func SetHspace(hspace float64) {
	if !paramsControl.UsePython {
		return
	}
	io.Ff(&pythonBuffer, "plt.subplots_adjust(hspace=%g)\n", hspace)
}

// SetVspace sets vertical space between subplots
func SetVspace(vspace float64) {
	if !paramsControl.UsePython {
		return
	}
	io.Ff(&pythonBuffer, "plt.subplots_adjust(vspace=%g)\n", vspace)
}

// Equal sets same scale for both axes
func Equal() {
	if !paramsControl.UsePython {
		return
	}
	io.Ff(&pythonBuffer, "plt.axis('equal')\n")
}

// AxisOff hides axes
func AxisOff() {
	if !paramsControl.UsePython {
		return
	}
	io.Ff(&pythonBuffer, "plt.axis('off')\n")
}

// SetAxis sets axes limits
func SetAxis(xmin, xmax, ymin, ymax float64) {
	if !paramsControl.UsePython {
		return
	}
	io.Ff(&pythonBuffer, "plt.axis([%g, %g, %g, %g])\n", xmin, xmax, ymin, ymax)
}

// AxisXmin sets minimum x
func AxisXmin(xmin float64) {
	if !paramsControl.UsePython {
		return
	}
	io.Ff(&pythonBuffer, "plt.axis([%g, plt.axis()[1], plt.axis()[2], plt.axis()[3]])\n", xmin)
}

// AxisXmax sets maximum x
func AxisXmax(xmax float64) {
	if !paramsControl.UsePython {
		return
	}
	io.Ff(&pythonBuffer, "plt.axis([plt.axis()[0], %g, plt.axis()[2], plt.axis()[3]])\n", xmax)
}

// AxisYmin sets minimum y
func AxisYmin(ymin float64) {
	if !paramsControl.UsePython {
		return
	}
	io.Ff(&pythonBuffer, "plt.axis([plt.axis()[0], plt.axis()[1], %g, plt.axis()[3]])\n", ymin)
}

// AxisYmax sets maximum y
func AxisYmax(ymax float64) {
	if !paramsControl.UsePython {
		return
	}
	io.Ff(&pythonBuffer, "plt.axis([plt.axis()[0], plt.axis()[1], plt.axis()[2], %g])\n", ymax)
}

// AxisXrange sets x-range (i.e. limits)
func AxisXrange(xmin, xmax float64) {
	if !paramsControl.UsePython {
		return
	}
	io.Ff(&pythonBuffer, "plt.axis([%g, %g, plt.axis()[2], plt.axis()[3]])\n", xmin, xmax)
}

// AxisYrange sets y-range (i.e. limits)
func AxisYrange(ymin, ymax float64) {
	if !paramsControl.UsePython {
		return
	}
	io.Ff(&pythonBuffer, "plt.axis([plt.axis()[0], plt.axis()[1], %g, %g])\n", ymin, ymax)
}

// AxisRange sets x and y ranges (i.e. limits)
func AxisRange(xmin, xmax, ymin, ymax float64) {
	if !paramsControl.UsePython {
		return
	}
	io.Ff(&pythonBuffer, "plt.axis([%g, %g, %g, %g])\n", xmin, xmax, ymin, ymax)
}

// AxisLims sets x and y limits
func AxisLims(lims []float64) {
	if !paramsControl.UsePython {
		return
	}
	io.Ff(&pythonBuffer, "plt.axis([%g, %g, %g, %g])\n", lims[0], lims[1], lims[2], lims[3])
}


// Hist draws histogram
func Hist(x [][]float64, labels []string, params *A) {
	if !paramsControl.UsePython {
		return
	}
	uid := pythonGenUID()
	sx := io.Sf("x%d", uid)
	sy := io.Sf("y%d", uid)
	genList(&pythonBuffer, sx, x)
	genStrArray(&pythonBuffer, sy, labels)
	io.Ff(&pythonBuffer, "plt.hist(%s,label=r'%s'", sx, sy)
	updateBufferAndClose(&pythonBuffer, params, true, false)
}

// Grid2d draws grid lines of 2D grid
//   withIDs -- add text with IDs numbered by looping over {X,Y}[j][i] (j:outer, i:inner)
func Grid2d(X, Y [][]float64, withIDs bool, paramsLines, paramsIDs *A) {
	if !paramsControl.UsePython {
		return
	}
	if len(X) < 2 || len(Y) < 2 {
		return
	}
	if paramsLines == nil {
		paramsLines = &A{C: "#427ce5", Lw: 0.8, NoClip: true}
	}
	if paramsIDs == nil {
		paramsIDs = &A{C: C(2, 0), Fsz: 6}
	}
	n1 := len(X)
	n0 := len(X[0])
	xx := make([]float64, 2) // min,max
	yy := make([]float64, 2) // min,max
	idx := 0
	for j := 0; j < n1; j++ {
		for i := 0; i < n0; i++ {
			if i > 0 {
				xx[0], xx[1] = X[j][i-1], X[j][i]
				yy[0], yy[1] = Y[j][i-1], Y[j][i]
				Plot(xx, yy, paramsLines)
			}
			if j > 0 {
				xx[0], xx[1] = X[j-1][i], X[j][i]
				yy[0], yy[1] = Y[j-1][i], Y[j][i]
				Plot(xx, yy, paramsLines)
			}
			if withIDs {
				Text(X[j][i], Y[j][i], io.Sf("%d", idx), paramsIDs)
				idx++
			}
		}
	}
}

// ContourF draws filled contour and possibly with a contour of lines (if params.UnoLines=false)
func ContourF(x, y, z [][]float64, params *A) {
	if !paramsControl.UsePython {
		return
	}
	uid := pythonGenUID()
	sx := io.Sf("x%d", uid)
	sy := io.Sf("y%d", uid)
	sz := io.Sf("z%d", uid)
	genMat(&pythonBuffer, sx, x)
	genMat(&pythonBuffer, sy, y)
	genMat(&pythonBuffer, sz, z)
	a, colors, levels := paramsContour(params, z)
	io.Ff(&pythonBuffer, "c%d = plt.contourf(%s,%s,%s%s%s)\n", uid, sx, sy, sz, colors, levels)
	if !a.NoLines {
		io.Ff(&pythonBuffer, "cc%d = plt.contour(%s,%s,%s,colors=['k']%s,linewidths=[%g])\n", uid, sx, sy, sz, levels, a.Lw)
		if !a.NoLabels {
			io.Ff(&pythonBuffer, "plt.clabel(cc%d,inline=%d,fontsize=%g)\n", uid, pyBool(!a.NoInline), a.Fsz)
		}
	}
	if !a.NoCbar {
		io.Ff(&pythonBuffer, "cb%d = plt.colorbar(c%d, format='%s')\n", uid, uid, a.NumFmt)
		if a.CbarLbl != "" {
			io.Ff(&pythonBuffer, "cb%d.ax.set_ylabel(r'%s')\n", uid, a.CbarLbl)
		}
	}
	if a.SelectC != "" {
		io.Ff(&pythonBuffer, "ccc%d = plt.contour(%s,%s,%s,colors=['%s'],levels=[%g],linewidths=[%g],linestyles=['-'])\n", uid, sx, sy, sz, a.SelectC, a.SelectV, a.SelectLw)
	}
}

// ContourL draws a contour with lines only
func ContourL(x, y, z [][]float64, params *A) {
	if !paramsControl.UsePython {
		return
	}
	uid := pythonGenUID()
	sx := io.Sf("x%d", uid)
	sy := io.Sf("y%d", uid)
	sz := io.Sf("z%d", uid)
	genMat(&pythonBuffer, sx, x)
	genMat(&pythonBuffer, sy, y)
	genMat(&pythonBuffer, sz, z)
	a, colors, levels := paramsContour(params, z)
	io.Ff(&pythonBuffer, "c%d = plt.contour(%s,%s,%s%s%s)\n", uid, sx, sy, sz, colors, levels)
	if !a.NoLabels {
		io.Ff(&pythonBuffer, "plt.clabel(c%d,inline=%d,fontsize=%g)\n", uid, pyBool(!a.NoInline), a.Fsz)
	}
	if a.SelectC != "" {
		io.Ff(&pythonBuffer, "cc%d = plt.contour(%s,%s,%s,colors=['%s'],levels=[%g],linewidths=[%g],linestyles=['-'])\n", uid, sx, sy, sz, a.SelectC, a.SelectV, a.SelectLw)
	}
}

// Quiver draws vector field
func Quiver(x, y, gx, gy [][]float64, params *A) {
	if !paramsControl.UsePython {
		return
	}
	uid := pythonGenUID()
	sx := io.Sf("x%d", uid)
	sy := io.Sf("y%d", uid)
	sgx := io.Sf("gx%d", uid)
	sgy := io.Sf("gy%d", uid)
	genMat(&pythonBuffer, sx, x)
	genMat(&pythonBuffer, sy, y)
	genMat(&pythonBuffer, sgx, gx)
	genMat(&pythonBuffer, sgy, gy)
	io.Ff(&pythonBuffer, "plt.quiver(%s,%s,%s,%s", sx, sy, sgx, sgy)
	if params != nil {
		if params.Scale > 0 {
			io.Ff(&pythonBuffer, ",scale=%g", params.Scale)
		}
	}
	updateBufferAndClose(&pythonBuffer, params, false, false)
}

// Grid adds grid to plot
func Grid(params *A) {
	if !paramsControl.UsePython {
		return
	}
	io.Ff(&pythonBuffer, "plt.grid(")
	updateBufferFirstArgsAndClose(&pythonBuffer, params, false, false)
}

// Legend adds legend to plot
func Legend(params *A) {
	if !paramsControl.UsePython {
		return
	}
	loc, ncol, hlen, fsz, frame, out, outX := paramsLeg(params)
	uid := pythonGenUID()
	io.Ff(&pythonBuffer, "h%d, l%d = plt.gca().get_legend_handles_labels()\n", uid, uid)
	io.Ff(&pythonBuffer, "if len(h%d) > 0 and len(l%d) > 0:\n", uid, uid)
	if out == 1 {
		io.Ff(&pythonBuffer, "    d%d = %s\n", uid, outX)
		io.Ff(&pythonBuffer, "    l%d = plt.legend(bbox_to_anchor=d%d, ncol=%d, handlelength=%g, prop={'size':%g}, loc=3, mode='expand', borderaxespad=0.0, columnspacing=1, handletextpad=0.05)\n", uid, uid, ncol, hlen, fsz)
		io.Ff(&pythonBuffer, "    addToEA(l%d)\n", uid)
	} else {
		io.Ff(&pythonBuffer, "    l%d = plt.legend(loc=%s, ncol=%d, handlelength=%g, prop={'size':%g})\n", uid, loc, ncol, hlen, fsz)
		io.Ff(&pythonBuffer, "    addToEA(l%d)\n", uid)
	}
	if frame == 0 {
		io.Ff(&pythonBuffer, "    l%d.get_frame().set_linewidth(0.0)\n", uid)
	}
}

// LegendX adds legend to plot with given data instead of relying on labels
func LegendX(dat []*A, params *A) {
	if !paramsControl.UsePython {
		return
	}
	loc, ncol, hlen, fsz, frame, out, outX := paramsLeg(params)
	uid := pythonGenUID()
	io.Ff(&pythonBuffer, "h%d = [", uid)
	for i, d := range dat {
		if i > 0 {
			io.Ff(&pythonBuffer, ",\n")
		}
		if d != nil {
			io.Ff(&pythonBuffer, "lns.Line2D([], [], %s)", d.String(false, false))
		}
	}
	io.Ff(&pythonBuffer, "]\n")
	io.Ff(&pythonBuffer, "if len(h%d) > 0:\n", uid)
	if out == 1 {
		io.Ff(&pythonBuffer, "    d%d = %s\n", uid, outX)
		io.Ff(&pythonBuffer, "    l%d = plt.legend(handles=h%d, bbox_to_anchor=d%d, ncol=%d, handlelength=%g, prop={'size':%g}, loc=3, mode='expand', borderaxespad=0.0, columnspacing=1, handletextpad=0.05)\n", uid, uid, uid, ncol, hlen, fsz)
		io.Ff(&pythonBuffer, "    addToEA(l%d)\n", uid)
	} else {
		io.Ff(&pythonBuffer, "    l%d = plt.legend(handles=h%d, loc=%s, ncol=%d, handlelength=%g, prop={'size':%g})\n", uid, uid, loc, ncol, hlen, fsz)
		io.Ff(&pythonBuffer, "    addToEA(l%d)\n", uid)
	}
	if frame == 0 {
		io.Ff(&pythonBuffer, "    l%d.get_frame().set_linewidth(0.0)\n", uid)
	}
}

// Gll adds grid, labels, and legend to plot
func Gll(xl, yl string, params *A) {
	if !paramsControl.UsePython {
		return
	}
	hide := getHideList(params)
	if hide != "" {
		io.Ff(&pythonBuffer, "for spine in %s: plt.gca().spines[spine].set_visible(0)\n", hide)
	}
	io.Ff(&pythonBuffer, "plt.grid(color='grey', zorder=-1000)\n")
	io.Ff(&pythonBuffer, "plt.xlabel(r'%s')\n", xl)
	io.Ff(&pythonBuffer, "plt.ylabel(r'%s')\n", yl)
	Legend(params)
}

// SetLabels sets x-y axes labels
func SetLabels(x, y string, params *A) {
	if !paramsControl.UsePython {
		return
	}
	a := ""
	if params != nil {
		a = "," + params.String(false, false)
	}
	io.Ff(&pythonBuffer, "plt.xlabel(r'%s'%s);plt.ylabel(r'%s'%s)\n", x, a, y, a)
}

// SetXlabel sets x-label
func SetXlabel(xl string, params *A) {
	if !paramsControl.UsePython {
		return
	}
	io.Ff(&pythonBuffer, "plt.xlabel(r'%s')\n", xl)
}

// SetYlabel sets y-label
func SetYlabel(yl string, params *A) {
	if !paramsControl.UsePython {
		return
	}
	io.Ff(&pythonBuffer, "plt.ylabel(r'%s')\n", yl)
}

// Clf clears current figure
func Clf() {
	if !paramsControl.UsePython {
		return
	}
	io.Ff(&pythonBuffer, "plt.clf()\n")
}

// SetFontSizes sets font sizes
//   NOTE: this function also sets the FontSet, if not ""
func SetFontSizes(params *A) {
	if !paramsControl.UsePython {
		return
	}
	txt, lbl, leg, xtck, ytck, fontset := paramsFsz(params)
	io.Ff(&pythonBuffer, "plt.rcParams.update({\n")
	io.Ff(&pythonBuffer, "    'font.size'       : %g,\n", txt)
	io.Ff(&pythonBuffer, "    'axes.labelsize'  : %g,\n", lbl)
	io.Ff(&pythonBuffer, "    'legend.fontsize' : %g,\n", leg)
	io.Ff(&pythonBuffer, "    'xtick.labelsize' : %g,\n", xtck)
	io.Ff(&pythonBuffer, "    'ytick.labelsize' : %g,\n", ytck)
	io.Ff(&pythonBuffer, "    'mathtext.fontset': '%s'})\n", fontset)
}

// ZoomWindow adds another axes to plot a figure within the figure; e.g. a zoom window
//  lef, bot, wid, hei -- normalised figure coordinates: left,bottom,width,height
//  asOld -- handle to the previous axes
//  axNew -- handle to the new axes
func ZoomWindow(lef, bot, wid, hei float64, params *A) (axOld, axNew string) {
	if !paramsControl.UsePython {
		return
	}
	uid := pythonGenUID()
	clr := "#dcdcdc"
	if params != nil {
		clr = params.C
	}
	axOld = io.Sf("axOld%d", uid)
	io.Ff(&pythonBuffer, "%s = plt.gca()\n", axOld)
	axNew = io.Sf("axNew%d", uid)
	io.Ff(&pythonBuffer, "%s = plt.axes([%g,%g,%g,%g], facecolor='%s')\n", axNew, lef, bot, wid, hei, clr)
	return
}

// Sca sets current axes
func Sca(axName string) {
	if !paramsControl.UsePython {
		return
	}
	io.Ff(&pythonBuffer, "plt.sca(%s)\n", axName)
}

*/
