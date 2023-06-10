package main

import (
	"errors"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db *gorm.DB

func main() {
	log.Println("start server")
	dsn := "host=192.168.33.12 port=5432 user=gwp dbname=gwp password=gwp sslmode=disable"
	Db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("データベースに接続できませんでした。エラー: " + err.Error())
	}

	var user User
	Db.Where("name = ?", "Alice").First(&user)
	log.Println("user: ", user)

	var user1 User
	result1 := Db.First(&user1)
	log.Println("result1: ", result1)
	// check error ErrRecordNotFound
	if errors.Is(result1.Error, gorm.ErrRecordNotFound) {
		log.Fatal(result1.Error)
	}

	log.Println("count:", result1.RowsAffected)

	// 何も指定せず、単体取得
	user2 := User{}
	result2 := Db.Take(&user2)
	// SELECT * FROM users LIMIT 1;
	log.Println("take:", user2)
	if errors.Is(result2.Error, gorm.ErrRecordNotFound) {
		log.Fatal(result2.Error)
	}

	// 降順で単体取得
	user3 := User{}
	result3 := Db.Last(&user3)
	// SELECT * FROM users ORDER BY id DESC LIMIT 1;
	log.Println("last:", user3)
	if errors.Is(result3.Error, gorm.ErrRecordNotFound) {
		log.Fatal(result3.Error)
	}

	// プライマリーキーで取得
	Db.First(&user, 4)
	Db.First(&user, "id = ?", 4)

	find(Db)

}

type User struct {
	gorm.Model
	Name  string
	Email string
}

// 全件取得
func find(db *gorm.DB) {
	users := []User{}
	result := db.Find(&users)
	log.Println("user:", users)
	if result.Error != nil {
		log.Fatal(result.Error)
	}
	log.Println("count:", result.RowsAffected)
}
