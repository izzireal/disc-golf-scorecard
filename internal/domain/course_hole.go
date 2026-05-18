package domain

// CourseHole represents one specific hole on a course with tee information
type CourseHole struct {
	ID               int      `json:"id" db:"id"`
	CourseID         int      `json:"course_id" db:"course_id"`
	HoleNumber       int      `json:"hole_number" db:"hole_number"`
	TeeColor         string   `json:"tee_color" db:"tee_color"`
	Par              int      `json:"par" db:"par"`
	DistanceFeet     *int     `json:"distance_feet,omitempty" db:"distance_feet"`
	DifficultyRating *float64 `json:"difficulty_rating,omitempty" db:"difficulty_rating"`
	Description      string   `json:"description,omitempty" db:"description"`
}

// NewCourseHole is used when adding holes to a course
type NewCourseHole struct {
	CourseID         int      `json:"course_id"`
	HoleNumber       int      `json:"hole_number"`
	TeeColor         string   `json:"tee_color"`
	Par              int      `json:"par"`
	DistanceFeet     *int     `json:"distance_feet,omitempty"`
	DifficultyRating *float64 `json:"difficulty_rating,omitempty"`
	Description      string   `json:"description,omitempty"`
}
