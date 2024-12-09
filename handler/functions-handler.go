package handler

import "net/http"

func InitilizeHandler() {
	http.Handle("/", http.FileServer(http.Dir("./static")))
}
