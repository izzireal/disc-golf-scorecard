package repository

import (
	"database/sql"
	"disc-golf-scorecard/internal/domain"
)

// RoundRepository handles all database operations for rounds
type RoundRepository struct {
	db *sql.DB
}

func NewRoundRepository(db *sql.DB) *RoundRepository {
	return &RoundRepository{db: db}
}

// Create starts a new round
func (r *RoundRepository) Create(round *domain.Round) error {
	query := `
        INSERT INTO rounds (course_id, date_played, group_code, notes)
        VALUES ($1, $2, $3, $4)
        RETURNING id, created_at`

	err := r.db.QueryRow(query, round.CourseID, round.DatePlayed, round.GroupCode, round.Notes).
		Scan(&round.ID, &round.CreatedAt)

	return err
}

// GetByID gets a round with its participants
func (r *RoundRepository) GetByID(id int) (*domain.Round, error) {
	round := &domain.Round{}

	query := `
        SELECT id, course_id, date_played, group_code, notes, created_at 
        FROM rounds WHERE id = $1`

	err := r.db.QueryRow(query, id).Scan(
		&round.ID, &round.CourseID, &round.DatePlayed,
		&round.GroupCode, &round.Notes, &round.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	return round, nil
}
