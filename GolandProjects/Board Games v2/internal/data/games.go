package data

import (
	"context"
	"database/sql"
	"errors"
	_ "github.com/lib/pq"
	"muhlik20033.bgv2.net/internal/validator"
	"time"
)

type Game struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"-"`
	Title     string    `json:"title"`
	Price     int32     `json:"price"`
	Color     string    `json:"color"`
	Material  string    `json:"material"`
	Ages      string    `json.:"ages"`
	Version   int32     `json:"version"`
}

func ValidateGame(v *validator.Validator, game *Game) {
	v.Check(game.Title != "", "title", "must be provided")
	v.Check(len(game.Title) <= 500, "title", "must not be more than 500 bytes long")
	//v.Check(game.Price != 0, "price", "must be provided")
	//v.Check(game.Price < 0, "price", "must be greater than 0")
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
        INSERT INTO games (title, price, color, material, ages)
        VALUES ($1, $2, $3, $4, $5)
        RETURNING id, created_at, version`
	args := []interface{}{game.Title, game.Price, game.Color, game.Material, game.Ages}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	return m.DB.QueryRowContext(ctx, query, args...).Scan(&game.ID, &game.CreatedAt, &game.Version)
}

func (m GameModel) Get(id int64) (*Game, error) {
	if id < 1 {
		return nil, ErrRecordNotFound
	}
	query := `
	SELECT id, created_at, title, price, color, material, ages, version
	FROM games
	WHERE id = $1`

	var game Game
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := m.DB.QueryRowContext(ctx, query, id).Scan(
		&game.ID,
		&game.CreatedAt,
		&game.Title,
		&game.Price,
		&game.Color,
		&game.Material,
		&game.Ages,
		&game.Version,
	)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}
	return &game, nil
}

func (m GameModel) Update(game *Game) error {
	query := `
		UPDATE games
		SET title = $1, price = $2, color = $3, material = $4, ages = $5, version = version + 1
		WHERE id = $6 AND version = $7
		RETURNING version`

	args := []interface{}{
		game.Title,
		game.Price,
		game.Color,
		game.Material,
		game.Ages,
		game.ID,
		game.Version,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := m.DB.QueryRowContext(ctx, query, args...).Scan(&game.Version)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return ErrEditConflict
		default:
			return err
		}
	}
	return nil
}

func (m GameModel) Delete(id int64) error {
	if id < 1 {
		return ErrRecordNotFound
	}
	query := `
	    DELETE FROM games
	    WHERE id = $1`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	result, err := m.DB.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return ErrRecordNotFound
	}
	return nil
}
