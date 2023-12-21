package db

import (
	"log"

	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	db, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=your_dbname password=your_password sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	DB = db
	db.AutoMigrate(&BinanceData{})

	// Важно: Не забудьте вернуть открытое соединение

}
