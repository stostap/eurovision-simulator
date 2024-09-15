package router

import (
    "eurovision-simulator/controllers"
    "github.com/gin-gonic/gin"
)

// SetupRouter sets up the routes and assigns them to the appropriate controllers
func SetupRouter(db *gorm.DB, votingSimulator *controllers.VotingSimulator) *gin.Engine {
    router := gin.Default()

    // Route to simulate voting for a contest
    router.POST("/simulate-voting/:contestID", func(c *gin.Context) {
        contestID := c.Param("contestID")
        
        // Simulate voting for the contest (just an example: fetch contestants for the contest)
        contestants := []models.Contestant{/* Fetch contestants from DB */}
        
        err := votingSimulator.SimulateVoting(contestID, contestants)
        if err != nil {
            c.JSON(500, gin.H{"error": err.Error()})
            return
        }
        c.JSON(200, gin.H{"message": "Voting simulation completed"})
    })

    return router
}
