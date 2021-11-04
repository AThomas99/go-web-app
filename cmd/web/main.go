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

	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)

	fmt.Println("Starting application on port number: ", portNumber)
	http.ListenAndServe(portNumber, nil)

}
