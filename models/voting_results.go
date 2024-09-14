package models

import (
    "gorm.io/gorm"
    "encoding/json"
)

// VotingResults tracks votes given to a contestant in a specific contest.
type VotingResults struct {
    ID                    uint            `json:"id" gorm:"primaryKey"`
    ContestantID          uint            `json:"contestant_id" gorm:"foreignKey"` // The contestant receiving the votes
    ContestID             uint            `json:"contest_id" gorm:"foreignKey"`    // SemiFinal or Final contest ID
    JuryVotes             int             `json:"jury_votes"`                      // Total jury votes
    PublicVotes           int             `json:"public_votes"`                    // Total public votes
    TotalScore            int             `json:"total_score"`                     // Total score (jury + public)

    // Maps storing how each contestant voted for this contestant (Jury/Public)
    JuryVotesByContestant   map[uint]int   `json:"jury_votes_by_contestant" gorm:"type:json"`
    PublicVotesByContestant map[uint]int   `json:"public_votes_by_contestant" gorm:"type:json"`
}

// BeforeSave serializes maps into JSON format before saving to the DB.
func (v *VotingResults) BeforeSave(tx *gorm.DB) (err error) {
    if len(v.JuryVotesByContestant) > 0 {
        juryVotesJSON, err := json.Marshal(v.JuryVotesByContestant)
        if err != nil {
            return err
        }
        v.JuryVotesByContestant = map[uint]int{}
        json.Unmarshal(juryVotesJSON, &v.JuryVotesByContestant)
    }

    if len(v.PublicVotesByContestant) > 0 {
        publicVotesJSON, err := json.Marshal(v.PublicVotesByContestant)
        if err != nil {
            return err
        }
        v.PublicVotesByContestant = map[uint]int{}
        json.Unmarshal(publicVotesJSON, &v.PublicVotesByContestant)
    }

    return nil
}
