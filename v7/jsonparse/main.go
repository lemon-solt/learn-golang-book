package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

type Post struct {
	Id       int       `json:"id"`
	Content  string    `json:"content"`
	Author   Author    `json:"author"`
	Comments []Comment `json:"comments"`
}

type Comment struct {
	Id      int    `json:"id"`
	Content string `json:"content"`
	Author  string `json:"author"`
}

type Author struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func main() {
	jsonFIle, err := os.Open("post.json")
	if err != nil {
		log.Fatalln(err)
	}

	defer jsonFIle.Close()

	// marashalパターン
	jsonData, err := ioutil.ReadAll(jsonFIle)
	if err != nil {
		log.Fatalln(err)
	}

	var post Post
	json.Unmarshal(jsonData, &post)
	fmt.Println(post)

	// decorderパターン
	decorder := json.NewDecoder(jsonFIle)
	for {
		var post2 Post
		err2 := decorder.Decode(&post2)
		if err2 == io.EOF {
			break
		}
		if err2 != nil {
			log.Fatalln("error decode", err2)
			return
		}
		fmt.Println("post2: ", post2)
	}

	// json作成1
	output, _ := json.MarshalIndent(&post, "", "\t\t")
	err = ioutil.WriteFile("post2.json", output, 0644)

	// json作成2 こっちが一般的な気がする
	jsonNewFile, _ := os.Create("post3.json")
	encoder := json.NewEncoder(jsonNewFile)
	encoder.Encode(&post)

}
