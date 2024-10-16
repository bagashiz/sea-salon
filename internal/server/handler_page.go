package server

import (
	"net/http"

	"github.com/angelofallars/htmx-go"
	"github.com/bagashiz/sea-salon/internal/web"
	"github.com/bagashiz/sea-salon/internal/web/template"
)

// staticFiles serves the static files such as CSS, JavaScript, and images.
func staticFiles() handlerFunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		http.FileServerFS(web.Assets).ServeHTTP(w, r)
		return nil
	}
}

// notFound is the handler for the 404 page.
func notFound() handlerFunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		statusCode := http.StatusNotFound
		w.WriteHeader(statusCode)
		return template.ErrorPage(statusCode).Render(r.Context(), w)
	}
}

// index is the handler for the landing page.
func index() handlerFunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		if htmx.IsHTMX(r) {
			return template.LandingPage().Render(r.Context(), w)
		}
		return template.Index().Render(r.Context(), w)
	}
}

// registerPage is the handler for the registration page and form component.
func registerPage() handlerFunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		if htmx.IsHTMX(r) {
			return template.Register().Render(r.Context(), w)
		}
		return template.RegisterPage().Render(r.Context(), w)
	}
}

// loginPage is the handler for the login page and form component.
func loginPage() handlerFunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		if htmx.IsHTMX(r) {
			return template.Login().Render(r.Context(), w)
		}
		return template.LoginPage().Render(r.Context(), w)
	}
}
