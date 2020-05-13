// Copyright 2016 The Gosl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package plt

// ParamsControl holds parameters to setup plots
type ParamsControl struct {
	UsePython bool    // use Python instead of JS-server
	TmpDir    string  // temporary dir for gosl.py file. default = "/tmp/gosl/plt"
	OutDir    string  // output directory for figures. default = "/tmp/gosl/plt"
	FigEps    bool    // figure: save EPS file instead of PNG
	FigDpi    int     // figure: dpi to be used when saving figure
	FigProp   float64 // figure: proportion: height = width * prop
	FigWidth  float64 // figure: width. Example: get this from LaTeX using \showthe\columnwidth
}

// ParamsPlot holds parameters to style lines, shapes (e.g. polygons) and text
type ParamsPlot struct {
	// lines
	Color           string  // color
	Alpha           float64 // transparency coefficient
	Marker          string  // marker
	LineStyle       string  // linestyle
	LineWidth       float64 // linewidth; -1 => default
	MarkerSize      int     // marker size; -1 => default
	Label           string  // label
	MarkEvery       int     // mark-every; -1 => default
	Zindex          int     // z-order
	MarkerEdgeColor string  // marker edge color
	MarkerEdgeWidth float64 // marker edge width
	MarkerVoid      bool    // void marker => markeredgecolor='C', markerfacecolor='none'
	NoClip          bool    // turn clipping off

	// shapes
	ShapeFaceColor string  // shapes: face color
	ShapeEdgeColor string  // shapes: edge color
	ShapeScale     float64 // shapes: scale information
	ShapeStyle     string  // shapes: style information
	ShapeClosed    bool    // shapes: closed shape
}

// ParamsText holds parameters to style text and fontsizes
type ParamsText struct {
	// text
	AlignHoriz string  // horizontal alignment; e.g. 'center'
	AlignVert  string  // vertical alignment; e.g. 'center'
	Rotation   float64 // rotation

	// fontsizes
	Size     float64 // font size
	SizeLbl  float64 // font size of labels
	SizeLeg  float64 // font size of legend
	SizeXtck float64 // font size of x-ticks
	SizeYtck float64 // font size of y-ticks
	FontSet  string  // font set: e.g. 'stix', 'stixsans' [default]
}

// ParamsFrame holds parameters to style the figure area
type ParamsFrame struct {
	// frame
	HideLeft   bool // hide left frame border
	HideRight  bool // hide right frame border
	HideBottom bool // hide bottom frame border
	HideTop    bool // hide top frame border

	// other options
	UseAxCoords  bool // the given x-y coordinates correspond to axes coords
	UseFigCoords bool // the given x-y coordinates correspond to figure coords
}

// ParamsLeg holds parameters to style the legend
type ParamsLeg struct {
	Loc   string    // legend: location. e.g.: right, center left, upper right, lower right, best, center, lower left, center right, upper left, upper center, lower center
	Ncol  int       // legend: number of columns
	Hlen  float64   // legend: handle length
	Frame bool      // legend: frame on
	Out   bool      // legend: outside
	OutX  []float64 // legend: normalised coordinates to put legend outside frame
}

// ParamsContour holds parameters for contours
type ParamsContour struct {
	Colors   []string  // contour: colors
	Nlevels  int       // contour: number of levels (overridden by Levels when it's not nil)
	Levels   []float64 // contour: levels (may be nil)
	CmapIdx  int       // contour: colormap index
	NumFmt   string    // contour: number format; e.g. "%g" or "%.2f"
	NoLines  bool      // contour: do not add lines on top of filled contour
	NoLabels bool      // contour: do not add labels
	NoInline bool      // contour: do not draw labels 'inline'
	NoCbar   bool      // contour: do not add colorbar
	CbarLbl  string    // contour: colorbar label
	SelectV  float64   // contour: selected value
	SelectC  string    // contour: color to mark selected level. empty means no selected line
	SelectLw float64   // contour: zero level linewidth
}

