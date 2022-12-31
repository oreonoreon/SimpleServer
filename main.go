package main

import (
	"io"
	"net/http"
)

func main() {
	var port string
	port = "80"

	http.HandleFunc("/", func(w http.ResponseWriter, request *http.Request) {
		io.WriteString(w, "Hello, world!\n")
	})
	http.ListenAndServe(":"+port, nil)
}
