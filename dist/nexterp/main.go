package main

import (
	"github.com/otaja/nexterp/data"
	"github.com/otaja/nexterp/handlers"
	"github.com/otaja/nexterp/middleware"

	"github.com/otaja/celeritas"
)

type application struct {
	App        *celeritas.Celeritas
	Handlers   *handlers.Handlers
	Models     data.Models
	Middleware *middleware.Middleware
}

func main() {
	c := initApplication()
	c.App.ListenAndServe()
}