// Params3d holds parameters for 3D plots
type Params3d struct {
	Rstride int  // 3d: row stride
	Cstride int  // 3d: column stride
	Surf    bool // 3d: generate surface
	Wire    bool // 3d: generate wireframe
}

// ParamsHist holds parameters for histograms
type ParamsHist struct {
	Colors  []string // histogram: colors
	Type    string   // histogram: type; e.g. "bar"
	Stacked bool     // histogram: stacked
	NoFill  bool     // histogram: do not fill bars
	Nbins   int      // histogram: number of bins
	Normed  bool     // histogram: normed
}

func (o ParamsPlot) String() (l string) {
	if paramsControl.UsePython {
		return pythonParamsPlot(&o)
	}
	return l
}

/*
// String returns a string representation of parameters
func (o Params) String(forHistogram, for3dPoints bool) (l string) {


	pythonAddToCmd(&l, o.AxCoords, "transform=plt.gca().transAxes")

	// plot and basic options
	if for3dPoints {
		addToCmd(&l, o.Ms > 0, io.Sf("s=%d", o.Ms))
		addToCmd(&l, o.Mec != "", io.Sf("edgecolor='%s'", o.Mec))
		if o.Void {
			addToCmd(&l, o.Void, "c='none'")
		} else {
			addToCmd(&l, o.C != "", io.Sf("c='%s'", o.C))
		}
		addToCmd(&l, o.Void && o.Mec == "", io.Sf("edgecolor='%s'", o.C))
	} else {

	// text and extra parameters
	addToCmd(&l, o.Ha != "", io.Sf("ha='%s'", o.Ha))
	addToCmd(&l, o.Va != "", io.Sf("va='%s'", o.Va))
	addToCmd(&l, o.Rot > 0, io.Sf("rotation=%g", o.Rot))
	addToCmd(&l, o.Fsz > 0, io.Sf("fontsize=%g", o.Fsz))

	// histograms
	if forHistogram {
		addToCmd(&l, len(o.Colors) > 0, io.Sf("color=%s", strings2list(o.Colors)))
		addToCmd(&l, len(o.Type) > 0, io.Sf("histtype='%s'", o.Type))
		addToCmd(&l, o.Stacked, "stacked=1")
		addToCmd(&l, o.NoFill, "fill=0")
		addToCmd(&l, o.Nbins > 0, io.Sf("bins=%d", o.Nbins))
		addToCmd(&l, o.Normed, "normed=1")
	}
	return
}



// floats2list converts slice of floats to string representing a Python list
func floats2list(vals []float64) (l string) {
	l = "["
	for i, v := range vals {
		if i > 0 {
			l += ","
		}
		l += io.Sf("%g", v)
	}
	l += "]"
	return
}

// strings2list converts slice of strings to string representing a Python list
func strings2list(vals []string) (l string) {
	l = "["
	for i, v := range vals {
		if i > 0 {
			l += ","
		}
		l += io.Sf("'%s'", v)
	}
	l += "]"
	return
}


// argsLeg returns legend parameters
func argsLeg(args *Params) (loc string, ncol int, hlen, fsz float64, frame int, out int, outX string) {
	loc = "'best'"
	ncol = 1
	hlen = 3.0
	fsz = 8.0
	frame = 0
	out = 0
	outX = "[0.0, 1.02, 1.0, 0.102]"
	if args == nil {
		return
	}
	if args.LegLoc != "" {
		loc = io.Sf("'%s'", args.LegLoc)
	}
	if args.LegNcol > 0 {
		ncol = args.LegNcol
	}
	if args.LegHlen > 0 {
		hlen = args.LegHlen
	}
	if args.FszLeg > 0 {
		fsz = args.FszLeg
	}
	if args.LegFrame {
		frame = 1
	}
	if args.LegOut {
		out = 1
	}
	if len(args.LegOutX) == 4 {
		outX = io.Sf("[%g, %g, %g, %g]", args.LegOutX[0], args.LegOutX[1], args.LegOutX[2], args.LegOutX[3])
	}
	return
}

// argsFsz allocates args if nil, and sets default fontsizes
func argsFsz(args *A) (txt, lbl, leg, xtck, ytck float64, fontset string) {
	txt, lbl, leg, xtck, ytck, fontset = 11, 10, 9, 8, 8, "stixsans"
	if args == nil {
		return
	}
	if args.Fsz > 0 {
		txt = args.Fsz
	}
	if args.FszLbl > 0 {
		lbl = args.FszLbl
	}
	if args.FszLeg > 0 {
		leg = args.FszLeg
	}
	if args.FszXtck > 0 {
		xtck = args.FszXtck
	}
	if args.FszYtck > 0 {
		ytck = args.FszYtck
	}
	if args.FontSet != "" {
		fontset = args.FontSet
	}
	return
}

// argsFigsize returns figure size data. Defaults are selected if args == nil
func argsFigData(args *A) (figType string, dpi, width, height int) {
	figType, dpi = "png", 150
	prop := 0.75
	widthPt := 400.0
	if args != nil {
		if args.Dpi > 0 {
			dpi = args.Dpi
		}
		if args.Eps {
			figType = "eps"
		}
		if args.Prop > 0 {
			prop = args.Prop
		}
		if args.WidthPt > 0 {
			widthPt = args.WidthPt
		}
	}
	w := widthPt / 72.27 // width in inches
	h := w * prop
	width, height = int(w), int(h)
	return
}

// argsWireSurf collects parameters for Wireframe or Surface
func argsWireSurf(args *A, surf bool) (l string) {
	if args != nil {
		if surf {
			if args.C == "" {
				l += io.Sf(",cmap=getCmap(%d)", args.CmapIdx)
			}
		}
		if args.Rstride < 1 {
			args.Rstride = 1
		}
		if args.Rstride > 0 {
			l += io.Sf(",rstride=%d", args.Rstride)
		}
		if args.Cstride < 1 {
			args.Cstride = 1
		}
		if args.Cstride > 0 {
			l += io.Sf(",cstride=%d", args.Cstride)
		}
	}
	return
}

// argsContour allocates args if nil, sets default parameters, and return formatted parameters
func argsContour(in *A, Z [][]float64) (out *A, colors, levels string) {
	out = in
	if out == nil {
		out = new(A)
	}
	if out.NumFmt == "" {
		out.NumFmt = "%g"
	}
	if out.SelectLw < 0.01 {
		out.SelectLw = 3.0
	}
	if out.Lw < 0.01 {
		out.Lw = 1.0
	}
	if out.Fsz < 0.01 {
		out.Fsz = 8.0
	}
	if len(out.Colors) > 0 {
		colors = io.Sf(",colors=%s", strings2list(out.Colors))
	} else {
		colors = io.Sf(",cmap=getCmap(%d)", out.CmapIdx)
	}
	if len(out.Levels) > 0 {
		levels = io.Sf(",levels=%s", floats2list(out.Levels))
	} else {
		if out.Nlevels > 1 {
			levels = ",levels=" + getContourLevels(out.Nlevels, Z)
		}
	}
	return
}

// getContourLevels computes the list of levels based on min and max values in Z
//  Note: the search for min and max is not very efficient for very large matrix
func getContourLevels(nlevels int, Z [][]float64) (l string) {
	if nlevels < 2 {
		return
	}
	if len(Z) < 1 {
		return
	}
	if len(Z[0]) < 1 {
		return
	}
	minZ, maxZ := Z[0][0], Z[0][0]
	for i := 0; i < len(Z); i++ {
		for j := 0; j < len(Z[i]); j++ {
			if Z[i][j] < minZ {
				minZ = Z[i][j]
			}
			if Z[i][j] > maxZ {
				maxZ = Z[i][j]
			}
		}
	}
	delZ := (maxZ - minZ) / float64(nlevels-1)
	l = "["
	for i := 0; i < nlevels; i++ {
		if i > 0 {
			l += ","
		}
		l += io.Sf("%g", minZ+float64(i)*delZ)
	}
	l += "]"
	return
}

*/
