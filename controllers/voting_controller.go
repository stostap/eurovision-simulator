package controllers

import (
	"eurovision-simulator/models"
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

func (sim *VotingSimulator) SimulateVoting(c *models.Contest) error {
	rand.Seed(time.Now().UnixNano())

	// Points for the top 10 contestants
	pointsDistribution := []int{12, 10, 8, 7, 6, 5, 4, 3, 2, 1}

	for _, voter := range c.Contestants {
		vr := models.VotingResults{
			ContestantID:            voter.ID,
			ContestID:               c.ID,
			JuryVotesByContestant:   make(map[uint]int),
			PublicVotesByContestant: make(map[uint]int)}
		c.Voting = append(c.Voting, vr)
		/*if err := sim.DB.Create(&vr).Error; err != nil {
			return err
		}*/
	}

	for _, voter := range c.Contestants {
		var candidates []models.Contestant
		for _, can := range c.Contestants {
			if voter.ID != can.ID {
				candidates = append(candidates, can)
			}
		}

		rand.Shuffle(len(candidates), func(i, j int) {
			candidates[i], candidates[j] = candidates[j], candidates[i]
		})

		// Select top 10 contestants for jury voting
		topContestantsForJury := candidates[:10]

		for i, pc := range topContestantsForJury {
			idx := slices.IndexFunc(c.Voting, func(vr models.VotingResults) bool {
				return pc.ID == vr.ContestantID
			})
			c.Voting[idx].JuryVotes += pointsDistribution[i]
			c.Voting[idx].TotalScore += pointsDistribution[i]
			c.Voting[idx].JuryVotesByContestant[voter.ID] = pointsDistribution[i]
		}

		rand.Shuffle(len(candidates), func(i, j int) {
			candidates[i], candidates[j] = candidates[j], candidates[i]
		})

		// Select top 10 contestants for public voting
		topContestantsForPublic := candidates[:10]

		for i, pc := range topContestantsForPublic {
			idx := slices.IndexFunc(c.Voting, func(vr models.VotingResults) bool {
				return pc.ID == vr.ContestantID
			})
			c.Voting[idx].PublicVotes += pointsDistribution[i]
			c.Voting[idx].TotalScore += pointsDistribution[i]
			c.Voting[idx].PublicVotesByContestant[voter.ID] = pointsDistribution[i]
		}

	}

	if err := sim.DB.Create(&c.Voting).Error; err != nil {
		return err
	}
	//sim.DB.Model(&c).Update("voting", c.Voting)
	return nil
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
