package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"sync"
)

var (
	counter    int
	mutex      sync.Mutex // Мьютекс для защиты доступа к counter
	valid_args = []string{"handler_functions", "handler_directory"}
)

func serverFile(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, r.URL.Path[1:])
}

func homeString(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Дом")
}

func incrementCounter(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	counter++
	fmt.Fprintf(w, strconv.Itoa(counter))
	mutex.Unlock()
}

func startHandleFunctions() {
	http.HandleFunc("/", serverFile)

	http.HandleFunc("/home", homeString)

	http.HandleFunc("/counter", incrementCounter)
}

func initilizeHandler() {
	http.Handle("/", http.FileServer(http.Dir("./static")))
}

func contains(options []string, str string) bool {
	for _, option := range options {
		if option == str {
			return true
		}

	}
	return false
}

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
	if len(args) <= 0 || !contains(valid_args, args[0]) {
		printUsage()
		return
	} else {

		switch args[0] {

		case "handler_functions":
			startHandleFunctions()

		case "handler_directory":
			initilizeHandler()

		}
	}

	http.ListenAndServe(":8080", nil)
}
