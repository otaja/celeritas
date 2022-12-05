package main

import (
	"log"
	"github.com/otaja/nexterp/data"
	"github.com/otaja/nexterp/handlers"
	"github.com/otaja/nexterp/middleware"
	"os"

	"github.com/otaja/celeritas"
)

func initApplication() *application {
	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	// init Celeritas
	cel := &celeritas.Celeritas{}
	err = cel.New(path)
	if err != nil {
		log.Fatal(err)
	}

	cel.AppName = "myApp"

	myMiddleware := &middleware.Middleware{
		App: cel,
	}

	myHandlers := &handlers.Handlers{
		App: cel,
	}

	app := &application{
		App:        cel,
		Handlers:   myHandlers,
		Middleware: myMiddleware,
	}

	app.App.Routes = app.routes()

	app.Models = data.New(app.App.DB.Pool)

	myHandlers.Models = app.Models

	app.Middleware.Models = app.Models

	return app
}
