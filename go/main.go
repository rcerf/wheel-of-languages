// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type Page struct {
	Title string
	Body  []byte
}

func (p *Page) save() error {
	filename := p.Title + ".html"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	path := "/public/templates/"
	filename := title + ".html"
	fmt.Println(filename))

	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	var title string
	if len(r.URL.Path[1:]) < 1 {
		title = "index"
	} else {
		title = r.URL.Path[1:]
	}
	p, _ := loadPage(title)
	fmt.Println(string(p.Body))
	fmt.Fprintf(w, string(p.Body))
}

func main() {
	http.HandleFunc("/", viewHandler)
	http.ListenAndServe(":8080", nil)
}
