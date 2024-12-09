package main

import (
	"flag"
	"net/http"

	"github.com/Noviiich/golang-web-server/pkg/handler"
)

var (
	useHandlerFunctions bool
)

func main() {
	flag.BoolVar(&useHandlerFunctions, "use-handler-functions", true, "Runs the Web Browser with handler functions")
	flag.Parse()

	if useHandlerFunctions {
		handler.InitilizeHandleFunctions()
	} else {
		handler.InitilizeHandler()
	}

	http.ListenAndServe(":8080", nil)
}
