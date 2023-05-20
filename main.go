package main

import (
	"fmt"
	"net/http"
)

func handler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(writer, "vagrant golang %s")
}

func main() {
	fmt.Println("start server")
	http.HandleFunc("/", handler)
	http.ListenAndServe(":7000", nil)

}
