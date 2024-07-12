package server

import (
	"net/http"

	"github.com/alexedwards/scs/v2"
	"github.com/bagashiz/sea-salon/internal/app/user"
)

func addRoutes(mux *http.ServeMux, sessionManager *scs.SessionManager, userRepo user.ReadWriter) {
	mux.Handle("GET /assets/", handle(staticFiles()))

	mux.Handle("GET /", handle(notFound()))
	mux.Handle("GET /{$}", handle(index()))

	mux.Handle("GET /register/", handle(registerPage()))
	mux.Handle("GET /login/", handle(loginPage()))

	mux.Handle("POST /register", handle(register(userRepo)))
}
