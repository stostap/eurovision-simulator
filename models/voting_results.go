package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

// Custom type for storing map[uint]int as JSON in the database
type UintIntMap map[uint]int

// Implement the driver.Valuer interface for saving the map as JSON
func (m UintIntMap) Value() (driver.Value, error) {
	// Marshal the map into JSON format
	return json.Marshal(m)
}

// Implement the sql.Scanner interface for reading the map from JSON
func (m *UintIntMap) Scan(value interface{}) error {
	if value == nil {
		*m = make(map[uint]int) // Initialize an empty map if no value is found
		return nil
	}

	// Convert the value (which is stored as []byte in the DB) into JSON
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("failed to convert database value to byte array")
	}

	// Unmarshal the JSON into the map
	return json.Unmarshal(bytes, m)
}

// VotingResults tracks votes given to a contestant in a specific contest.
type VotingResults struct {
	ID           uint `json:"id" gorm:"primaryKey"`
	ContestantID uint `json:"contestant_id" gorm:"foreignKey"` // The contestant receiving the votes
	ContestID    uint `json:"contest_id" gorm:"foreignKey"`    // SemiFinal or Final contest ID
	JuryVotes    int  `json:"jury_votes"`                      // Total jury votes
	PublicVotes  int  `json:"public_votes"`                    // Total public votes
	TotalScore   int  `json:"total_score"`                     // Total score (jury + public)

	// Maps storing how each contestant voted for this contestant (Jury/Public)
	JuryVotesByContestant   UintIntMap `json:"jury_votes_by_contestant" gorm:"type:json"`
	PublicVotesByContestant UintIntMap `json:"public_votes_by_contestant" gorm:"type:json"`
}

// BeforeSave serializes maps into JSON format before saving to the DB.
/*func (v *VotingResults) BeforeSave(tx *gorm.DB) (err error) {
	fmt.Println("error here\n")
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
}*/
