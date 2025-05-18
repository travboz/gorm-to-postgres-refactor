package models

import "errors"

var (
	ErrRecordNotFound = errors.New("record not found")
	ErrEditConflict   = errors.New("edit conflict") // when multiple people try and edit a quest simultaneously - a race-condition error
)
