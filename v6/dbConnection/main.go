package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Post struct {
	Id      int
	Content string
	Author  string
}

var Db *sql.DB

func init() {
	var err error
	Db, err = sql.Open("postgres", "host=192.168.33.12 port=5432 user=postgres dbname=postgres password=postgres sslmode=disable")
	fmt.Println("db: ", Db, err)

	if err := Db.Ping(); err != nil {
		fmt.Println("接続エラー:", Db.Ping())
		panic(err)
	}

	if err != nil {
		fmt.Println("db error.")
		panic(err)
	}
}

func Posts(limit int) (posts []Post, err error) {
	rows, err := Db.Query("select id, content, author from posts limit $1", limit)
	if err != nil {
		return
	}

	for rows.Next() {
		post := Post{}
		err = rows.Scan(&post.Id, &post.Content, &post.Author)
		if err != nil {
			return
		}

		posts = append(posts, post)
	}

	rows.Close()
	return
}

func GetPost(id int) (post Post, err error) {
	post = Post{}
	err = Db.QueryRow("select id, content, author from posts where id = $1", id).Scan(&post.Id, &post.Content, &post.Author)
	return
}

func (post *Post) Create() (err error) {
	statement := "insert into posts (content, author) values ($1, $2) returning id"
	fmt.Println("query: ", statement)
	stmt, err := Db.Prepare(statement)
	fmt.Println("create prepare error: ", err)
	if err != nil {
		return
	}
	defer stmt.Close()
	err = stmt.QueryRow(post.Content, post.Author).Scan(&post.Id)
	fmt.Println("result: ", err)
	return
}

func (post *Post) Update() (err error) {
	_, err = Db.Exec("update posts set content = $2, author = $3 where id = $1", post.Id, post.Content, post.Author)
	return
}

func (post *Post) Delete() (err error) {
	_, err = Db.Exec("delete from posts where id = $1", post.Id)
	return
}

func main() {
	post := Post{Content: "愛はあるんか", Author: "俺"}
	fmt.Println("post1: ", post)
	post.Create()
	fmt.Println("post2: ", post)

	readPost, _ := GetPost(post.Id)
	fmt.Println("readpost: ", readPost)

	readPost.Content = "ぼんじゅーる"
	readPost.Author = "俺じゃない"
	readPost.Update()

	posts, _ := Posts(10)
	fmt.Println("posts: ", posts)
	// readPost.Delete()

	for _, pos := range posts {
		fmt.Println(pos)
		pos.Delete()
	}

}
