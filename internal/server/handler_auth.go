package server

import (
	"net/http"

	"github.com/alexedwards/scs/v2"
	"github.com/angelofallars/htmx-go"
	"github.com/bagashiz/sea-salon/internal/app/user"
	"github.com/bagashiz/sea-salon/internal/web/template"
)

// register is the handler for the registration form submission.
func register(userService *user.Service) handlerFunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		formValues, err := parseForm(r)
		if err != nil {
			return err
		}

		_, err = userService.CreateAccount(
			r.Context(),
			formValues["full_name"],
			formValues["phone_number"],
			formValues["email"],
			formValues["password"],
			formValues["confirm_password"],
		)
		if err != nil {
			return htmx.NewResponse().
				StatusCode(http.StatusUnprocessableEntity).
				Retarget("#auth").
				PreventPushURL().
				RenderTempl(r.Context(), w, template.RegisterForm(err))
		}

		return htmx.NewResponse().
			StatusCode(http.StatusCreated).
			Retarget("#auth").
			Reswap("transition:true").
			PushURL("/login/").
			RenderTempl(r.Context(), w, template.LoginForm(true, nil))
	}
}

// login is the handler for the login form submission.
func login(sessionManager *scs.SessionManager, userService *user.Service) handlerFunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		formValues, err := parseForm(r)
		if err != nil {
			return err
		}

		currentUser, err := userService.GetUserByEmail(
			r.Context(),
			formValues["email"],
			formValues["password"],
		)
		if err != nil {
			return htmx.NewResponse().
				StatusCode(http.StatusUnprocessableEntity).
				Retarget("#auth").
				PreventPushURL().
				RenderTempl(r.Context(), w, template.LoginForm(false, err))
		}

		sessionManager.Put(r.Context(), "user_id", currentUser.ID.String())

		return htmx.NewResponse().
			StatusCode(http.StatusOK).
			Retarget("main").
			Reswap("transition:true").
			PushURL("/").
			RenderTempl(r.Context(), w, template.LandingPage()) // TODO: redirect to the user's dashboard
	}
}
