package models

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

var DB *gorm.DB

func connectDatabase() {
	dsn := fmt.Sprintf(
		"user=%s paassword=%s dbname=%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
	)
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connetc to database")
	}

	database.AutoMigrate(&Head{}, &Page{})

	DB = database
}



