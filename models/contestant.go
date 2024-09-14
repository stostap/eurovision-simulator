package models

// Contestant represents a participant in the contest.
type Contestant struct {
    ID        uint   `json:"id" gorm:"primaryKey"`
    Name      string `json:"name"`
    Country   string `json:"country"`
}