package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"sync"
)

var counter int
var mutex sync.Mutex // Мьютекс для защиты доступа к counter

func InitilizeHandleFunctions() {
	http.HandleFunc("/", serverFile)

	http.HandleFunc("/home", homeString)

	http.HandleFunc("/counter", incrementCounter)
}

func serverFile(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "../../static")
}

func homeString(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Дом")
}

func incrementCounter(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Initializing the handler functions...")
	mutex.Lock()
	counter++
	fmt.Fprint(w, strconv.Itoa(counter))
	mutex.Unlock()
	fmt.Println("Handler functions have been initialized")
}
