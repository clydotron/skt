package main

import (
	"net/http"

	"github.com/bmizerany/pat"
)

func (app *application) routes() http.Handler {

	mux := pat.New()
	mux.Get("/", http.HandlerFunc(app.home))
	mux.Get("/api/ping", http.HandlerFunc(app.ping))
	return mux
}
