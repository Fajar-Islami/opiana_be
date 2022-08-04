package config

import (
	"fmt"
	"os"

	"opiana_code_test/entity"
	"opiana_code_test/seed"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// SetupDatabaseConnection is creating a new connection to our database
func SetupDatabaseConnection() *gorm.DB {
	err := godotenv.Load()
	if err != nil {
		panic("Failed to load env file")
	}

	dbUser := os.Getenv("mysql.DB_USER")
	dbPass := os.Getenv("mysql.DB_PASS")
	dbHost := os.Getenv("mysql.DB_HOST")
	dbName := os.Getenv("mysql.DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to create a connection to your database")
	}

	db.AutoMigrate(&entity.Post{}, &entity.User{}, &entity.PostType{})

	// Seed data
	var count int64
	if db.Migrator().HasTable(&entity.PostType{}) {
		// count row
		db.Model(&entity.PostType{}).Count(&count)
		if count < 1 {
			db.CreateInBatches(&seed.PostTypeSeed, len(seed.PostTypeSeed))
		}
	}

	return db

}

// CloseDatabaseConnection method is closing a connection between your app and your db
func CloseDatabaseConnection(db *gorm.DB) {
	dbSQL, err := db.DB()
	if err != nil {
		panic("Failed to close connection to database")
	}

	dbSQL.Close()

}
