package main

import (
	"errors"
	"fmt"
	"muhlik20033.bgv2.net/internal/data"
	"muhlik20033.bgv2.net/internal/validator"
	"net/http"
)

func (app *application) createGameHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Title    string `json:"title"`
		Price    int32  `json:"price"`
		Color    string `json:"color"`
		Material string `json:"material"`
		Ages     string `json:"ages"`
	}
	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}
	game := &data.Game{
		Title:    input.Title,
		Price:    input.Price,
		Color:    input.Color,
		Material: input.Material,
		Ages:     input.Ages,
	}
	v := validator.New()

	if data.ValidateGame(v, game); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}
	//fmt.Fprintf(w, "%+v\n", input)

	err = app.models.Games.Insert(game)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	headers := make(http.Header)
	headers.Set("Location", fmt.Sprintf("/v1/games/%d", game.ID))

	err = app.writeJSON(w, http.StatusCreated, envelope{"game": game}, headers)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}

}

func (app *application) showGameHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}
	game, err := app.models.Games.Get(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}
	err = app.writeJSON(w, http.StatusOK, envelope{"game": game}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) updateGameHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}
	// Retrieve  movie record as normal.
	game, err := app.models.Games.Get(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	var input struct {
		Title    *string `json:"title"`
		Price    *int32  `json:"price"`
		Color    *string `json:"color"`
		Material *string `json:"material"`
		Ages     *string `json:"ages"`
	}
	err = app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	if input.Title != nil {
		game.Title = *input.Title
	}
	if input.Price != nil {
		game.Price = *input.Price
	}
	if input.Price != nil {
		game.Color = *input.Color
	}
	if input.Material != nil {
		game.Material = *input.Material
	}
	if input.Ages != nil {
		game.Ages = *input.Ages
	}

	v := validator.New()
	if data.ValidateGame(v, game); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	err = app.models.Games.Update(game)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrEditConflict):
			app.editConflictResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"game": game}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) deleteGameHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	err = app.models.Games.Delete(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}
	err = app.writeJSON(w, http.StatusOK, envelope{"message": "game successfully deleted"}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) listGamesHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Title string
		Color string
		data.Filters
	}
	v := validator.New()
	qs := r.URL.Query()

	input.Title = app.readString(qs, "title", "")
	input.Color = app.readString(qs, "color", "")

	input.Filters.Page = app.readInt(qs, "page", 1, v)
	input.Filters.PageSize = app.readInt(qs, "page_size", 20, v)

	input.Filters.Sort = app.readString(qs, "sort", "id")
	input.Filters.SortSafelist = []string{"id", "title", "color", "-id", "-title", "-color"}

	if data.ValidateFilters(v, input.Filters); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	games, metadata, err := app.models.Games.GetAll(input.Title, input.Color, input.Filters)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"games": games, "metadata": metadata}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

// BODY='{"title":"Flip Uno","price":250,"color":"Gray","material":"Laminated cardboard","ages":"+7"}'
// BODY='{"title":"Codenames","price":300,"color":"Gray","material":"Plastic", "ages":"+14"}'
// BODY='{"title":"One Piece Zoro And Sanji Starter Deck","price":280,"color":"White","material":"Laminated cardboard", "ages":"+12"}'
// BODY='{"title":"Scout","price":210,"color":"Yellow","material":"Carbon Fiber", "ages":"+9"}'
// psql "postgres://greenlight:pa55word@localhost/greenlight?sslmode=disable"
// SELECT * FROM games;
