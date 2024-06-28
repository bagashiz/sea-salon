package server

import (
	"net/http"

	"github.com/bagashiz/sea-salon/web"
	"github.com/bagashiz/sea-salon/web/templates"
)

// The staticFiles function serves the static files such as CSS, JavaScript, and images.
func staticFiles() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.FileServerFS(web.Assets).ServeHTTP(w, r)
	})
}

// The index function is the handler for the index page.
func index() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_ = templates.Index().Render(r.Context(), w)
	})
}
