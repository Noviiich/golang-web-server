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

func echoString(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello")
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
	http.HandleFunc("/", echoString)

	http.HandleFunc("/home", homeString)

	http.HandleFunc("/counter", incrementCounter)

	http.ListenAndServe(":8080", nil)
}
