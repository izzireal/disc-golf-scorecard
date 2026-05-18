package repository

import (
	"database/sql"
	"disc-golf-scorecard/internal/domain"
)

// CourseRepository handles all database operations for courses
type CourseRepository struct {
	db *sql.DB
}

// NewCourseRepository creates a new repository
func NewCourseRepository(db *sql.DB) *CourseRepository {
	return &CourseRepository{db: db}
}

// Create adds a new course to the database
func (r *CourseRepository) Create(course *domain.Course) error {
	query := `
        INSERT INTO courses (name, location)
        VALUES ($1, $2)
        RETURNING id, created_at`

	err := r.db.QueryRow(query, course.Name, course.Location).
		Scan(&course.ID, &course.CreatedAt)

	return err
}

// GetByID retrieves a course by ID (with holes if requested)
func (r *CourseRepository) GetByID(id int) (*domain.Course, error) {
	course := &domain.Course{}

	query := `SELECT id, name, location, created_at FROM courses WHERE id = $1`
	err := r.db.QueryRow(query, id).Scan(&course.ID, &course.Name, &course.Location, &course.CreatedAt)
	if err != nil {
		return nil, err
	}

	return course, nil
}
