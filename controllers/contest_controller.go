package controllers

import (
	"errors"
	"eurovision-simulator/models"
	"log"
	"math/rand"
	"time"

	"gorm.io/gorm"
)

// populateSemiFinals takes a list of contestants, splits them into two semi-finals,
// and stores them in the database.
func populateSemiFinals(db *gorm.DB, contestants []models.Contestant) (*models.Contest, *models.Contest, error) {
	// Validate input
	if len(contestants) == 0 {
		return nil, nil, errors.New("no contestants available to populate semi-finals")
	}

	rand.Seed(time.Now().UnixNano()) // Initialize random seed

	// Shuffle contestants to simulate a random voting preference
	shuffleContestants(contestants)

	// Split contestants into two groups
	sf1Contestants, sf2Contestants := splitContestants(contestants)

	// Create the semi-final contests
	sf1, err := createContest(db, "Semi-Final 1", sf1Contestants)
	if err != nil {
		log.Printf("Error creating Semi-Final 1: %v", err)
		return nil, nil, err
	}

	sf2, err := createContest(db, "Semi-Final 2", sf2Contestants)
	if err != nil {
		log.Printf("Error creating Semi-Final 2: %v", err)
		return nil, nil, err
	}

	return sf1, sf2, nil
}

// shuffleContestants shuffles the order of the contestants.
func shuffleContestants(contestants []models.Contestant) {
	log.Printf("Shuffling %d contestants", len(contestants))
	rand.Shuffle(len(contestants), func(i, j int) {
		contestants[i], contestants[j] = contestants[j], contestants[i]
	})
}

// splitContestants splits the contestants into two equal parts.
// If there is an odd number, the second group will have one extra contestant.
func splitContestants(contestants []models.Contestant) ([]models.Contestant, []models.Contestant) {
	midpoint := len(contestants) / 2
	log.Printf("Splitting contestants into two groups: %d in Semi-Final 1, %d in Semi-Final 2",
		len(contestants[:midpoint]), len(contestants[midpoint:]))

	return contestants[:midpoint], contestants[midpoint:]
}

// createContest creates a new contest (either Semi-Final or Final) with the given contestants.
func createContest(db *gorm.DB, name string, contestants []models.Contestant) (*models.Contest, error) {
	contest := models.Contest{
		Name:        name,
		ContestType: "SemiFinal", // Assuming it's a semi-final here; adjust if needed
		Contestants: contestants,
	}

	log.Printf("Creating %s with %d contestants", name, len(contestants))
	if err := db.Create(&contest).Error; err != nil {
		return nil, err
	}

	return &contest, nil
}
