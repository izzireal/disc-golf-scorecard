package domain

import "time"

// Round represents one complete disc golf session (18 holes usually)
type Round struct {
	ID         int       `json:"id" db:"id"`
	CourseID   int       `json:"course_id" db:"course_id"`
	DatePlayed time.Time `json:"date_played" db:"date_played"`
	GroupCode  string    `json:"group_code,omitempty" db:"group_code"`
	Notes      string    `json:"notes,omitempty" db:"notes"`
	CreatedAt  time.Time `json:"created_at" db:"created_at"`

	// Relationships (loaded when needed)
	Course       *Course            `json:"course,omitempty"`
	Participants []RoundParticipant `json:"participants,omitempty"`
}

// RoundParticipant links a player to a round
type RoundParticipant struct {
	ID         int  `json:"id" db:"id"`
	RoundID    int  `json:"round_id" db:"round_id"`
	PlayerID   int  `json:"player_id" db:"player_id"`
	FinalScore *int `json:"final_score,omitempty" db:"final_score"` // can be null until round finished
}
