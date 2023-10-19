package data

import (
	"muhlik20033.bgv2.net/internal/validator"
	"time"
)

type Game struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"-"`
	Title     string    `json:"title"`
	Color     string    `json:"color"`
	Material  string    `json:"material"`
	Ages      string    `json.:"ages"`
	Version   int32     `json:"version"`
}

func ValidateGame(v *validator.Validator, game *Game) {
	v.Check(game.Title != "", "title", "must be provided")
	v.Check(len(game.Title) <= 500, "title", "must not be more than 500 bytes long")
	v.Check(game.Color != "", "color", "must be provided")
	v.Check(len(game.Color) <= 100, "color", "must not be more than 500 bytes long")
	v.Check(game.Material != "", "material", "must be provided")
	v.Check(len(game.Material) <= 100, "material", "must not be more than 100 bytes long")
	v.Check(game.Ages != "", "ages", "must be provided")
	v.Check(len(game.Ages) <= 10, "ages", "must not be more than 10 bytes long")
}
