package controllers

import (
	"eurovision-simulator/models"
	"log"

	"gorm.io/gorm"
)

// populateContestants adds a list of contestants to the database.
// It takes a list of country names and adds each country as a contestant.
// Returns a slice of models.Contestant or an error if the operation fails.
func populateContestants(db *gorm.DB, contestants []string) ([]models.Contestant, error) {
	if len(contestants) == 0 {
		return nil, nil // Return early if there are no contestants
	}

	eligibleContestants := make([]models.Contestant, len(contestants))

	// Create the contestant objects
	for i, country := range contestants {
		eligibleContestants[i] = models.Contestant{
			Name:    country,
			Country: country,
		}
	}

	// Insert contestants in bulk for better performance
	if err := db.Create(&eligibleContestants).Error; err != nil {
		log.Printf("Error creating contestants: %v", err)
		return nil, err
	}

	log.Printf("Successfully populated %d contestants", len(eligibleContestants))
	return eligibleContestants, nil
}
