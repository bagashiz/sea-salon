package server

import "net/http"

func addRoutes(mux *http.ServeMux) {
	mux.Handle("GET /assets/", handle(staticFiles()))

	mux.Handle("GET /", handle(notFound()))
	mux.Handle("GET /{$}", handle(index()))

	mux.Handle("GET /register/", handle(register()))
	mux.Handle("GET /login/", handle(login()))
}
