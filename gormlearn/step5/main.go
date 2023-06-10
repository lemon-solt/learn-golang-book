package main

import (
	"fmt"
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
	Db.AutoMigrate(&Paformer{})

	// user := []Paformer{
	// 	{Name: "サカズキ", Email: "hoge@example.com"},
	// }
	// Db.Create(&user)

	// updates(Db)
	updatesAll(Db)
}

type Paformer struct {
    gorm.Model
    Name  string
    Email string
	Age int
	IsActive bool
}

// 複数のカラムを更新する
func updates(db *gorm.DB) {
	result := db.Model(&Paformer{}).Where("id = 1").Updates(Paformer{Name: "サボ", Age: 10, IsActive: true})
	if result.Error != nil {
		log.Fatal(result.Error)
	}
	log.Println("count:", result.RowsAffected)

	// user := Paformer{}
	// db.Where("id = 1").Take(&user)
	// log.Println("user:", user)
}

// 一括更新
func updatesAll(db *gorm.DB) {
	paformer := Paformer{
		Name:     "海賊王ルフィ",
		Age:      100,
		IsActive: true,
	}
	result := db.Where("name = ?", "jon").Updates(&paformer)
	if result.Error != nil {
		log.Fatal(result.Error)
	}
	log.Println("count:", result.RowsAffected)

	paformers := []Paformer{}
	db.Find(&paformers)
	fmt.Println("paformers:", paformers)
}