// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package routes

import (
	"fmt"
	"net/http"
)

func Routes() {
	http.Handle("/", http.FileServer(http.Dir("./public")))
	http.HandleFunc("/links", linksHandler)
}

func linksHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("links handler is handling")
}
