package models

// Eurovision represents the whole event with multiple contests.
type Eurovision struct {
    ID         uint    `json:"id" gorm:"primaryKey"`
    Year       int     `json:"year"`             // The year of the contest
    // Foreign key relationships to semi-finals and final
    SemiFinal1ID uint      `json:"semi_final_1_id"`
    SemiFinal1   Contest   `json:"semi_final_1" gorm:"foreignKey:SemiFinal1ID;references:ID"`

    SemiFinal2ID uint      `json:"semi_final_2_id"`
    SemiFinal2   Contest   `json:"semi_final_2" gorm:"foreignKey:SemiFinal2ID;references:ID"`

    FinalID      uint      `json:"final_id"`
    Final        Contest   `json:"final" gorm:"foreignKey:FinalID;references:ID"`
}
