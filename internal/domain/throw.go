package domain

// Throw represents one individual throw during a hole
type Throw struct {
	ID          int `json:"id" db:"id"`
	HoleScoreID int `json:"hole_score_id" db:"hole_score_id"`
	ThrowNumber int `json:"throw_number" db:"throw_number"` // 1 = drive, 2 = upshot, etc.

	DiscUsed  string `json:"disc_used,omitempty" db:"disc_used"`
	ThrowType string `json:"throw_type,omitempty" db:"throw_type"` // Backhand, Forehand, Roller, etc.
	Outcome   string `json:"outcome,omitempty" db:"outcome"`       // Fairway Hit, OB, Tree, Approach, In Basket, etc.

	DistanceFeet *int   `json:"distance_feet,omitempty" db:"distance_feet"`
	Notes        string `json:"notes,omitempty" db:"notes"`
}

// NewThrow is used when recording a new throw
type NewThrow struct {
	HoleScoreID  int    `json:"hole_score_id"`
	ThrowNumber  int    `json:"throw_number"`
	DiscUsed     string `json:"disc_used,omitempty"`
	ThrowType    string `json:"throw_type,omitempty"`
	Outcome      string `json:"outcome,omitempty"`
	DistanceFeet *int   `json:"distance_feet,omitempty"`
	Notes        string `json:"notes,omitempty"`
}
