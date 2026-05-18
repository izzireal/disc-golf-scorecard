package domain

// HoleScore represents the summary of what happened on one hole during a round
type HoleScore struct {
	ID         int    `json:"id" db:"id"`
	RoundID    int    `json:"round_id" db:"round_id"`
	PlayerID   int    `json:"player_id" db:"player_id"`
	HoleNumber int    `json:"hole_number" db:"hole_number"`
	TeeColor   string `json:"tee_color" db:"tee_color"`

	Score int    `json:"score" db:"score"`           // Final score for the hole (e.g. 3, 4, 5)
	Putts *int   `json:"putts,omitempty" db:"putts"` // optional
	Notes string `json:"notes,omitempty" db:"notes"`

	// Optional detailed throws (loaded only when user wants full history)
	Throws []Throw `json:"throws,omitempty"`
}

// NewHoleScore is used when creating a new score entry
type NewHoleScore struct {
	RoundID    int    `json:"round_id"`
	PlayerID   int    `json:"player_id"`
	HoleNumber int    `json:"hole_number"`
	TeeColor   string `json:"tee_color"`
	Score      int    `json:"score"`
	Putts      *int   `json:"putts,omitempty"`
	Notes      string `json:"notes,omitempty"`
}
