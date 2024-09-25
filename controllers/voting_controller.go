package controllers

import (
	"errors"
	"eurovision-simulator/models"
	"log"
	"math/rand"
	"slices"
	"time"

	"gorm.io/gorm"
)

// VotingSimulator handles the voting simulation logic
type VotingSimulator struct {
	DB *gorm.DB
}

// NewVotingSimulator creates a new VotingSimulator
func NewVotingSimulator(db *gorm.DB) *VotingSimulator {
	return &VotingSimulator{DB: db}
}

// SimulateVoting performs the voting simulation for a given contest
func (sim *VotingSimulator) SimulateVoting(c *models.Contest) error {
	startTime := time.Now() // Track the start time for performance metrics
	rand.Seed(time.Now().UnixNano())

	log.Printf("Starting voting simulation for contest: %s (ID: %d)", c.Name, c.ID)

	if len(c.Contestants) <= 1 {
		log.Println("Error: Not enough contestants to perform voting.")
		return errors.New("not enough contestants to perform voting")
	}

	// Points for the top 10 contestants
	pointsDistribution := []int{12, 10, 8, 7, 6, 5, 4, 3, 2, 1}

	// Initialize VotingResults for each contestant
	if err := sim.initializeVotingResults(c); err != nil {
		log.Printf("Error initializing voting results for contest: %s (ID: %d), Error: %v", c.Name, c.ID, err)
		return err
	}

	for _, voter := range c.Contestants {
		log.Printf("Processing votes from voter: %s (ID: %d)", voter.Name, voter.ID)

		// Create a slice of candidates excluding the voter themselves
		candidates := sim.excludeSelf(c.Contestants, voter.ID)

		// Shuffle and calculate Jury votes
		log.Printf("Distributing Jury votes from %s (ID: %d)", voter.Name, voter.ID)
		sim.processVotes(c, candidates, voter, pointsDistribution, true)

		// Shuffle and calculate Public votes
		log.Printf("Distributing Public votes from %s (ID: %d)", voter.Name, voter.ID)
		sim.processVotes(c, candidates, voter, pointsDistribution, false)
	}

	// Save Voting results to the database
	log.Printf("Saving voting results for contest: %s (ID: %d) to the database", c.Name, c.ID)
	if err := sim.DB.Create(&c.Voting).Error; err != nil {
		log.Printf("Error saving voting results for contest: %s (ID: %d), Error: %v", c.Name, c.ID, err)
		return err
	}

	log.Printf("Voting simulation for contest: %s (ID: %d) completed in %v", c.Name, c.ID, time.Since(startTime))
	return nil
}

// initializeVotingResults initializes empty VotingResults for each contestant
func (sim *VotingSimulator) initializeVotingResults(c *models.Contest) error {
	log.Printf("Initializing voting results for contest: %s (ID: %d)", c.Name, c.ID)
	for _, voter := range c.Contestants {
		vr := models.VotingResults{
			ContestantID:            voter.ID,
			ContestID:               c.ID,
			JuryVotesByContestant:   make(map[uint]int),
			PublicVotesByContestant: make(map[uint]int),
		}
		c.Voting = append(c.Voting, vr)
	}
	return nil
}

// excludeSelf filters out the contestant voting from the candidate pool
func (sim *VotingSimulator) excludeSelf(contestants []models.Contestant, voterID uint) []models.Contestant {
	log.Printf("Excluding voter (ID: %d) from candidate pool", voterID)
	var candidates []models.Contestant
	for _, can := range contestants {
		if can.ID != voterID {
			candidates = append(candidates, can)
		}
	}
	rand.Shuffle(len(candidates), func(i, j int) {
		candidates[i], candidates[j] = candidates[j], candidates[i]
	})
	return candidates
}

// processVotes handles the logic for voting (either Jury or Public) for a set of contestants
func (sim *VotingSimulator) processVotes(c *models.Contest, candidates []models.Contestant, voter models.Contestant, pointsDistribution []int, isJury bool) {
	voteType := "Jury"
	if !isJury {
		voteType = "Public"
	}

	log.Printf("Processing %s votes from voter: %s (ID: %d)", voteType, voter.Name, voter.ID)

	// Take the top 10 candidates
	topContestants := candidates[:min(len(candidates), 10)]

	// Assign points based on the points distribution
	for i, pc := range topContestants {
		idx := slices.IndexFunc(c.Voting, func(vr models.VotingResults) bool {
			return pc.ID == vr.ContestantID
		})

		// Apply points to either Jury or Public voting
		if isJury {
			c.Voting[idx].JuryVotes += pointsDistribution[i]
			c.Voting[idx].TotalScore += pointsDistribution[i]
			c.Voting[idx].JuryVotesByContestant[voter.ID] = pointsDistribution[i]
			log.Printf("%s voted for %s (ID: %d) with %d Jury points", voter.Name, pc.Name, pc.ID, pointsDistribution[i])
		} else {
			c.Voting[idx].PublicVotes += pointsDistribution[i]
			c.Voting[idx].TotalScore += pointsDistribution[i]
			c.Voting[idx].PublicVotesByContestant[voter.ID] = pointsDistribution[i]
			log.Printf("%s voted for %s (ID: %d) with %d Public points", voter.Name, pc.Name, pc.ID, pointsDistribution[i])
		}
	}
}

// min returns the smaller of two integers
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
