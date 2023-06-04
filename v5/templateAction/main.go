package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"text/template"
	"time"
)

func proccess(w http.ResponseWriter, r *http.Request) {
	fmt.Println("proccess accessed")
	t, _ := template.ParseFiles("v5/templateAction/tmp1.html")
	rand.Seed(time.Now().Unix())
	t.Execute(w, rand.Intn(10) > 5)
}

func proccess2(w http.ResponseWriter, r *http.Request) {
	fmt.Println("proccess accessed")
	t, _ := template.ParseFiles("v5/templateAction/tmp2.html")
	t.Execute(w, []string{"そこに", "愛は", "あるんか", ""})
}
func proccess3(w http.ResponseWriter, r *http.Request) {
	fmt.Println("proccess accessed")
	t, _ := template.ParseFiles("v5/templateAction/tmp3.html")
	t.Execute(w, []string{"そこに", "愛は", "あるんか", ""})
}

func proccess4(w http.ResponseWriter, r *http.Request) {
	fmt.Println("proccess accessed")
	t, _ := template.ParseFiles("v5/templateAction/tmp4.html")
	t.Execute(w, "愛")
}

func proccess5(w http.ResponseWriter, r *http.Request) {
	fmt.Println("proccess accessed")
	t, _ := template.ParseFiles("v5/templateAction/inc1.html", "v5/templateAction/inc2.html")
	t.Execute(w, "愛")
}

func proccess6(w http.ResponseWriter, r *http.Request) {
	fmt.Println("proccess accessed")
	t, _ := template.ParseFiles("v5/templateAction/vari.html")

	variable := map[string]int{
		"key1": 1,
		"key2": 2,
	}
	t.Execute(w, variable)
}

// func lambda2(str string) string {
// 	layout := "2001-05-01"
// 	return fmt.Sprintln(layout, str)
// }

func lambda(t time.Time) string {
	layout := "2006-01-02" // 定型っぽい、他の値にすると謎の状態になる: template: 29290-31-55
	return t.Format(layout)
}

func proccess7(w http.ResponseWriter, r *http.Request) {

	lamdaMap := template.FuncMap{"fdate": lambda}

	fmt.Println("proccess accessed7")
	// ファイル名が一致しない場合、エラーとなる
	t := template.New("tmp5.html").Funcs(lamdaMap)
	t, _ = t.ParseFiles("tmp5.html")
	t.Execute(w, time.Now())
}

func proccess8(w http.ResponseWriter, r *http.Request) {

	fmt.Println("proccess accessed8")
	t, _ := template.ParseFiles("layout.html", "child.html")
	t.ExecuteTemplate(w, "layout", "やぁ")
}

func proccess9(w http.ResponseWriter, r *http.Request) {
	fmt.Println("proccess accessed9")
	var t *template.Template
	t, _ = template.ParseFiles("block.html", "child.html") // blockコンテンツでの表示するものが無い場合に
	t.ExecuteTemplate(w, "layout", "")
}

func main() {
	server := http.Server{
		Addr: ":7001",
	}
	// fmt.Println("run serve\nhttp://localhost:7001")

	http.HandleFunc("/proccess", proccess)
	http.HandleFunc("/proccess2", proccess2)
	http.HandleFunc("/proccess3", proccess3)
	http.HandleFunc("/proccess4", proccess4)
	http.HandleFunc("/proccess5", proccess5)
	http.HandleFunc("/proccess6", proccess6)
	http.HandleFunc("/proccess7", proccess7)
	http.HandleFunc("/proccess8", proccess8)
	http.HandleFunc("/proccess9", proccess9)
	server.ListenAndServe()
}
