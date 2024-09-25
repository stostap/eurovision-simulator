package main

import (
	"eurovision-simulator/controllers"
	"eurovision-simulator/models"
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Function to create a Eurovision event for a specific year
func createEurovisionEvent(db *gorm.DB) {
	// Create SemiFinal 1
	semiFinal1 := models.Contest{
		Name:        "SemiFinal 1",
		ContestType: "SemiFinal",
	}

	// Create SemiFinal 2
	semiFinal2 := models.Contest{
		Name:        "SemiFinal 2",
		ContestType: "SemiFinal",
	}

	// Create Final
	finalContest := models.Contest{
		Name:        "Final",
		ContestType: "Final",
	}

	// Create the Eurovision event for a specific year
	eurovisionEvent := models.Eurovision{
		Year:       2024,
		SemiFinal1: semiFinal1,
		SemiFinal2: semiFinal2,
		Final:      finalContest,
	}

	// Save Eurovision event, along with contests (SemiFinal 1, SemiFinal 2, Final)
	db.Create(&eurovisionEvent)
}

func main() {
	// Initialize GORM database connection (using SQLite for simplicity)
	db, err := gorm.Open(sqlite.Open("eurovision.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect to the database")
	}

	// AutoMigrate all models (Contestant, Contest, VotingResults, Eurovision)
	db.AutoMigrate(&models.Contestant{}, &models.Contest{}, &models.VotingResults{}, &models.Eurovision{})

	eurovisionCtrl := controllers.NewEurovisionController(db)

	event, err := eurovisionCtrl.StartEurovision()
	votingSim := controllers.NewVotingSimulator(db)
	votingSim.SimulateVoting(&event.SemiFinal1)

	fmt.Printf("Eurovision %d:\n", event.Year)
	fmt.Printf("  SemiFinal 1: %v\n", event.SemiFinal1)
	fmt.Printf("  SemiFinal 2: %v\n", event.SemiFinal2)
	fmt.Printf("  Final: %s\n", event.Final.Name)
}
