package main

import (
	"fmt"
	"github.com/otaja/nexterp/data"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/otaja/celeritas/mailer"
)

func (a *application) routes() *chi.Mux {
	// middleware must become before routes

	// add routes here

	a.Get("/", a.Handlers.Home)

	// static routes

	fileServer := http.FileServer(http.Dir("./public"))
	a.App.Routes.Handle("/public/*", http.StripPrefix("/public", fileServer))

	return a.App.Routes
}
