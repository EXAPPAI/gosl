// Copyright 2016 The Gosl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package plt

import (
	"bytes"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/cpmech/gosl/chk"
	"github.com/cpmech/gosl/io"
)

// buffer holding Python commands
var pythonBuffer bytes.Buffer

// buffer holding Python extra artists commands
var pythonBufferEa bytes.Buffer

// flag indicating that Python axes3d ('AX3D') has been created
var pythonAxes3dCreated bool

// PythonCommands adds Python commands to be called when plotting
func PythonCommands(text string) {
	io.Ff(&pythonBuffer, text)
}

// PythonFile loads Python file and copy its contents to temporary buffer
func PythonFile(filename string) {
	b := io.ReadFile(filename)
	io.Ff(&pythonBuffer, string(b))
	return
}

// pythonShow runs python and then shows figure
func pythonShow() {
	io.Ff(&pythonBuffer, "plt.show()\n")
	pythonRun("")
}

// pythonSave executes python script and generates figure (creates the output directory too)
//  NOTE: the file name will be fnkey + .png (default) or .eps depending on the Reset function
func pythonSave(fnkey string) {
	empty := paramsControl.OutDir == "" || fnkey == ""
	if empty {
		chk.Panic("directory and filename key must not be empty\n")
	}
	err := os.MkdirAll(paramsControl.OutDir, 0777)
	if err != nil {
		chk.Panic("cannot create directory to save figure file:\n%v\n", err)
	}
	fileExt := ".png"
	if paramsControl.FigEps {
		fileExt = ".eps"
	}
	fn := filepath.Join(paramsControl.OutDir, fnkey+fileExt)
	io.Ff(&pythonBuffer, "plt.savefig(r'%s', bbox_inches='tight', bbox_extra_artists=EXTRA_ARTISTS)\n", fn)
	pythonRun(fn)
}

// pythonShowSave runs python to show figure and then saves figure file
func pythonShowSave(fnkey string) {
	empty := paramsControl.OutDir == "" || fnkey == ""
	if empty {
		chk.Panic("directory and filename key must not be empty\n")
	}
	uid := pythonGenUID()
	io.Ff(&pythonBuffer, "fig%d = plt.gcf()\n", uid)
	io.Ff(&pythonBuffer, "plt.show()\n")
	err := os.MkdirAll(paramsControl.OutDir, 0777)
	if err != nil {
		chk.Panic("cannot create directory to save figure file:\n%v\n", err)
	}
	fileExt := ".png"
	if paramsControl.FigEps {
		fileExt = ".eps"
	}
	fn := filepath.Join(paramsControl.OutDir, fnkey+fileExt)
	io.Ff(&pythonBuffer, "fig%d.savefig(r'%s', bbox_inches='tight', bbox_extra_artists=EXTRA_ARTISTS)\n", uid, fn)
	pythonRun("")
}

// pythonInit initalizes python buffers and sets default values
//
//   The following data is set:
//     fontsizes:
//        args.Fsz     float64 // font size
//        args.FszLbl  float64 // font size of labels
//        args.FszLeg  float64 // font size of legend
//        args.FszXtck float64 // font size of x-ticks
//        args.FszYtck float64 // font size of y-ticks
//     figure data:
//        args.Dpi     int     // dpi to be used when saving figure. default = 96
//        args.Png     bool    // save png file
//        args.Eps     bool    // save eps file
//        args.Prop    float64 // proportion: height = width * prop
//        args.WidthPt float64 // width in points. Get this from LaTeX using \showthe\columnwidth
func pythonInit() {

	// clear buffer and start python code
	pythonBuffer.Reset()
	pythonBufferEa.Reset()
	io.Ff(&pythonBuffer, pythonHeader)
	pythonAxes3dCreated = false

	/*
		txt, lbl, leg, xtck, ytck, fontset := argsFsz(args)
		figType, dpi, width, height := argsFigData(args)
		io.Ff(&pythonBuffer, "plt.rcdefaults()\n")
		io.Ff(&pythonBuffer, "plt.rcParams.update({\n")
		io.Ff(&pythonBuffer, "    'font.size'       : %g,\n", txt)
		io.Ff(&pythonBuffer, "    'axes.labelsize'  : %g,\n", lbl)
		io.Ff(&pythonBuffer, "    'legend.fontsize' : %g,\n", leg)
		io.Ff(&pythonBuffer, "    'xtick.labelsize' : %g,\n", xtck)
		io.Ff(&pythonBuffer, "    'ytick.labelsize' : %g,\n", ytck)
		io.Ff(&pythonBuffer, "    'mathtext.fontset': '%s',\n", fontset)
		io.Ff(&pythonBuffer, "    'figure.figsize'  : [%d,%d],\n", width, height)
		switch figType {
		case "eps":
			io.Ff(&pythonBuffer, "    'backend'            : 'ps',\n")
			io.Ff(&pythonBuffer, "    'text.usetex'        : True,\n")  // very IMPORTANT to avoid Type 3 fonts
			io.Ff(&pythonBuffer, "    'ps.useafm'          : True,\n")  // very IMPORTANT to avoid Type 3 fonts
			io.Ff(&pythonBuffer, "    'pdf.use14corefonts' : True})\n") // very IMPORTANT to avoid Type 3 fonts
			fileExt = ".eps"
		default:
			io.Ff(&pythonBuffer, "    'savefig.dpi'     : %d})\n", dpi)
			fileExt = ".png"
		}
	*/
}

// pythonRun calls Python to generate plot
func pythonRun(fn string) {
	// write file
	goslPy := filepath.Join(paramsControl.TmpDir, "gosl.py")
	io.WriteFile(goslPy, &pythonBufferEa, &pythonBuffer)
	python := os.Getenv("PYTHON")
	if python == "" {
		python = "python"
	}

	// set command
	cmd := exec.Command(python, goslPy)
	var out, serr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &serr

	// call Python
	err := cmd.Run()
	if err != nil {
		chk.Panic("call to Python failed:\n%v\n", serr.String())
	}

	// show filename
	if fn != "" {
		io.Pf("file <%s> written\n", fn)
	}

	// show output
	io.Pf("%s", out.String())
}

// pythonGenUID returns an unique id for python variables
func pythonGenUID() int { return pythonBuffer.Len() }

// pythonBool converts Go bool to Python bool
func pythonBool(flag bool) string {
	if flag {
		return "True"
	}
	return "False"
}

const pythonHeader = `### file generated by Gosl #################################################
import numpy as np
import matplotlib.pyplot as plt
import matplotlib.ticker as tck
import matplotlib.patches as pat
import matplotlib.path as pth
import matplotlib.patheffects as pff
import matplotlib.lines as lns
import mpl_toolkits.mplot3d as m3d
NaN = np.NaN
EXTRA_ARTISTS = []
def addToEA(obj):
    if obj!=None: EXTRA_ARTISTS.append(obj)
COLORMAPS = [plt.cm.bwr, plt.cm.RdBu, plt.cm.hsv, plt.cm.jet, plt.cm.terrain, plt.cm.pink, plt.cm.Greys]
def getCmap(idx): return COLORMAPS[idx %% len(COLORMAPS)]
`
