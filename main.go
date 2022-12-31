package main

import (
	"flag"
	"fmt"
	"github.com/apex/gateway"
	"io"
	"log"
	"net/http"
)

func main() {
	port := flag.Int("port", -1, "specify a port to use http rather than AWS Lambda")
	flag.Parse()
	listener := gateway.ListenAndServe
	portStr := "n/a"
	if *port != -1 {
		portStr = fmt.Sprintf(":%d", *port)
		listener = http.ListenAndServe
		http.HandleFunc("/", func(w http.ResponseWriter, request *http.Request) {
			io.WriteString(w, "Hello, world!\n")
		})
	}
	log.Fatal(listener(portStr, nil))
}
