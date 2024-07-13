package server

import (
	"net/http"

	"github.com/angelofallars/htmx-go"
	"github.com/bagashiz/sea-salon/internal/app/user"
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

// register is the handler for the registration form submission.
func register(userRepo user.ReadWriter) handlerFunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		validationErrors := make(map[string]string)
		req := make(map[string]string)

		if err := r.ParseForm(); err != nil {
			return handlerError{http.StatusBadRequest, err.Error()}
		}

		req["full_name"] = r.FormValue("full_name")
		req["phone_number"] = r.FormValue("phone_number")
		req["email"] = r.FormValue("email")
		req["password"] = r.FormValue("password")
		req["confirm_password"] = r.FormValue("confirm_password")

		if req["password"] != req["confirm_password"] {
			validationErrors["confirm_password"] = "passwords do not match"
		}

		fullName := req["full_name"]
		if fullName == "" {
			validationErrors["full_name"] = "name cannot be empty"
		}

		phoneNumber, err := user.NewPhone(req["phone_number"])
		if err != nil {
			validationErrors["phone_number"] = err.Error()
		}

		email, err := user.NewEmail(req["email"])
		if err != nil {
			validationErrors["email"] = err.Error()
		}

		password, err := user.NewPassword(req["password"])
		if err != nil {
			validationErrors["password"] = err.Error()
		}

		hashedPassword, err := password.Hash()
		if err != nil {
			validationErrors["password"] = err.Error()
		}

		if len(validationErrors) > 0 {
			return htmx.NewResponse().
				StatusCode(http.StatusUnprocessableEntity).
				Retarget("#register-card").
				PreventPushURL().
				AddTrigger(
					htmx.TriggerObject("register-validation", map[string]any{
						"errors": validationErrors,
						"values": req,
					}),
				).
				RenderTempl(r.Context(), w, template.RegisterForm())
		}

		u, err := user.NewUser(fullName, phoneNumber, email, user.Password(hashedPassword), "customer")
		if err != nil {
			return htmx.NewResponse().
				StatusCode(http.StatusUnprocessableEntity).
				Retarget("#register-card").
				PreventPushURL().
				AddTrigger(
					htmx.TriggerDetail("register-error", err.Error()),
				).
				RenderTempl(r.Context(), w, template.RegisterForm())
		}
		u.Create()

		if err := userRepo.CreateUser(r.Context(), u); err != nil {
			if err == user.ErrUserExists {
				return htmx.NewResponse().
					StatusCode(http.StatusUnprocessableEntity).
					Retarget("#register-card").
					PreventPushURL().
					AddTrigger(
						htmx.TriggerDetail("register-error", err.Error()),
					).
					RenderTempl(r.Context(), w, template.RegisterForm())
			}
		}

		return htmx.NewResponse().
			StatusCode(http.StatusCreated).
			Retarget("main").
			Reswap("transition:true").
			PushURL("/login").
			RenderTempl(r.Context(), w, template.Login())
	}
}
