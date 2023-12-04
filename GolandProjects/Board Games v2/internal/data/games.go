package data

import (
	"database/sql"
	"github.com/lib/pq"
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

type GameModel struct {
	DB *sql.DB
}

func (m GameModel) Insert(game *Game) error {
	query := `
        INSERT INTO games (title, color, material, ages)
        VALUES ($1, $2, $3, $4)
        RETURNING id, created_at, version`
	args := []interface{}{game.Title, game.Color, game.Material, pq.Array(game.Ages)}
	return m.DB.QueryRow(query, args...).Scan(&game.ID, &game.CreatedAt, &game.Version)
}

func (m GameModel) Get(id int64) (*Game, error) {
	return nil, nil
}

func (m GameModel) Update(game *Game) error {
	return nil
}

func (m GameModel) Delete(id int64) error {
	return nil
}

type MockGameModel struct{}

func (m MockGameModel) Insert(game *Game) error {
	return nil
}
func (m MockGameModel) Get(id int64) (*Game, error) {
	return nil, nil
}
func (m MockGameModel) Update(game *Game) error {
	return nil
}
func (m MockGameModel) Delete(id int64) error {
	return nil
}
