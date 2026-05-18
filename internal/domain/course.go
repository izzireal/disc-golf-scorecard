package domain

import "time"

// Course represents a disc golf course (the physical place)
type Course struct {
	ID        int       `json:"id" db:"id"`
	Name      string    `json:"name" db:"name"`
	Location  string    `json:"location" db:"location"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`

	// We can load holes later when needed
	Holes []CourseHole `json:"holes,omitempty"`
}
