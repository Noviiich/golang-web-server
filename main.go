package main

import (
	"fmt"
	"net/http"
	"strconv"
	"sync"
)

var (
	counter int
	mutex   sync.Mutex // Мьютекс для защиты доступа к counter
)

func serverFile(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, r.URL.Path[1:])
}

func homeString(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Дом")
}

func incrementCounter(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	counter++
	fmt.Fprintf(w, strconv.Itoa(counter))
	mutex.Unlock()
}

func main() {
	http.HandleFunc("/", serverFile)

	http.HandleFunc("/home", homeString)

	http.HandleFunc("/counter", incrementCounter)

	http.ListenAndServe(":8080", nil)
}
