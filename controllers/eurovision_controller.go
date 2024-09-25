package controllers

import (
	"eurovision-simulator/models"
	"log"

	"gorm.io/gorm"
)

// EurovisionController handles Eurovision-related operations
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

	// Preload semi-final and final contestants
	if err := controller.DB.
		Preload("SemiFinal1.Contestants").
		Preload("SemiFinal2.Contestants").
		Preload("Final.Contestants").
		First(&eurovision).Error; err != nil {
		log.Printf("Error fetching Eurovision details: %v", err)
		return nil, err
	}

	return &eurovision, nil
}

// StartEurovision initializes the Eurovision event with two semi-finals and a final
func (controller *EurovisionController) StartEurovision() (*models.Eurovision, error) {
	eurovision := models.Eurovision{Year: 2024}

	// Generate contestant list from predefined countries
	countries := generateCountriesList()

	// Populate contestants
	contestants, err := populateContestants(controller.DB, countries)
	if err != nil {
		log.Printf("Error populating contestants: %v", err)
		return nil, err
	}

	// Populate semi-finals
	sf1, sf2, err := populateSemiFinals(controller.DB, contestants)
	if err != nil {
		log.Printf("Error populating semi-finals: %v", err)
		return nil, err
	}

	// Assign semi-finals to Eurovision struct
	eurovision.SemiFinal1 = *sf1
	eurovision.SemiFinal2 = *sf2

	// Save Eurovision instance to the database
	if err := controller.DB.Create(&eurovision).Error; err != nil {
		log.Printf("Error creating Eurovision event: %v", err)
		return nil, err
	}

	log.Println("Eurovision event successfully created")
	return &eurovision, nil
}

// GetVotingResults fetches the voting results for the final stage of Eurovision
func (controller *EurovisionController) GetVotingResults() ([]models.VotingResults, error) {
	var eurovision models.Eurovision

	// Fetch voting results for the final
	if err := controller.DB.Preload("Final.VotingResults").First(&eurovision).Error; err != nil {
		log.Printf("Error fetching voting results: %v", err)
		return nil, err
	}

	return eurovision.Final.Voting, nil
}

// generateCountriesList returns a list of countries participating in Eurovision
func generateCountriesList() []string {
	return []string{
		"Albania", "Armenia", "Australia", "Austria", "Azerbaijan", "Belgium", "Bulgaria", "Croatia",
		"Cyprus", "Czech Republic", "Denmark", "Estonia", "Finland", "France", "Georgia", "Germany",
		"Greece", "Iceland", "Ireland", "Israel", "Italy", "Latvia", "Lithuania", "Malta", "Moldova",
		"Netherlands", "North Macedonia", "Norway", "Poland", "Portugal", "Romania", "San Marino",
		"Serbia", "Slovenia", "Spain", "Sweden", "Switzerland", "Ukraine", "United Kingdom",
	}
}
