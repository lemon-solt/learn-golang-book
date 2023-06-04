package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"io/ioutil"
)

type Post struct {
	Id      int
	Content string
	Author  string
}

func store(data interface{}, filename string) {
	buffer := new(bytes.Buffer)
	encoder := gob.NewEncoder(buffer)
	err := encoder.Encode(data)
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile(filename, buffer.Bytes(), 0600)
	if err != nil {
		panic(err)
	}
}

func load(data interface{}, filename string) {
	raw, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	buffer := bytes.NewBuffer(raw)
	dec := gob.NewDecoder(buffer)
	err = dec.Decode(data)
	if err != nil {
		panic(err)
	}
}

func main() {
	posts := []Post{
		{Id: 1, Content: "愛はあるんか", Author: "ないです"},
		{Id: 2, Content: "愛はあるんか2", Author: "ないです2"},
		{Id: 3, Content: "愛はあるんか3", Author: "ないです3"},
	}

	store(posts, "post1")
	var postRead []Post
	load(&postRead, "post1")
	fmt.Println("post read: ", postRead)
}
