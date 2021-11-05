package main

import (
	"fmt"
	"log"
	"net/http"

	"hello-world/pkg/config"
	"hello-world/pkg/handlers"
	"hello-world/pkg/render"
)

const portNumber = ":8080"


func main() {


	var app config.AppConfig

	tc ,err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	fmt.Println("Starting application on port number: ", portNumber)

	// srv - server variable for server properties
	srv := http.Server{
		Addr: portNumber,
		Handler: routes(&app),
	}
	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
