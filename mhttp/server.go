package mhttp

import (
	"fmt"
	"net/http"
)

func InitilizeHandler() {
	fmt.Println("Initializing handler with FileServer")
	http.Handle("/", http.FileServer(http.Dir("./static")))
	fmt.Println("Handler has been initialized")
}

func InitilizeHandleFunctions() {
	fmt.Println("Initializing the handler functions...")

	http.HandleFunc("/", serverFile)
	http.HandleFunc("/home", homeString)
	http.HandleFunc("/counter", incrementCounter)

	fmt.Println("Handler functions have been initialized")
}
