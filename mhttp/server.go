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
	InitilizeFileServer() error
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

func (fs ServerConfig) InitilizeFileServer() error {
	fmt.Println("Initializing handler with FileServer")

	if err := utils.ValidateFolder(fs.StaticDirectory); err != nil {
		return errors.New("couldn't validate static folder, error: " + err.Error())
	}

	http.Handle("/", http.FileServer(http.Dir(fs.StaticDirectory)))

	fmt.Println("FileServer has been initialized")
	return nil
}

func (fs ServerConfig) InitilizeHandleFunctions() error {
	fmt.Println("Initializing the handler functions...")

	if err := utils.ValidateFolder(fs.StaticDirectory); err != nil {
		return errors.New("couldn't validate static folder, error: " + err.Error())
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
