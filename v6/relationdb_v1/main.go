package main

import (
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/lib/pq"
)

type Post struct {
	Id       int
	Content  string
	Author   string
	Comments []Comment
}

type Comment struct {
	Id      int
	Content string
	Author  string
	Post    *Post
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

func (comment *Comment) Create() (err error) {
	if comment.Post == nil {
		err = errors.New("投稿が見つからない")
		return
	}
	err = Db.QueryRow("insert into comments (content, author, post_id) values (?,?,?) returning id", comment.Content, comment.Author, comment.Post.Id).Scan(&comment.Id)
	return
}

func GetPost(id int) (post Post, err error) {
	post = Post{}
	post.Comments = []Comment{}
	err = Db.QueryRow("select id, content, author from posts where id = $1", id).Scan(&post.Id, &post.Content, &post.Author)

	rows, err := Db.Query("select id, content, author from comments where post_id = $1", id)
	if err != nil {
		return
	}

	for rows.Next() {
		comment := Comment{Post: &post}
		err = rows.Scan(&comment.Id, &comment.Content, &comment.Author)
		if err != nil {
			return
		}
		post.Comments = append(post.Comments, comment)
	}
	rows.Close()
	return

}

func (post *Post) Create() (err error) {
	err = Db.QueryRow("insert into posts (content, author) values ($1, $2) returning id", post.Content, post.Author).Scan(&post.Id)
	return
}

func main() {
	post := Post{Content: "愛はあるんか", Author: "俺やろ"}
	post.Create()

	comment := Comment{Content: "ひましたり", Author: "暇人", Post: &post}
	comment.Create()
	readPost, _ := GetPost(post.Id)

	fmt.Println(readPost)
}
