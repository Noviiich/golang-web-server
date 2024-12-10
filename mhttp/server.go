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

type IServer interface {
	InitilizeHandler()
	InitilizeHandleFunctions()
	ListenAndServe()
}

func NewServer(url, port, staticDirectory string) ServerConfig {
	var serverConfig = ServerConfig{
		StaticDirectory: staticDirectory,
		url:             url,
		port:            port,
	}

	return serverConfig
}

func (fs ServerConfig) InitilizeHandler() {
	fmt.Println("Initializing handler with FileServer")

	http.Handle("/", http.FileServer(http.Dir(fs.StaticDirectory)))

	fmt.Println("FileServer has been initialized")
}

func (fs ServerConfig) InitilizeHandleFunctions() {
	fmt.Println("Initializing the handler functions...")

	http.HandleFunc("/", serverFile)
	http.HandleFunc("/home", homeString)
	http.HandleFunc("/counter", incrementCounter)

	fmt.Println("Handler functions have been initialized")
}

func (fs ServerConfig) ListenAndServe() {
	fmt.Printf("Starting server on %s:%s...\n", fs.url, fs.port)
	http.ListenAndServe(fs.url+":"+fs.port, nil)
}
