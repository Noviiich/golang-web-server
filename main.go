package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/Noviiich/golang-web-server/pkg/handler"
	"github.com/Noviiich/golang-web-server/pkg/utils"
)

var (
	valid_args = []string{"handler_functions", "handler_directory"}
)

func printUsage() {
	fmt.Println("Usages:")
	fmt.Println("go run main.go option")
	fmt.Println("  options:")
	fmt.Println("    handler_functions")
	fmt.Println("    handler_directory")
	fmt.Println("  example:")
	fmt.Println("go run main.go handler_directory")
}

func main() {
	args := os.Args[1:]
	if len(args) <= 0 || !utils.StringArrayContains(valid_args, args[0]) {
		printUsage()
		return
	} else {

		switch args[0] {

		case "handler_functions":
			handler.InitilizeHandleFunctions()

		case "handler_directory":
			handler.InitilizeHandler()

		}
	}

	http.ListenAndServe(":8080", nil)
}
