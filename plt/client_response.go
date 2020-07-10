// Copyright 2016 The Gosl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package plt

// response holds the data received from the Observer
type response struct {
	PlotName string `json:"plotName"`
	Action   string `json:"action"`
	Payload  string `json:"payload"`
}
