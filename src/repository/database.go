package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"urlshortener.com/devgym/jr/models"
)

var DB *gorm.DB

// type Connection struct {
// 	Host         string
// 	Port         int32
// 	User         string
// 	Password     string
// 	DatabaseName string
// }

func Init() *gorm.DB {
	//dsn := "host=localhost port=5432 user=postgres password=postgres dbname=shorten_url sslmode=disable"
	dsn := "postgres://postgres:postgres@localhost:5432/devgym_shorten_url"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("I got a problen when initializing database")
	}

	db.AutoMigrate(&models.ShortenUrl{})

	return db
}
