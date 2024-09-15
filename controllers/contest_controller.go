package controllers

import (
	"eurovision-simulator/models"
	"time"

	"math/rand"

	"gorm.io/gorm"
)

func populateSemiFinals(db *gorm.DB, contestants []models.Contestant) (*models.Contest, *models.Contest, error) {
	rand.Seed(time.Now().UnixNano()) // Initialize random seed
	// Shuffle contestants to simulate a random voting preference
	rand.Shuffle(len(contestants), func(i, j int) {
		contestants[i], contestants[j] = contestants[j], contestants[i]
	})

	midpoint := len(contestants) / 2
	sf1 := models.Contest{
		Name:        "Semi-Final 1",
		ContestType: "SemiFinal",
		Contestants: contestants[:midpoint],
	}
	err := db.Create(&sf1).Error
	if err != nil {
		return nil, nil, err
	}

	sf2 := models.Contest{
		Name:        "Semi-Final 2",
		ContestType: "SemiFinal",
		Contestants: contestants[midpoint:],
	}
	err = db.Create(&sf2).Error
	if err != nil {
		return nil, nil, err
	}

	return &sf1, &sf2, nil
}
