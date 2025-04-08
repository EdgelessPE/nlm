package db

import (
	"fmt"
	"nlm/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		config.ENV.DB_HOST,
		config.ENV.DB_USER,
		config.ENV.DB_PASSWORD,
		config.ENV.DB_NAME,
		config.ENV.DB_PORT)

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
