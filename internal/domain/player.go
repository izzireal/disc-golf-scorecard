package domain

import "time"

// Player represents a person who plays disc golf
type Player struct {
	ID        int       `json:"id" db:"id"`
	Name      string    `json:"name" db:"name"`
	Nickname  string    `json:"nickname,omitempty" db:"nickname"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

// PlayerSummary is a lighter version used in leaderboards and lists
type PlayerSummary struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Nickname string `json:"nickname,omitempty"`
}
