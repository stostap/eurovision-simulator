package models

// Contest represents either a SemiFinal or Final stage.
type Contest struct {
	ID            uint            `json:"id" gorm:"primaryKey"`
	Name          string          `json:"name"`                                              // E.g., "SemiFinal 1", "SemiFinal 2", "Final"
	ContestType   string          `json:"contest_type"`                                      // E.g., "SemiFinal" or "Final"
	Contestants   []Contestant    `json:"contestants" gorm:"many2many:contest_contestants;"` // Many-to-many relation with Contestant
	VotingResults []VotingResults `json:"voting_results" gorm:"foreignKey:ContestID"`
}
