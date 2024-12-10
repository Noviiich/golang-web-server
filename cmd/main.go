package main

import (
	"flag"

	"github.com/Noviiich/golang-web-server/mhttp"
)

var (
	useHandlerFunctions bool
)

const (
	url             string = "localhost"
	port            string = "8080"
	staticDirectory string = "./static"
)

func main() {
	flag.BoolVar(&useHandlerFunctions, "use-handler-functions", true, "Runs the Web Browser with handler functions")
	flag.Parse()

	server := mhttp.NewServer(url, port, staticDirectory)

	if useHandlerFunctions {
		server.InitilizeHandleFunctions()
	} else {
		server.InitilizeHandler()
	}
	server.ListenAndServe()
}
