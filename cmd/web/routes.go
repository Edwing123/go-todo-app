package main

import (
	"net/http"

	"github.com/bmizerany/pat"
)

// Returns the router for the server
func (app *application) getRouter() http.Handler {
	router := pat.New()
	return router
}
