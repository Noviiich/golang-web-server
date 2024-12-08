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

func main() {
	http.HandleFunc("/", echoString)

	http.HandleFunc("/", homeString)

	http.ListenAndServe(":8080", nil)
}
