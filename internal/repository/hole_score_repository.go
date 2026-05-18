package repository

import (
	"context"
	"database/sql"
	"fmt"

	"disc-golf-mvp/internal/domain"
)

// HoleScoreRepository handles all database operations for hole scores
type HoleScoreRepository struct {
	db *sql.DB
}

// NewHoleScoreRepository creates a new repository instance
func NewHoleScoreRepository(db *sql.DB) *HoleScoreRepository {
	return &HoleScoreRepository{db: db}
}

// Create inserts a new hole score
func (r *HoleScoreRepository) Create(ctx context.Context, score *domain.HoleScore) error {
	query := `
		INSERT INTO hole_scores (
			round_id, player_id, hole_number, score, 
			putts, fairway_hit, disc_used, notes
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		RETURNING id, created_at`

	err := r.db.QueryRowContext(ctx, query,
		score.RoundID,
		score.PlayerID,
		score.HoleNumber,
		score.Score,
		score.Putts,
		score.FairwayHit,
		score.DiscUsed,
		score.Notes,
	).Scan(&score.ID, &score.CreatedAt)

	if err != nil {
		return fmt.Errorf("failed to create hole score: %w", err)
	}
	return nil
}

// GetByRoundAndPlayer returns all hole scores for a player in a round
func (r *HoleScoreRepository) GetByRoundAndPlayer(ctx context.Context, roundID, playerID int64) ([]*domain.HoleScore, error) {
	query := `
		SELECT id, round_id, player_id, hole_number, score, 
		       putts, fairway_hit, disc_used, notes, created_at
		FROM hole_scores
		WHERE round_id = $1 AND player_id = $2
		ORDER BY hole_number ASC`

	rows, err := r.db.QueryContext(ctx, query, roundID, playerID)
	if err != nil {
		return nil, fmt.Errorf("failed to query hole scores: %w", err)
	}
	defer rows.Close()

	var scores []*domain.HoleScore
	for rows.Next() {
		score := &domain.HoleScore{}
		err := rows.Scan(
			&score.ID,
			&score.RoundID,
			&score.PlayerID,
			&score.HoleNumber,
			&score.Score,
			&score.Putts,
			&score.FairwayHit,
			&score.DiscUsed,
			&score.Notes,
			&score.CreatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan hole score: %w", err)
		}
		scores = append(scores, score)
	}

	return scores, nil
}

// Update updates an existing hole score (useful for corrections)
func (r *HoleScoreRepository) Update(ctx context.Context, score *domain.HoleScore) error {
	query := `
		UPDATE hole_scores
		SET score = $1, putts = $2, fairway_hit = $3,
		    disc_used = $4, notes = $5
		WHERE id = $6`

	result, err := r.db.ExecContext(ctx, query,
		score.Score,
		score.Putts,
		score.FairwayHit,
		score.DiscUsed,
		score.Notes,
		score.ID,
	)
	if err != nil {
		return fmt.Errorf("failed to update hole score: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return fmt.Errorf("no hole score found with id %d", score.ID)
	}

	return nil
}
