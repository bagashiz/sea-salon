package server

import (
	"net/http"

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
		w.WriteHeader(http.StatusNotFound)
		return template.NotFound().Render(r.Context(), w)
	}
}

// index is the handler for the landing page.
func index() handlerFunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		if isHTMXRequest(r) {
			return template.LandingPage().Render(r.Context(), w)
		}
		return template.Index().Render(r.Context(), w)
	}
}

// register is the handler for the registration page and form component.
func register() handlerFunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		if isHTMXRequest(r) {
			return template.RegisterForm().Render(r.Context(), w)
		}
		return template.Register().Render(r.Context(), w)
	}
}

// login is the handler for the login page and form component.
func login() handlerFunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		if isHTMXRequest(r) {
			return template.LoginForm().Render(r.Context(), w)
		}
		return template.Login().Render(r.Context(), w)
	}
}

// isHTMXRequest checks request headers to determine if the request is an htmx request.
func isHTMXRequest(r *http.Request) bool {
	return r.Header.Get("HX-Request") == "true"
}
