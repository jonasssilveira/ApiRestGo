package database

import (
	"booksAPiProject/migrations"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"time"
)

var db *gorm.DB

func StardDB() *gorm.DB {
	str := "host=localhost port=25432 user=admin dbname=books sslmode=disable password=123456"
	db, err := gorm.Open(postgres.Open(str), &gorm.Config{})

	if err != nil {
		log.Fatal("Error: ", err)
	}

	config, _ := db.DB()

	config.SetMaxIdleConns(10)
	config.SetMaxOpenConns(100)
	config.SetConnMaxLifetime(time.Hour)

	migrations.RunMigrations(db)
	return db
}
func GetDatabase() *gorm.DB {
	return db
}
