package utils

import (
	"eurovision-simulator/models"
	"fmt"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

// InitializeDB sets up the database connection
func InitializeDB() *gorm.DB {
	if db == nil {
		var err error
		db, err = gorm.Open(sqlite.Open("eurovision.db"), &gorm.Config{})
		if err != nil {
			log.Fatal("Failed to connect to database:", err)
		}

		// Auto migrate the models (this will create/update tables based on your structs)
		err = db.AutoMigrate(&models.Eurovision{}, &models.Contest{}, &models.Contestant{}, &models.VotingResults{})
		if err != nil {
			log.Fatal("Failed to migrate database:", err)
		}

		fmt.Println("Database connection established and migrated")
	}
	return db
}
