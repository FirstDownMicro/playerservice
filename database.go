package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func CreateConnection() (*gorm.DB, error) {
	// host := os.Getenv("DB_HOST")
	// user := os.Getenv("DB_USER")
	// DBName := os.Getenv("DB_NAME")
	// password := os.Getenv("DB_PASSWORD")

	//docker.for.mac.localhost
	return gorm.Open(
		"postgres",
		"host=docker.for.mac.localhost port=5432 user=jessetellez dbname=footballdb password=password1 sslmode=disable",
	)
}
