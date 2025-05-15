package env

import "database/sql"

type Env struct {
	DB *sql.DB
}

func NewEnv(db *sql.DB) *Env {
	return &Env{db}
}
