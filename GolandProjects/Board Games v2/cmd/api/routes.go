package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()

	router.NotFound = http.HandlerFunc(app.notFoundResponse)
	router.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowedResponse)

	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)

	router.HandlerFunc(http.MethodGet, "/v1/games", app.requirePermission("movies:read", app.listGamesHandler))
	router.HandlerFunc(http.MethodPost, "/v1/games", app.requirePermission("movies:write", app.createGameHandler))
	router.HandlerFunc(http.MethodGet, "/v1/games/:id", app.requirePermission("movies:read", app.showGameHandler))
	router.HandlerFunc(http.MethodPatch, "/v1/games/:id", app.requirePermission("movies:write", app.updateGameHandler))
	router.HandlerFunc(http.MethodDelete, "/v1/games/:id", app.requirePermission("movies:write", app.deleteGameHandler))

	router.HandlerFunc(http.MethodPost, "/v1/users", app.registerUserHandler)
	router.HandlerFunc(http.MethodPut, "/v1/users/activated", app.activateUserHandler)

	router.HandlerFunc(http.MethodPost, "/v1/tokens/authentication", app.createAuthenticationTokenHandler)

	return app.recoverPanic(app.enableCORS(app.rateLimit(app.authenticate(router))))
}
