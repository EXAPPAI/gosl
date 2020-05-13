// Copyright 2016 The Gosl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package plt

import (
	"bytes"

	"github.com/cpmech/gosl/io"
)

// pythonGenMat generates matrix
func pythonGenMat(buf *bytes.Buffer, name string, a [][]float64) {
	io.Ff(buf, "%s=np.array([", name)
	for i := range a {
		io.Ff(buf, "[")
		for j := range a[i] {
			io.Ff(buf, "%g,", a[i][j])
		}
		io.Ff(buf, "],")
	}
	io.Ff(buf, "],dtype=float)\n")
}

// pythonGenList generates list
func pythonGenList(buf *bytes.Buffer, name string, a [][]float64) {
	io.Ff(buf, "%s=[", name)
	for i := range a {
		io.Ff(buf, "[")
		for j := range a[i] {
			io.Ff(buf, "%g,", a[i][j])
		}
		io.Ff(buf, "],")
	}
	io.Ff(buf, "]\n")
}

// pythonGenArray generates the NumPy text corresponding to an array of float point numbers
func pythonGenArray(buf *bytes.Buffer, name string, u []float64) {
	io.Ff(buf, "%s=np.array([", name)
	for i := range u {
		io.Ff(buf, "%g,", u[i])
	}
	io.Ff(buf, "],dtype=float)\n")
}

// pythonGen2Arrays generates the NumPy text corresponding to 2 arrays of float point numbers
func pythonGen2Arrays(buf *bytes.Buffer, nameA, nameB string, a, b []float64) {
	pythonGenArray(buf, nameA, a)
	pythonGenArray(buf, nameB, b)
}

// genStrArray generates the NumPy text corresponding to an array of strings
func pythonGenStrArray(buf *bytes.Buffer, name string, u []string) {
	io.Ff(buf, "%s=[", name)
	for i := range u {
		io.Ff(buf, "r'%s',", u[i])
	}
	io.Ff(buf, "]\n")
}
