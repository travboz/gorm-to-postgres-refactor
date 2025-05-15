package models

import (
	"context"
	"database/sql"
	"time"
)

type Quest struct {
	ID          int        `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Reward      int        `json:"reward"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   *time.Time `json:"-"`
}

func CreateNewQuest(db *sql.DB, quest Quest) error {
	query := `
		INSERT INTO quests (title, description, reward) 
		VALUES ($1, $2, $3)
		RETURNING id, created_at;`

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	args := []any{quest.Title, quest.Description, quest.Reward}

	var q Quest

	err := db.QueryRowContext(ctx, query, args...).Scan(&q.ID, &q.CreatedAt)
	if err != nil {
		return err
	}

	return nil
}
