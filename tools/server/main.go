// Copyright 2016 The Gosl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/cpmech/lootbag/neto"
)

var addr = flag.String("addr", ":8081", "http service address")

func main() {
	flag.Parse()
	hub := NewHub()
	http.HandleFunc("/provider", neto.Ehandler(func(w http.ResponseWriter, r *http.Request) {
		hub.AddProvider(w, r)
	}))
	http.HandleFunc("/observer", neto.Ehandler(func(w http.ResponseWriter, r *http.Request) {
		hub.AddObserver(w, r)
	}))
	log.Printf("... running on %v ...", *addr)
	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
		os.Exit(1)
	}
}
