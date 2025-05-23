package models

import (
	"context"
	"database/sql"
	"errors"
	"time"
)

var (
	dbTimeout = 5 * time.Second
)

type Quest struct {
	ID          int64      `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Reward      int        `json:"reward"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   *time.Time `json:"-"`
	Version     int32      `json:"version"`
}

func CreateNewQuest(db *sql.DB, quest Quest) error {
	query := `
		INSERT INTO quests (title, description, reward)
		VALUES ($1, $2, $3)
		RETURNING id, created_at, version;`

	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	args := []any{quest.Title, quest.Description, quest.Reward}

	var q Quest

	err := db.QueryRowContext(ctx, query, args...).Scan(&q.ID, &q.CreatedAt, &q.Version)
	if err != nil {
		return err
	}

	return nil
}

func GetAllQuests(db *sql.DB) ([]*Quest, error) {
	query := `
		SELECT id, title, description, reward, created_at, updated_at
		FROM quests;`

	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
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

func GetQuestByID(db *sql.DB, id int64) (*Quest, error) {
	query := `
		SELECT id, title, description, reward, created_at, updated_at, version
		FROM quests
		WHERE id = $1;`

	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	var q Quest

	args := id

	err := db.QueryRowContext(ctx, query, args).Scan(
		&q.ID,
		&q.Title,
		&q.Description,
		&q.Reward,
		&q.CreatedAt,
		&q.UpdatedAt,
		&q.Version,
	)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound

		default:
			return nil, err
		}
	}

	return &q, nil

}

func UpdateQuest(db *sql.DB, quest *Quest) error {
	query := `
	UPDATE quests
	SET title = $1, description = $2, reward = $3, version = version + 1
	WHERE id = $4 AND version = $5
	RETURNING version;
	`

	args := []any{
		quest.Title,
		quest.Description,
		quest.Reward,
		quest.ID,
		quest.Version,
	}

	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	err := db.QueryRowContext(ctx, query, args...).Scan(&quest.Version)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return ErrEditConflict
		default:
			return err
		}
	}

	return err
}

func DeleteQuest(db *sql.DB, id int64) error {
	query := `
		DELETE FROM quests
		WHERE id = $1;`

	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	result, err := db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	// Call the RowsAffected() method on the sql.Result object to get the number of rows
	// affected by the query.
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	// If no rows were affected, we know that the movies table didn't contain a record
	// with the provided ID at the moment we tried to delete it. In that case we
	// return an ErrRecordNotFound error.
	if rowsAffected == 0 {
		return ErrRecordNotFound
	}

	return nil
}
