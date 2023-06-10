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
	update(Db)

}

type User struct {
    gorm.Model
    Name  string
    Email string
}


// 更新(upsert)
func update(db *gorm.DB) {
	result := db.Model(&User{}).Where("id = 5").Update("name", "ジョージ")
	if result.Error != nil {
		log.Fatal(result.Error)
	}
	log.Println("count:", result.RowsAffected)

	user := User{}
	db.Where("id = 4").Take(&user)
	log.Println("user:", user)
}
