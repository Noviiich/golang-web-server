package handler

import (
	"fmt"
	"net/http"
)

func InitilizeHandler() {
	fmt.Println("Initializing handler with FileServer")
	http.Handle("/", http.FileServer(http.Dir("../../static")))
	fmt.Println("Handler has been initialized")
}
