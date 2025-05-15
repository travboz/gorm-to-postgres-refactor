package models

import (
	"database/sql"
	"time"
)

type Quest struct {
	ID          uint      `json:"id" gorm:"primary_key"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Reward      int       `json:"reward"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func CreateNewQuest(db *sql.DB, quest Quest) error {
	return nil
}
