package data

import (
	"database/sql"
	"errors"
)

var (
	ErrRecordNotFound = errors.New("record not found")
	ErrEditConflict   = errors.New("edit conflict")
)

type Models struct {
	Games interface {
		Insert(game *Game) error
		Get(id int64) (*Game, error)
		Update(game *Game) error
		Delete(id int64) error
	}
}

func NewModels(db *sql.DB) Models {
	return Models{
		Games: GameModel{DB: db},
	}
}
