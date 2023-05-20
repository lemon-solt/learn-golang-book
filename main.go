package main

import (
	"fmt"
	"net/http"
)

func handler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprint(writer, "vagrant golang %s", request.URL.Path[1:])
}

func main() {
	fmt.Println("start server")
	http.HandleFunc("/", handler)
	http.ListenAndServe(":7000", nil)
}
