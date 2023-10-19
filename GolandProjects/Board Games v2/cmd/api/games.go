package main

import (
	"fmt"
	"muhlik20033.bgv2.net/internal/data"
	"muhlik20033.bgv2.net/internal/validator"
	"net/http"
	"time"
)

func (app *application) createGameHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Title    string `json:"title"`
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
		Color:    input.Color,
		Material: input.Material,
		Ages:     input.Ages,
	}
	v := validator.New()
	if data.ValidateGame(v, game); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}
	fmt.Fprintf(w, "%+v\n", input)
}

func (app *application) showGameHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}
	game := data.Game{
		ID:        id,
		CreatedAt: time.Now(),
		Title:     "Uno Flip",
		Color:     "Violet",
		Material:  "Laminated cardboard",
		Ages:      "+7",
		Version:   1,
	}
	err = app.writeJSON(w, http.StatusOK, envelope{"game": game}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
