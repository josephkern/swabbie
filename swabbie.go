package main

import (
	
	"http"
	"io/ioutil"
	"os"
	"fmt"
//	"flag"
)

type Page struct {
	Title string
	Body  []byte
}

func indexPage(w http.ResponseWriter, r *http.Request) (title string, err os.Error) {
	title = r.URL.Path[1:]
	if title != "" {
		return
	}
	title = "index.html"
	return
}

// Load a file from the OS, return a Page struct or error
func loadPage(title string)(*Page, os.Error) {
	body, err := ioutil.ReadFile(title)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	title, err := indexPage(w, r)
	if err != nil {
		return
	}
	p, err := loadPage(title)
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