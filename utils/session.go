package utils

import "eurovision-simulator/models"

// Global Eurovision event variable
var eurovisionEvent *models.Eurovision

// GetEurovisionSession returns the current Eurovision session or initializes a new one
func GetEurovisionSession() *models.Eurovision {
	if eurovisionEvent == nil {
		eurovisionEvent = &models.Eurovision{Year: 2024}
	}
	return eurovisionEvent
}
