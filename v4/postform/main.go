package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

/*
json: PostFormの値ではない為、解析されない
application/jsonで送信される情報は、form-dataやx-www-form-urlencodedではなく、
goの解析メソッドは、json以外の情報を受け取る仕様になっている。その為、別の方法でjsonを受け取る必要がある
*/
func postform(w http.ResponseWriter, r *http.Request) {
	// r.ParseForm()
	fmt.Println(r.Form)
	fmt.Fprintln(w, r.Form)     // postの値とurlに同じキーを含む場合はどちらも取得できる
	fmt.Fprintln(w, r.PostForm) // postのみを優先する場合はPostFormを私用する
	fmt.Fprintln(w, r.PostFormValue("hello"))
}

func fileform(w http.ResponseWriter, r *http.Request) {
	// r.ParseMultipartForm(1024)
	// fileheader := r.MultipartForm.File["file"][0]
	// file, err := fileheader.Open()

	file, _, err := r.FormFile("file")
	if err == nil {
		data, err := ioutil.ReadAll(file)
		if err == nil {
			fmt.Fprintln(w, string(data))
		}
	}
}

func main() {
	server := http.Server{
		Addr: ":7001",
	}
	fmt.Printf("run serve")

	http.HandleFunc("/postform", postform)
	http.HandleFunc("/fileform", fileform)
	server.ListenAndServe()
}
