package main

import (
	
	"http"
	"io/ioutil"
	"os"
	"fmt"
)

type Page struct {
	Body  []byte
}

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