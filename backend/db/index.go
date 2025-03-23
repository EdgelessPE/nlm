package db

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"))

	config := &gorm.Config{
		PrepareStmt: true,
	}

	db, err := gorm.Open(postgres.Open(dsn), config)
	if err != nil {
		panic("failed to connect database")
	}

	// 启用 UUID 扩展
	db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";")

	DB = db
}
