// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"./db"
	"./routes"
	"log"
	"net/http"
)

func main() {
	port := ":8080"
	routes.Routes()
	log.Printf("listening on" + port)
	http.ListenAndServe(port, nil)
}
