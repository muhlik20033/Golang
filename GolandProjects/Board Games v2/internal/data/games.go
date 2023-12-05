package data

import (
	"database/sql"
	"errors"
	_ "errors"
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
	v.Check(game.Price == 0, "price", "must be provided")
	v.Check(game.Price < 0, "price", "must be greater than 0")
	v.Check(len(game.Title) <= 500, "title", "must not be more than 500 bytes long")
	v.Check(game.Color != "", "color", "must be provided")
	v.Check(len(game.Color) <= 500, "color", "must not be more than 500 bytes long")
	v.Check(game.Material != "", "material", "must be provided")
	v.Check(len(game.Material) <= 500, "material", "must not be more than 100 bytes long")
	v.Check(game.Ages != "", "ages", "must be provided")
	v.Check(len(game.Ages) <= 10, "ages", "must not be more than 10 bytes long")
}

type GameModel struct {
	DB *sql.DB
}

func (m GameModel) Insert(game *Game) error {
	query := `
        INSERT INTO Games (title, price, color, material, ages)
        VALUES ($1, $2, $3, $4, $5)
        RETURNING id, created_at, version`
	args := []interface{}{game.Title, game.Price, game.Color, game.Material, game.Ages}
	return m.DB.QueryRow(query, args...).Scan(&game.ID, &game.CreatedAt, &game.Version)
}

func (m GameModel) Get(id int64) (*Game, error) {
	// The PostgreSQL bigserial type that we're using for the movie ID starts
	// auto-incrementing at 1 by default, so we know that no movies will have ID values
	// less than that. To avoid making an unnecessary database call, we take a shortcut
	// and return an ErrRecordNotFound error straight away.
	if id < 1 {
		return nil, ErrRecordNotFound
	}
	// Define the SQL query for retrieving the movie data.
	query := `
	SELECT id, created_at, title, price, color, material, ages, version
	FROM Games
	WHERE id = $1`
	// Declare a Movie struct to hold the data returned by the query.
	var game Game
	// Execute the query using the QueryRow() method, passing in the provided id value
	// as a placeholder parameter, and scan the response data into the fields of the
	// Movie struct. Importantly, notice that we need to convert the scan target for the
	// genres column using the pq.Array() adapter function again.
	err := m.DB.QueryRow(query, id).Scan(
		&game.ID,
		&game.CreatedAt,
		&game.Title,
		&game.Price,
		&game.Color,
		&game.Material,
		&game.Ages,
		&game.Version,
	)
	// Handle any errors. If there was no matching movie found, Scan() will return
	// a sql.ErrNoRows error. We check for this and return our custom ErrRecordNotFound
	// error instead.
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
	// Declare the SQL query for updating the record and returning the new version
	// number.
	query := `
		UPDATE Games
		SET title = $1, price = $2, color = $3, material = $4, ages = $5, version = version + 1
		WHERE id = $3
		RETURNING version`
	// Create an args slice containing the values for the placeholder parameters.
	args := []interface{}{
		game.Title,
		game.Price,
		game.Color,
		game.Material,
		game.Ages,
		game.ID,
	}
	return m.DB.QueryRow(query, args...).Scan(&game.Version)
}

func (m GameModel) Delete(id int64) error {
	if id < 1 {
		return ErrRecordNotFound
	}
	// Construct the SQL query to delete the record.
	query := `
	DELETE FROM Games
	WHERE id = $1`
	result, err := m.DB.Exec(query, id)
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
