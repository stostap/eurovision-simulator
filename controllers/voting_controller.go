package controllers

import (
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

// SimulateVoting simulates both jury and public voting for a given contest
// where contestants themselves are the voters, but they cannot vote for themselves.
/*func (sim *VotingSimulator) SimulateVoting(contestID uint, contestants []models.Contestant) error {
	rand.Seed(time.Now().UnixNano()) // Initialize random seed

	// Points for the top 10 contestants
	pointsDistribution := []int{12, 10, 8, 7, 6, 5, 4, 3, 2, 1}

	for _, voter := range contestants {
		// Shuffle contestants to simulate a random voting preference
		rand.Shuffle(len(contestants), func(i, j int) {
			contestants[i], contestants[j] = contestants[j], contestants[i]
		})

		// Filter out the voter from the shuffled list (so they can't vote for themselves)
		eligibleContestants := make([]models.Contestant, 0)
		for _, contestant := range contestants {
			if contestant.ID != voter.ID {
				eligibleContestants = append(eligibleContestants, contestant)
			}
		}

		// Select top 10 contestants for jury voting
		topContestantsForJury := eligibleContestants[:10]

		// Shuffle the list again for public voting
		rand.Shuffle(len(eligibleContestants), func(i, j int) {
			eligibleContestants[i], eligibleContestants[j] = eligibleContestants[j], eligibleContestants[i]
		})

		// Select top 10 contestants for public voting
		topContestantsForPublic := eligibleContestants[:10]

		// Generate jury votes for this voter (acting as a "country")
		juryVotes := make(map[uint]int)
		publicVotes := make(map[uint]int)

		// Assign jury votes
		for i, contestant := range topContestantsForJury {
			juryVotes[contestant.ID] = pointsDistribution[i]
		}

		// Assign public votes (this can be different, we will reuse the same distribution for now)
		for i, contestant := range topContestantsForPublic {
			publicVotes[contestant.ID] = pointsDistribution[i]
		}

		// Store results for each contestant
		for contestantID, juryPoints := range juryVotes {
			publicPoints := publicVotes[contestantID] // Public points may be zero if contestant is not in top 10 for public votes

			// Find or create VotingResults for the contestant
			votingResult := models.VotingResults{
				ContestID:    contestID,
				ContestantID: contestantID,
				JuryVotes:    map[uint]int{voter.ID: juryPoints},
				PublicVotes:  map[uint]int{voter.ID: publicPoints},
				TotalScore:   juryPoints + publicPoints,
			}
			sim.DB.Create(&votingResult)
		}
	}
	return nil
}*/
