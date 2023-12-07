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
		game.Material = *input.Material // Note that we don't need to dereference a slice.
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
		app.serverErrorResponse(w, r, err)
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

// BODY='{"title":"Moana","year":2016,"runtime":"107 mins", "genres":["animation","adventure"]}'
//curl.exe -i -d "{"title":"Flip Uno","price":250,"color":"Gray","material":"Laminated cardboard","ages":"+7"}" localhost:4000/v1/games
//{"title":"Codenames","price":300,"color":"Gray","material":"Plastic", "ages":"+14"}
//{"title":"One Piece Zoro And Sanji Starter Deck","price":280,"color":"White","material":"Laminated cardboard", "ages":"+12"}
//'{"title":"Scout","price":210,"color":"Yellow","material":"Carbon Fiber", "ages":"+9"}'
//$BODYFromFile = Get-Content -Path "request.json" -Raw
//$BODY | Out-File -FilePath "request.json" -Encoding UTF8
