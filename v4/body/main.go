package main

import (
	"fmt"
	"net/http"
)

func body(w http.ResponseWriter, r *http.Request) {
	// h := r.Header
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)
	fmt.Print(string(body))
	fmt.Fprintln(w, string(body))
}

func main() {
	server := http.Server{
		Addr: ":7000",
	}

	http.HandleFunc("/body", body)
	server.ListenAndServe()
}
