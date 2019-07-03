package main

import (
	"fmt"
	"net/http"
	"os"
)

const defaultPort = "8080"
const defaultHostname = "default-hostname"

func main() {
	os.Getenv("PORT")
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	errorMessage := ""
	hostname, err := os.Hostname()
	if err != nil {
		hostname = defaultHostname
		errorMessage = err.Error()
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, `{"hostname":"%s", "error": "%s"}`, hostname, errorMessage)
	})

	if err := http.ListenAndServe("0.0.0.0:"+port, nil); err != nil {
		panic(err)
	}
}
