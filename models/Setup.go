package models

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)
import _ "gorm.io/driver/sqlite"

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open(sqlite.Open("test.database"), &gorm.Config{})

	if err != nil {
		panic("Cannot connect to database!")
	}

	database.AutoMigrate(&Book{})

	DB = database
}
