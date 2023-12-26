package handlers

import (
	"github.com/Lincxx/go-web-app/pkg/config"
	"github.com/Lincxx/go-web-app/pkg/render"
	"net/http"
)

// repo pattern

// Repo the repository used by the handlers
var Repo *Repository

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

// NewRepo create a new resposit
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for the handles
func NewHandlers(r *Repository) {
	Repo = r
}

// handler func
// Home is the home page handler
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.gohtml")
}

// About is the about page handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page.gohtml")
}
