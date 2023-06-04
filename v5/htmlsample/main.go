package main

import (
	"fmt"
	"net/http"
	"text/template"
)

func proccess(w http.ResponseWriter, r *http.Request) {
	fmt.Println("proccess accessed")
	t, _ := template.ParseFiles("v5/htmlsample/temp1.html")
	t.Execute(w, "hello world!")
}

func proccess2(w http.ResponseWriter, r *http.Request) {
	fmt.Println("proccess2 accessed")
	t := template.Must(template.ParseFiles("v5/htmlsample/temp2.html"))
	t.Execute(w, "愛はあるんか")
}

func proccess3(w http.ResponseWriter, r *http.Request) {
	fmt.Println("proccess3 accessed")
	t, _ := template.ParseFiles("v5/htmlsample/temp1.html", "v5/htmlsample/temp2.html", "v5/htmlsample/temp3.html")
	t.ExecuteTemplate(w, "temp3.html", "愛はあるんか") // templateのpathではなく、ファイル名を指定する
}

func proccess4(w http.ResponseWriter, r *http.Request) {
	str := `<h1>愛はシンプル {{ . }}</h1>`
	t := template.New("str.html")
	t, _ = t.Parse(str)
	t.Execute(w, "愛はあるんです")
}
func main() {
	server := http.Server{
		Addr: ":7001",
	}
	fmt.Println("run serve\nhttp://localhost:7001")

	http.HandleFunc("/proccess", proccess)
	http.HandleFunc("/proccess2", proccess2)
	http.HandleFunc("/proccess3", proccess3)
	http.HandleFunc("/proccess4", proccess4)
	server.ListenAndServe()
}
