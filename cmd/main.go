package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/Noviiich/golang-web-server/handler"
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
	fmt.Println("Starting server on localhost:8080...")
	http.ListenAndServe(":8080", nil)
}
