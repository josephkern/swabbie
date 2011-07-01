package main

import (
	
	"http"
	"io/ioutil"
	"os"
	"fmt"
//	"flag"
)

type Page struct {
	Body  []byte
}

// Load a file from the OS, return a Page struct or error
func loadPage(w http.ResponseWriter, r *http.Request)(*Page, os.Error) {
	title := r.URL.Path[1:]
	if title == "" {
		title = "index.html"
	}
	body, err := ioutil.ReadFile(title)
	if err != nil {
		return nil, err
	}
	return &Page{Body: body}, nil
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	p, err := loadPage(w, r)
	if err != nil {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "%s", p.Body)
}

func main() {
	http.HandleFunc("/", rootHandler)
	http.ListenAndServe(":8080", nil)
}