package server

import (
	"net/http"

	"github.com/bagashiz/sea-salon/internal/web"
	"github.com/bagashiz/sea-salon/internal/web/template"
)

// staticFiles serves the static files such as CSS, JavaScript, and images.
func staticFiles() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.FileServerFS(web.Assets).ServeHTTP(w, r)
	})
}

// notFound is the handler for the 404 page.
func notFound() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		_ = template.NotFound().Render(r.Context(), w)
	})
}

// index is the handler for the landing page.
func index() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("HX-Request") == "true" {
			_ = template.LandingPage().Render(r.Context(), w)
			return
		}
		_ = template.Index().Render(r.Context(), w)
	})
}

// register is the handler for the registration page and form component.
func register() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("HX-Request") == "true" {
			_ = template.RegisterForm().Render(r.Context(), w)
			return
		}
		_ = template.Register().Render(r.Context(), w)
	})
}

// login is the handler for the login page and form component.
func login() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("HX-Request") == "true" {
			_ = template.LoginForm().Render(r.Context(), w)
			return
		}
		_ = template.Login().Render(r.Context(), w)
	})
}
