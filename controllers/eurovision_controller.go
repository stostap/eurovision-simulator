package controllers

import (
	"eurovision-simulator/models"

	"gorm.io/gorm"
)

type EurovisionController struct {
	DB *gorm.DB
}

// NewEurovisionController creates a new instance of EurovisionController
func NewEurovisionController(db *gorm.DB) *EurovisionController {
	return &EurovisionController{DB: db}
}

// GetEurovision fetches the Eurovision event details (including semi-finals and final)
func (controller *EurovisionController) GetEurovision() (*models.Eurovision, error) {
	var eurovision models.Eurovision

	if err := controller.DB.Preload("SemiFinal1.Contestants").Preload("SemiFinal2.Contestants").
		Preload("Final.Contestants").First(&eurovision).Error; err != nil {
		return nil, err
	}

	return &eurovision, nil
}

// StartEurovision initializes the Eurovision event with two semi-finals and a final
func (controller *EurovisionController) StartEurovision() (*models.Eurovision, error) {
	eurovision := models.Eurovision{
		Year: 2024,
	}

	countries := []string{
		"Albania",
		"Armenia",
		"Australia",
		"Austria",
		"Azerbaijan",
		"Belgium",
		"Bulgaria",
		"Croatia",
		"Cyprus",
		"Czech Republic",
		"Denmark",
		"Estonia",
		"Finland",
		"France",
		"Georgia",
		"Germany",
		"Greece",
		"Iceland",
		"Ireland",
		"Israel",
		"Italy",
		"Latvia",
		/*"Lithuania",
		"Malta",
		"Moldova",
		"Netherlands",
		"North Macedonia",
		"Norway",
		"Poland",
		"Portugal",
		"Romania",
		"San Marino",
		"Serbia",
		"Slovenia",
		"Spain",
		"Sweden",
		"Switzerland",
		"Ukraine",
		"United Kingdom",*/
	}

	contestants, err := populateContestants(controller.DB, countries)
	if err != nil {
		return nil, err
	}

	sf1, sf2, err := populateSemiFinals(controller.DB, contestants)
	eurovision.SemiFinal1 = *sf1
	eurovision.SemiFinal2 = *sf2

	if err := controller.DB.Create(&eurovision).Error; err != nil {
		return nil, err
	}

	return &eurovision, nil
}

// GetVotingResults fetches the voting results for the final stage of Eurovision
func (controller *EurovisionController) GetVotingResults() ([]models.VotingResults, error) {
	var eurovision models.Eurovision

	// Fetch the Eurovision final voting results with contestants and their votes
	if err := controller.DB.Preload("Final.VotingResults").First(&eurovision).Error; err != nil {
		return nil, err
	}

	return eurovision.Final.Voting, nil
}
