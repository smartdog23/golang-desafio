package utils

import (
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"github.com/smartdog23/golang-desafio/domain"
	"log"
	"os"
	_ "github.com/lib/pq"
)


func ConnectDB() *gorm.DB {

	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	dsn := os.Getenv("dsn")

	db, err := gorm.Open("postgres", dsn)

	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
		panic(err)
	}

	//defer db.Close()

	db.AutoMigrate(&domain.User{})

	return db
}
