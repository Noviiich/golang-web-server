package main

import (
	"flag"
	"log"

	"github.com/Noviiich/golang-web-server/configs"
	"github.com/Noviiich/golang-web-server/mhttp"
)

var (
	useHandlerFunctions bool
	err                 error
)

func main() {
	flag.BoolVar(&useHandlerFunctions, "use-handler-functions", true, "Runs the Web Browser with handler functions")
	flag.Parse()

	server := mhttp.NewServer(configs.URL, configs.PORT, configs.STATIC_DIRECTORY)

	if useHandlerFunctions {
		err = server.InitilizeHandleFunctions()
	} else {
		err = server.InitilizeFileServer()
	}

	if err != nil {
		log.Fatal(err)
	}
	server.ListenAndServe()
}
