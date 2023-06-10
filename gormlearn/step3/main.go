package main

import (
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
	save(Db)
}

type User struct {
    gorm.Model
    Name  string
    Email string
}


// 更新(upsert)
func save(db *gorm.DB) {
	// 構造体にidが無い場合はinsertされる
	user1 := User{}
	user1.Name = "愛はあるんか２"
	result1 := db.Save(&user1)
	if result1.Error != nil {
		log.Fatal(result1.Error)
	}
	log.Println("count:", result1.RowsAffected)
	log.Println("user1:", user1)

	// 先にユーザーを取得する
	user2 := User{}
	db.First(&user2)

	// 構造体にidがある場合はupdateされる
	user2.Name = "そうかい"
	result2 := db.Save(&user2)
	if result2.Error != nil {
		log.Fatal(result2.Error)
	}
	log.Println("count:", result2.RowsAffected)
	log.Println("user2:", user2)
}