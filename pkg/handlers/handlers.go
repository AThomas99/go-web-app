package handlers

import (
	"hello-world/pkg/config"
	"hello-world/pkg/models"
	"hello-world/pkg/render"
	"net/http"
)

// Repo - the repository used by the handlers
var Repo *Repository

// Repository - Is the repository type
type Repository struct {
	App *config.AppConfig
}

// NewRepo - creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers - sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	// perform some logic

	// send the data to the template

	render.RenderTemplate(w, "home.page.html", &models.TemplateData{})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	// perform some logic
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello World"

	// send the data to the template
	render.RenderTemplate(w, "about.page.html", &models.TemplateData{
		StringMap: stringMap,
	})
}
