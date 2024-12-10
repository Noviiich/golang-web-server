package mhttp

import (
	"fmt"
	"net/http"
	"strconv"
	"sync"
)

var counter int
var mutex sync.Mutex
var staticDirectory = "./static"

func serverFile(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, staticDirectory)
}

func homeString(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Дом")
}

func incrementCounter(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	counter++
	fmt.Fprint(w, strconv.Itoa(counter))
	mutex.Unlock()
}
