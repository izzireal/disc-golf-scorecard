package repository

import (
	"database/sql"
	"disc-golf-scorecard/internal/domain"
)

// PlayerRepository handles all database operations for players
type PlayerRepository struct {
	db *sql.DB
}

// NewPlayerRepository creates a new player repository
func NewPlayerRepository(db *sql.DB) *PlayerRepository {
	return &PlayerRepository{db: db}
}

// Create adds a new player
func (r *PlayerRepository) Create(player *domain.Player) error {
	query := `
        INSERT INTO players (name, nickname)
        VALUES ($1, $2)
        RETURNING id, created_at`

	err := r.db.QueryRow(query, player.Name, player.Nickname).
		Scan(&player.ID, &player.CreatedAt)

	return err
}

// GetByID retrieves a player by ID
func (r *PlayerRepository) GetByID(id int) (*domain.Player, error) {
	player := &domain.Player{}

	query := `SELECT id, name, nickname, created_at FROM players WHERE id = $1`

	err := r.db.QueryRow(query, id).Scan(&player.ID, &player.Name, &player.Nickname, &player.CreatedAt)
	if err != nil {
		return nil, err
	}

	return player, nil
}

// List returns all players (useful for selecting players when creating a round)
func (r *PlayerRepository) List() ([]domain.PlayerSummary, error) {
	query := `SELECT id, name, nickname FROM players ORDER BY name`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var players []domain.PlayerSummary

	for rows.Next() {
		var p domain.PlayerSummary
		err := rows.Scan(&p.ID, &p.Name, &p.Nickname)
		if err != nil {
			return nil, err
		}
		players = append(players, p)
	}

	return players, nil
}
