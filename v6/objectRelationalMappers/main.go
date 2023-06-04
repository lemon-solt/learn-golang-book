package main

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

type Post struct {
	Id        int
	Content   string
	Author    string `sql:"not null"`
	Comments  []Comment
	CreatedAt time.Time
}

type Comment struct {
	Id        int
	Content   string
	Author    string `sql:"not null"`
	PostId    int
	CreatedAt time.Time
}

var Db *gorm.DB

func init() {
	var err error
	Db, err = gorm.Open("postgres", "host=192.168.33.12 port=5432 user=postgres dbname=postgres password=postgres sslmode=disable")
	fmt.Println("db: ", Db, err)

	// if err := Db.Ping(); err != nil {
	// 	fmt.Println("接続エラー:", Db.Ping())
	// 	panic(err)
	// }

	if err != nil {
		fmt.Println("db error.")
		panic(err)
	}

	Db.AutoMigrate(&Post{}, &Comment{})
}

func main() {
	post := Post{Content: "愛はあるんか？", Author: "俺なのか"}
	fmt.Println("post: ", post)
	Db.Create(&post)
	fmt.Println("new record post: ", post)

	comment := Comment{Content: "その愛ええやん", Author: "俺ではないやろ"}
	Db.Model(&post).Association("Comments").Append(comment)

	var readPost Post
	Db.Where("author = ?", "俺なのか").First(&readPost)

	var comments []Comment
	Db.Model(&readPost).Related(&comments)

	fmt.Println(comments)
}
