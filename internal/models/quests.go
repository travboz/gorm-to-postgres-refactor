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

func GetAllQuests(db *sql.DB) ([]*Quest, error) {
	query := `
		SELECT id, title, description, reward, created_at, updated_at
		FROM quests;`

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var quests []*Quest

	for rows.Next() {
		var q Quest

		err := rows.Scan(&q.ID, &q.Title, &q.Description, &q.Reward, &q.CreatedAt, &q.UpdatedAt)
		if err != nil {
			return nil, err
		}

		quests = append(quests, &q)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return quests, nil

}
