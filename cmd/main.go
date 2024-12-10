package main

import (
	"flag"

	"github.com/Noviiich/golang-web-server/configs"
	"github.com/Noviiich/golang-web-server/mhttp"
)

var (
	useHandlerFunctions bool
)

func main() {
	flag.BoolVar(&useHandlerFunctions, "use-handler-functions", true, "Runs the Web Browser with handler functions")
	flag.Parse()

	server := mhttp.NewServer(configs.URL, configs.PORT, configs.STATIC_DIRECTORY)

	if useHandlerFunctions {
		server.InitilizeHandleFunctions()
	} else {
		server.InitilizeHandler()
	}
	server.ListenAndServe()
}
