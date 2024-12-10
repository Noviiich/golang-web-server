package mhttp

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/Noviiich/golang-web-server/utils"
)

type ServerConfig struct {
	StaticDirectory string
	url             string
	port            string
}

type IServer interface {
	InitilizeHandler() error
	InitilizeHandleFunctions() error
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

func (fs ServerConfig) InitilizeHandler() error {
	fmt.Println("Initializing handler with FileServer")

	if !utils.FolderExists(fs.StaticDirectory) {
		return errors.New("'static' folder does not exists: " + fs.StaticDirectory)
	}

	http.Handle("/", http.FileServer(http.Dir(fs.StaticDirectory)))

	fmt.Println("FileServer has been initialized")
	return nil
}

func (fs ServerConfig) InitilizeHandleFunctions() error {
	fmt.Println("Initializing the handler functions...")

	if !utils.FolderExists(fs.StaticDirectory) {
		return errors.New("'static' folder does not exists: " + fs.StaticDirectory)
	}

	http.HandleFunc("/", serverFile)
	http.HandleFunc("/home", homeString)
	http.HandleFunc("/counter", incrementCounter)

	fmt.Println("Handler functions have been initialized")

	return nil
}

func (fs ServerConfig) ListenAndServe() {
	fmt.Printf("Starting server on %s:%s...\n", fs.url, fs.port)
	http.ListenAndServe(fs.url+":"+fs.port, nil)
}
