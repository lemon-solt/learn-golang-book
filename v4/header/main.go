package main

import (
	"fmt"
	"net/http"
)

func header(w http.ResponseWriter, r *http.Request) {
	// h := r.Header
	h := r.Header.Get("Accept-Encoding")
	fmt.Fprintln(w, h)
}

func main() {
	server := http.Server{
		Addr: ":7000",
	}

	http.HandleFunc("/header", header)
	server.ListenAndServe()
}
