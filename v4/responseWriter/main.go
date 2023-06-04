package main

import (
	"fmt"
	"net/http"
)

func writeHtml(w http.ResponseWriter, r *http.Request) {
	str := `<!DOCTYPE html>
		<html lang="en">
		<head>
			<meta charset="UTF-8">
			<meta http-equiv="X-UA-Compatible" content="IE=edge">
			<meta name="viewport" content="width=device-width, initial-scale=1.0">
			<title>Document</title>
		</head>
		<body>
			<div>
				<h1>そこに愛はあるんか</h1>
			</div>
		</body>
		</html>
		`
	w.Write([]byte(str))
}

func writeHeader(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(501)
	fmt.Fprintln(w, "愛はありません。")
}

func redirectHeader(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Location", "https://google.com")
	w.WriteHeader(302)
}

func main() {
	server := http.Server{
		Addr: ":7001",
	}
	fmt.Printf("run serve")

	http.HandleFunc("/writeHtml", writeHtml)
	http.HandleFunc("/writeHeader", writeHeader)
	http.HandleFunc("/redirectHeader", redirectHeader)
	server.ListenAndServe()
}
