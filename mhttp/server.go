package mhttp

import (
	"fmt"
	"net/http"
)

type ServerConfig struct {
	StaticDirectory string
	url             string
	port            string
}

func (fs ServerConfig) InitilizeHandler() {
	fmt.Println("Initializing handler with FileServer")

	http.Handle("/", http.FileServer(http.Dir(fs.StaticDirectory)))

	fmt.Println("FileServer has been initialized")
}

func InitilizeHandleFunctions() {
	fmt.Println("Initializing the handler functions...")

	http.HandleFunc("/", serverFile)
	http.HandleFunc("/home", homeString)
	http.HandleFunc("/counter", incrementCounter)

	fmt.Println("Handler functions have been initialized")
}
