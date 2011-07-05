package httplog

import (

	"http"
	"time"
	"fmt"
)

//perhaps write to a file as well? Or leave that for the application main ...

func newEntry (r *http.Request) (logline string) {
// accept a http.Request
// parse request struct
// add time stamp
// return string/io.Writer
}