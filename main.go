package main

import (
	"fmt"
	"net/http"
)

func echoString(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello")
}

func homeString(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Дом")
}

func incrementCounter(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Counter")
}

func main() {
	http.HandleFunc("/", echoString)

	http.HandleFunc("/home", homeString)

	http.HandleFunc("/counter", incrementCounter)

	http.ListenAndServe(":8080", nil)
}
