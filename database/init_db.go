package database

import (
	"fmt"
	"golang_api/models"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

//ConnectDB func
func ConnectDB() (*gorm.DB, error) {
	//Load env variables
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	databaseHost := os.Getenv("databaseHost")
	databasePort := os.Getenv("databasePort")
	databaseName := os.Getenv("databaseName")
	username := os.Getenv("databaseUser")
	password := os.Getenv("databasePassword")

	dbURI := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=disable", databaseHost, databasePort, databaseName, username, password)
	db, err := gorm.Open(postgres.Open(dbURI), &gorm.Config{})

	if err != nil {
		fmt.Println("error", err)
		panic(err)
	}

	runMigrations(db)
	return db, err
}

func runMigrations(db *gorm.DB) {
	db.AutoMigrate(
		&models.User{},
		&models.Book{},
	)
}

var conn, _ = ConnectDB()
