package main

import (
	"fmt"
	"net/http"
	"reflect"
	"runtime"
)


func love(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "愛はあるんか")
}

func log(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := runtime.FuncForPC(reflect.ValueOf(h).Pointer()).Name()
		fmt.Println("handler function called - ", name)
		h(w, r)
	}
}


func main(){
	server := http.Server{
		Addr: ":7000",
	}

	http.HandleFunc("/love", log(love))
	fmt.Println("run serve")
	server.ListenAndServe()
}