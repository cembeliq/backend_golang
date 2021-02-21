package database

import (
	"cembeliq_app/config"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {
	db, err := gorm.Open(sqlite.Open(config.EnvVariable("DATABASE_NAME")), &gorm.Config{})
	// db2, _ := db.DB()
	// db2.Close()
	if err != nil {
		panic("failed to connect database")
	}
	log.Print("Connection successfully...")
	return db
}
