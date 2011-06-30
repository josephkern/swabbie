package main

import (
	
	"fmt"
	"http"
	"io/ioutil"
	"os"
	"regexp"
)

type Page struct {
	Title string
	Body  []byte
}

var titleValidator = regexp.MustCompile("^[a-zA-Z0-9].+$")

func getPage(w http.ResponseWriter, r *http.Request) (title string, err os.Error) {
	title = r.URL.Path[lenPath:]
	if !titleValidator.MatchString(title) {
		http.NotFound(w, r)
		err = os.NewError("Invalid Page")
	}
	return
}

func loadPage(title string) (*Page, os.Error) {
	filename := title
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

const lenPath = len("/")

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	title, err := getPage(w, r)
	if err != nil {
		return
	}
	p, err := loadPage(title)
	if err != nil {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}
	fmt.Fprintf(w, "%s", p.Body)
}

func main() {
	http.HandleFunc("/", defaultHandler)
	http.ListenAndServe(":8080", nil)
}