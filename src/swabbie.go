package main

import (
	
	"http"
	"io/ioutil"
	"os"
	"fmt"
	"flag"
	"strconv"
)

var port *int = flag.Int("port", 8080, "Port for web server to listen on. Default 8080.")

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
//	req := http.Request.Write(r)
	// host - userid [day/month/year :hour:minute:second zone] request response objectsize referer user-agent
	// need to add in the response fields from the ResponseWriter interface
	fmt.Printf("%s - userid [time] %s %s %s %s\n", r.Host, r.Method, r.RawURL, r.Referer, r.UserAgent)
	fmt.Fprintf(w, "%s", p.Body)
}

func main() {
	flag.Parse()

	http.HandleFunc("/", rootHandler)
	
	Port := ":" + strconv.Itoa(*port)
	http.ListenAndServe(Port, nil)
}