package controllers

import (
	"eurovision-simulator/models"

	"gorm.io/gorm"
)

func populateContestants(db *gorm.DB, contestants []string) ([]models.Contestant, error) {
	eligibleContestants := make([]models.Contestant, 0)
	for _, country := range contestants {
		contestant := models.Contestant{
			Name:    country,
			Country: country,
		}
		eligibleContestants = append(eligibleContestants, contestant)
		err := db.Create(&contestant).Error
		if err != nil {
			return nil, err
		}
	}
	return eligibleContestants, nil
}
