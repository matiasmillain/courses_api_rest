package config

import (
	"fmt"
	"os"

	"github.com/matiasmillain/courses_api_rest/entity"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//SetupDatabaseConnection is creating a new connection to our database
func SetupDatabaseConnection() *gorm.DB {
	// errEnv := godotenv.Load()
	// if errEnv != nil {
	// 	panic("Failed to load env file")
	// }

	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf("host=%s user=%s password=%s port=%s dbname=%s sslmode=require TimeZone=America/Buenos_Aires", dbHost, dbUser, dbPass, dbPort, dbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to create a connection to database")
	}

	db.AutoMigrate(&entity.Course{}, &entity.User{})
	println("Database connected!")
	return db
}

//CloseDatabaseConnection method is closing a connection between your app and your db
func CloseDatabaseConnection(db *gorm.DB) {
	dbSQL, err := db.DB()
	if err != nil {
		panic("Failed to close connection from database")
	}
	dbSQL.Close()
}
