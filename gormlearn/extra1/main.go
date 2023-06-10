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

	log.Println("create db: ", Db)
}

type Paformer struct {
	gorm.Model
	Name     string
	Email    string
	Age      int
	IsActive bool
}