package database

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DbInstance struct {
	db *gorm.DB
}

var Database DbInstance

func ConnectDb() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}
	log.Println("Connection Opened to Database")
	db.Logger = logger.Default.LogMode(logger.Info)

	log.Println("Database Migrated")
	db.AutoMigrate(&Customer{}, &Token{})

	Database = DbInstance{db}
}
