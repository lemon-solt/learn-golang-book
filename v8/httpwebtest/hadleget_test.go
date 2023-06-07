package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

var mux *http.ServeMux
var writer *httptest.ResponseRecorder

func TestMain(m *testing.M) {
	setUp()
	code := m.Run()
	os.Exit(code)
}

func setUp() {
	mux = http.NewServeMux()
	mux.HandleFunc("/post/", handleRequest)
	writer = httptest.NewRecorder()
}

func TestHandleGet(t *testing.T) {
	request, _ := http.NewRequest("GET", "/post/6", nil)
	mux.ServeHTTP(writer, request)

	if writer.Code != 200 {
		t.Error("response code is error", writer.Code)
	}

	var post Post
	json.Unmarshal(writer.Body.Bytes(), &post)
	if post.Id != 6 {
		t.Error("cannnot retrieve json")
	}
}

func TestHandlePut(t *testing.T) {
	json := strings.NewReader(`{"content": "それやで", "author": "俺のやつ"`)
	request, _ := http.NewRequest("PUT", "/post/6", json)
	mux.ServeHTTP(writer, request)

	if writer.Code != 200 {
		t.Error("error code: ", writer.Code)
	}

}
