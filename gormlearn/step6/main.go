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
	selectUpdate(Db)
}

type Paformer struct {
    gorm.Model
    Name  string
    Email string
	Age int
	IsActive bool
}

// 一括更新
func selectUpdate(db *gorm.DB) {
	// Selectで指定することで更新されます
	result := db.Model(Paformer{}).Where("name = ?", "海賊王ルフィ").Select("name", "is_active").Updates(Paformer{Name: "マリオ", IsActive: false})
	if result.Error != nil {
		log.Fatal(result.Error)
	}
	log.Println("count:", result.RowsAffected)

	user := Paformer{}
	db.Where("id = 1").Take(&user)
	log.Println("user:", user)
}