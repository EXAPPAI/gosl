// Copyright 2016 The Gosl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package plt

import "github.com/cpmech/gosl/io"

/*
// updateBufferFirstParamsAndClose updates buffer with the first parameters and close with ")\n"
func updateBufferFirstParamsAndClose(buf *bytes.Buffer, params *Params, forHistogram, for3dPoints bool) {
	if buf == nil {
		return
	}
	if params == nil {
		io.Ff(buf, ")\n")
		return
	}
	txt := params.String(forHistogram, for3dPoints)
	if txt == "" {
		io.Ff(buf, ")\n")
		return
	}
	io.Ff(buf, txt+")\n")
}
*/

// pythonPutParams updates buffer with parameters and close with ")\n".
// NOTE: interface{} is a slice of Params
func pythonPutParams(params interface{}) {
	txt := ""
	io.Pf(">>>> %T\n", params)
	switch pp := params.(type) {
	case []ParamsPlot:
		if len(pp) > 0 {
			txt = "," + pp[0].String()
		}
	}
	io.Ff(&pythonBuffer, txt+")\n")
}

// pythonAddToCmd adds new option to list of commands separated with commas
func pythonAddToCmd(line *string, condition bool, delta string) {
	if condition {
		if len(*line) > 0 {
			*line += ","
		}
		*line += delta
	}
}

// pythonGetHideList returns a string representing the "spines-to-remove" list in Python
func pythonGetHideList(params ...ParamsFrame) (l string) {
	if len(params) < 1 {
		return
	}
	p := params[0]
	if p.HideLeft || p.HideRight || p.HideBottom || p.HideTop {
		c := ""
		pythonAddToCmd(&c, p.HideLeft, "'left'")
		pythonAddToCmd(&c, p.HideRight, "'right'")
		pythonAddToCmd(&c, p.HideBottom, "'bottom'")
		pythonAddToCmd(&c, p.HideTop, "'top'")
		l = "[" + c + "]"
	}
	return
}

func pythonParamsPlot(o *ParamsPlot) (l string) {
	if o.MarkerVoid && o.Color == "" {
		o.Color = "red"
	}
	pythonAddToCmd(&l, o.Color != "", io.Sf("color='%s'", o.Color))
	pythonAddToCmd(&l, o.MarkerSize > 0, io.Sf("ms=%d", o.MarkerSize))
	pythonAddToCmd(&l, o.MarkerEdgeColor != "", io.Sf("markeredgecolor='%s'", o.MarkerEdgeColor))
	pythonAddToCmd(&l, o.MarkerVoid, "markerfacecolor='none'")
	pythonAddToCmd(&l, o.MarkerVoid && o.MarkerEdgeColor == "", io.Sf("markeredgecolor='%s'", o.Color))
	pythonAddToCmd(&l, o.MarkerEdgeWidth > 0, io.Sf("mew=%g", o.MarkerEdgeWidth))
	pythonAddToCmd(&l, o.Alpha > 0, io.Sf("alpha=%g", o.Alpha))
	pythonAddToCmd(&l, o.Marker != "", io.Sf("marker='%s'", o.Marker))
	pythonAddToCmd(&l, o.LineStyle != "", io.Sf("linestyle='%s'", o.LineStyle))
	pythonAddToCmd(&l, o.LineWidth > 0, io.Sf("lw=%g", o.LineWidth))
	pythonAddToCmd(&l, o.Label != "", io.Sf("label='%s'", o.Label))
	pythonAddToCmd(&l, o.MarkEvery > 0, io.Sf("markevery=%d", o.MarkEvery))
	pythonAddToCmd(&l, o.Zindex > 0, io.Sf("zorder=%d", o.Zindex))
	pythonAddToCmd(&l, o.NoClip, "clip_on=0")
	pythonAddToCmd(&l, o.ShapeFaceColor != "", io.Sf("facecolor='%s'", o.ShapeFaceColor))
	pythonAddToCmd(&l, o.ShapeEdgeColor != "", io.Sf("edgecolor='%s'", o.ShapeEdgeColor))
	return
}
